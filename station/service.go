package station

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
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
		WithContext(context.Background()).
		VisitStations(c.Config)
}

func (c *config) initStation(ctx context.Context) error {
	s := StationFromContext(ctx)
	s.ID = ctx.Value("StationId").(string)
	return nil
}

func (c *config) initSensors(ctx context.Context) error {
	s := SensorsFromContext(ctx)

	// Set the VisitStation.ID
	stationConfig := StationFromContext(ctx)
	s.ID = stationConfig.ID + "." + ctx.Value("SensorId").(string)

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
