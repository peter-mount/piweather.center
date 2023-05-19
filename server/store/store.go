package store

import (
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
	mutex     sync.Mutex
	data      map[string]*Reading
	history   map[string][]*Reading
}

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
	return nil
}

func (s *Store) DeclareReading(name string, unit *value.Unit) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.data[name]; !exists {
		rec := &Reading{Name: name, Value: unit.Value(0)}
		s.data[name] = rec
		s.history[name] = []*Reading{rec}
	}
}

func (s *Store) Record(name string, value value.Value, time time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rec := &Reading{
		Name:  name,
		Value: value,
		Time:  time,
	}

	// Prepend to history, keeping only last 60 entries
	hist := s.history[name]
	hist = append(hist, rec)
	sort.SliceStable(hist, func(i, j int) bool {
		return hist[i].Time.After(hist[j].Time)
	})
	if len(hist) > 60 {
		hist = hist[:60]
	}
	s.history[name] = hist

	// Cache latest value which is at front
	s.data[name] = hist[0]
}

func (s *Store) GetReading(name string) *Reading {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.data[name]
}

func (s *Store) GetHistory(name string) []*Reading {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
