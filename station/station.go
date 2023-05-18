package station

import (
	"context"
	archiver "github.com/peter-mount/piweather.center/server/archiver"
	"github.com/peter-mount/piweather.center/station/payload"
)

// Station defines a Weather Station at a specific location.
// It consists of one or more Reading's
type Station struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Name of the station
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Location of the station
	Location Location `json:"location" xml:"location,omitempty" yaml:"location,omitempty"`
	// One or more Sensors collection
	Sensors map[string]*Sensors `json:"sensors" xml:"sensors" yaml:"sensors"`
}

func (s *Station) Init(ctx context.Context) error {
	return s.call(ctx, func(sensors *Sensors, ctx context.Context) error {
		return sensors.init(ctx)
	})
}

func (s *Station) call(ctx context.Context, f func(*Sensors, context.Context) error) error {
	child := context.WithValue(ctx, "Station", s)
	for id, sensors := range s.Sensors {
		child1 := context.WithValue(child, "SensorId", id)
		if err := f(sensors, child1); err != nil {
			return err
		}
	}
	return nil
}

// Sensors define a Reading collection within the Station.
// A Reading collection is
type Sensors struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Name of the Readings collection
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Source of data for this collection
	Source Source `json:"source" xml:"source" yaml:"source"`
	// Format of the message, default is json
	Format string
	// Timestamp Path to timestamp, "" for none
	Timestamp string
	// Reading's provided by this collection
	Readings map[string]*Reading `json:"readings" xml:"readings" yaml:"readings"`
}

func (s *Sensors) init(ctx context.Context) error {
	// Set the Station.ID
	station := ctx.Value("Station").(*Station)
	s.ID = station.ID + "." + ctx.Value("SensorId").(string)

	if err := s.call(ctx, func(sensor *Reading, ctx context.Context) error {
		return sensor.init(ctx)
	}); err != nil {
		return err
	}

	// Preload from logs
	_ = archiver.FromContext(ctx).Preload(ctx, s.ID, s.process)

	return nil
}

func (s *Sensors) Process(ctx context.Context) error {
	archiver.FromContext(ctx).Archive(payload.GetPayload(ctx))
	return s.process(ctx)
}

func (s *Sensors) process(ctx context.Context) error {
	return s.call(ctx, func(sensor *Reading, ctx context.Context) error {
		return sensor.process(ctx)
	})
}

func (s *Sensors) call(ctx context.Context, f func(*Reading, context.Context) error) error {
	child := context.WithValue(ctx, "Sensors", s)
	for id, sensor := range s.Readings {
		child2 := context.WithValue(child, "ReadingId", id)
		if err := f(sensor, child2); err != nil {
			return err
		}
	}
	return nil
}
