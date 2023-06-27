package memory

import (
	"context"
	"fmt"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Store is a version of Store and Archiver that can hold a single day's
// archives
type Store struct {
	storeDir string                 // Path to store directory
	config   *station.Stations      // Loaded stations
	history  map[string][]*Reading  // Loaded history
	visitor  station.VisitorBuilder // Visitor for loading data
}

func New(dir string, config *station.Stations) (*Store, error) {
	fs, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !fs.IsDir() {
		return nil, fmt.Errorf("%q not a directory", dir)
	}

	ms := &Store{
		storeDir: dir,
		config:   config,
	}
	ms.Clear()

	return ms, nil
}

// Close implements io.Closer
func (s *Store) Close() error {
	s.Clear()
	return nil
}

func (s *Store) Clear() {
	s.history = make(map[string][]*Reading)
}

func (s *Store) Get(id string) []*Reading {
	return s.history[id]
}

func (s *Store) GetAt(id string, t time.Time) *Reading {
	var r *Reading
	for _, reading := range s.Get(id) {
		if reading.Time.After(t) {
			return r
		}
		r = reading
	}
	return r
}

// Load a stations data for a time
func (s *Store) Load(t time.Time) error {
	s.Clear()

	s.visitor = station.NewVisitor().
		Sensors(value.ResetMap).
		Reading(s.processReading).
		CalculatedValue(s.calculate)

	ctx := value.WithMap(context.Background())
	ctx = context.WithValue(ctx, "time", t)

	return s.config.Accept(station.NewVisitor().
		Sensors(s.loadSensors).
		WithContext(ctx))
}

func ArchiveFileName(dir, name string, t time.Time) string {
	p := filepath.Join(strings.Split(name, ".")...)
	return filepath.Join(dir, p, t.UTC().Format("2006/01/02")+".log")
}

func (s *Store) loadSensors(ctx context.Context) error {
	sensors := station.SensorsFromContext(ctx)
	t := ctx.Value("time").(time.Time)

	fileName := ArchiveFileName(s.storeDir, sensors.ID, t)

	err := io.NewReader().
		ForEachLine(func(line string) error {
			p, err := payload.FromLog(line)
			if p != nil && err == nil {
				_ = s.visitor.
					WithContext(p.AddContext(ctx)).
					VisitSensors(sensors)
			}
			return nil
		}).
		Open(fileName)

	// It is perfectly fine for the file not to exist
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

func (s *Store) processReading(ctx context.Context) error {
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

			s.history[r.ID] = append(s.history[r.ID], &Reading{
				Name:  r.ID,
				Value: v,
				Time:  p.Time(),
			})
		}
	}
	return nil
}

func (s *Store) calculate(ctx context.Context) error {
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

	values.Put(calc.ID, result)

	s.history[calc.ID] = append(s.history[calc.ID], &Reading{
		Name:  calc.ID,
		Value: result,
		Time:  p.Time(),
	})

	return nil
}
