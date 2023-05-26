package station

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	kernel.RegisterAPI((*Config)(nil), &config{})
}

// Config provides access to the Stations config
type Config interface {
	Accept(v Visitor) error
	Stations() *Stations
}

type config struct {
	Config *Stations `kernel:"config,stations"`
}

func (c *config) Stations() *Stations {
	return c.Config
}

func (c *config) Accept(v Visitor) error {
	return v.VisitStations(c.Config)
}

func (c *config) Start() error {
	if c.Config == nil || len(*c.Config) == 0 {
		return fmt.Errorf("no configuration provided")
	}

	// Once loaded ensure the structure is intact and ids are set up
	return NewVisitor().
		Station(c.initStation).
		Sensors(c.initSensors).
		Reading(c.initReading).
		CalculatedValue(c.initCalculatedValue).
		WithContext(context.Background()).
		VisitStations(c.Config)
}

func (c *config) initStation(ctx context.Context) error {
	s := StationFromContext(ctx)
	s.ID = ctx.Value("StationId").(string)

	lat, err := util.ParseAngle(s.Location.Latitude)
	if err != nil {
		return err
	}

	lon, err := util.ParseAngle(s.Location.Longitude)
	if err != nil {
		return err
	}

	s.latLong = &coord.LatLong{
		Longitude: lon,
		Latitude:  lat,
		Altitude:  s.Location.Altitude,
		Name:      s.Location.Name,
		Notes:     s.Location.Notes,
	}

	return nil
}

func (c *config) initSensors(ctx context.Context) error {
	s := SensorsFromContext(ctx)

	// Set the VisitStation.ID
	stationConfig := StationFromContext(ctx)
	s.ID = stationConfig.ID + "." + ctx.Value("SensorId").(string)

	s.station = StationFromContext(ctx)

	return nil
}

func (c *config) initReading(ctx context.Context) error {
	parent := SensorsFromContext(ctx)
	s := ReadingFromContext(ctx)
	s.ID = parent.ID + "." + ctx.Value("ReadingId").(string)

	// If not ok then we will ignore the reading
	if u, ok := value.GetUnit(s.Type); ok {
		s.unit = u

		// Use the same unit, unless we declare an alternate
		s.useUnit = u
		if s.Use != "" {
			// TODO if the unit is not ok or there's no transform then this will fail over to use the src unit
			if u, ok := value.GetUnit(s.Use); ok && value.TransformAvailable(s.useUnit, u) {
				s.useUnit = u
			}
		}
	}
	return nil
}

func (c *config) initCalculatedValue(ctx context.Context) error {
	parent := SensorsFromContext(ctx)
	s := CalculatedValueFromContext(ctx)

	// Ensure ID is set and the Source entries have the same prefix
	prefix := parent.ID + "."
	s.ID = prefix + ctx.Value("ReadingId").(string)

	for i, src := range s.Source {
		s.Source[i] = prefix + src
	}

	// Lookup the Calculator to use
	calc, err := value.GetCalculator(s.Type)
	if err != nil {
		return err
	}
	s.calculator = calc

	// If we have Use set then try to convert to that unit
	if s.Use != "" {
		to, ok := value.GetUnit(s.Use)
		if ok {
			s.calculator = s.calculator.As(to)
		}
	}

	return nil
}
