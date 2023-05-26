package store

import (
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/template"
	"github.com/peter-mount/piweather.center/weather/value"
	"golang.org/x/net/context"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Store struct {
	Templates *template.Manager `kernel:"inject"`
	Cron      *cron.CronService `kernel:"inject"`
	Config    station.Config    `kernel:"inject"`
	mutex     sync.Mutex
	data      map[string]*Reading
	history   map[string][]*Reading
}

const (
	storeMaxAge = time.Hour * 26 // Max time to keep readings
)

func FromContext(ctx context.Context) *Store {
	return ctx.Value("local.store").(*Store)
}

func (s *Store) AddContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "local.store", s)
}

type Reading struct {
	Name  string
	Value value.Value
	Time  time.Time
}

func (r *Reading) String() string {
	return strings.Join([]string{
		r.Name,
		strconv.FormatFloat(r.Value.Float(), 'f', 3, 64),
		strconv.Itoa(int(r.Time.UTC().Unix())),
	}, " ")
}

func (s *Store) PostInit() error {
	s.Templates.AddFunction("getReadingKeys", s.GetKeys)
	s.Templates.AddFunction("getReading", s.GetReading)
	s.Templates.AddFunction("getReadingHistory", s.GetHistory)
	return nil
}

func (s *Store) Start() error {
	s.data = make(map[string]*Reading)
	s.history = make(map[string][]*Reading)

	// Register every reading in the store so we have an entry for them
	err := s.Config.Accept(station.NewVisitor().
		Reading(s.registerReading).
		WithContext(context.Background()))

	if err == nil {
		// Every 10 minutes clear down history
		_, err = s.Cron.AddTask("0/10 * * * ?", s.pruneTask)
	}

	return err
}

func (s *Store) registerReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)

	name := strings.ToLower(r.ID)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.data[name]; !exists {
		rec := &Reading{Name: name, Value: r.Unit().Value(0)}
		s.data[name] = rec
		s.history[name] = []*Reading{rec}
	}

	return nil
}

func (s *Store) ProcessReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)
	values := value.MapFromContext(ctx)
	if r.Unit() != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(r.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := r.Value(f)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			values.Put(r.ID, v)

			s.Record(r.ID, v, p.Time())
		}
	}
	return nil
}

func (s *Store) Calculate(ctx context.Context) error {
	// Get value.Time from Station and Payload
	sensors := station.SensorsFromContext(ctx)
	p := payload.GetPayload(ctx)
	t := sensors.Station().LatLong().Time(p.Time())

	calc := station.CalculatedValueFromContext(ctx)

	values := value.MapFromContext(ctx)
	args := values.GetAll(calc.Source...)

	result, err := calc.Calculate(t, args...)
	if err != nil {
		return err
	}

	s.Record(calc.ID, result, p.Time())

	return nil
}

func (s *Store) Record(name string, value value.Value, recTime time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	name = strings.ToLower(name)

	rec := &Reading{
		Name:  name,
		Value: value,
		Time:  recTime,
	}

	// Prepend to history, keeping only last 60 entries
	hist := s.history[name]
	hist = append(hist, rec)
	sort.SliceStable(hist, func(i, j int) bool {
		return hist[i].Time.Before(hist[j].Time)
	})

	// Prune the history as we add an entry
	hist = s.pruneHistory(hist)
	s.history[name] = hist

	// Cache latest value which is at end, but only if we have data
	if len(hist) > 0 {
		s.data[name] = hist[len(hist)-1]
	}
}

func (s *Store) GetReading(name string) *Reading {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	name = strings.ToLower(name)
	return s.data[name]
}

func (s *Store) GetHistory(name string) []*Reading {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	name = strings.ToLower(name)
	return s.history[name]
}

func (s *Store) GetKeys() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var keys []string
	for k, _ := range s.data {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func (s *Store) pruneTask(_ context.Context) error {
	keys := s.GetKeys()
	for _, k := range keys {
		s.prune(k)
	}
	return nil
}

func (s *Store) prune(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if hist, exists := s.history[key]; exists {
		s.history[key] = s.pruneHistory(hist)
	}
}

func (s *Store) pruneHistory(hist []*Reading) []*Reading {
	cutoff := time.Now().Add(-storeMaxAge)

	// >2 so we keep the last 2 entries regardless of how old
	// they are - so on status page we have a last received time
	// if a sensor is offline.
	for len(hist) > 1 && hist[0].Time.Before(cutoff) {
		hist = hist[1:]
	}

	return hist
}
