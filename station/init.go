package station

import (
	"context"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

func (s *Stations) Init() error {
	return NewVisitor().
		Station(s.initStation).
		Sensors(s.initSensors).
		Reading(s.initReading).
		CalculatedValue(s.initCalculatedValue).
		WithContext(context.Background()).
		VisitStations(s)
}

func (s *Stations) initStation(ctx context.Context) error {
	station := StationFromContext(ctx)
	station.ID = ctx.Value("StationId").(string)

	lat, err := util.ParseAngle(station.Location.Latitude)
	if err != nil {
		return err
	}

	lon, err := util.ParseAngle(station.Location.Longitude)
	if err != nil {
		return err
	}

	station.latLong = &coord.LatLong{
		Longitude: lon,
		Latitude:  lat,
		Altitude:  station.Location.Altitude,
		Name:      station.Location.Name,
		Notes:     station.Location.Notes,
	}

	return nil
}

func (s *Stations) initSensors(ctx context.Context) error {
	sensors := SensorsFromContext(ctx)

	// Set the VisitStation.ID
	stationConfig := StationFromContext(ctx)
	sensors.ID = stationConfig.ID + "." + ctx.Value("SensorId").(string)

	sensors.station = StationFromContext(ctx)

	return nil
}

func (s *Stations) initReading(ctx context.Context) error {
	parent := SensorsFromContext(ctx)
	reading := ReadingFromContext(ctx)
	reading.ID = parent.ID + "." + ctx.Value("ReadingId").(string)
	reading.sensors = parent

	// If not ok then we will ignore the reading
	if u, ok := value.GetUnit(reading.Type); ok {
		reading.unit = u

		// Use the same unit, unless we declare an alternate
		reading.useUnit = u
		if reading.Use != "" {
			// TODO if the unit is not ok or there's no transform then this will fail over to use the src unit
			if u, ok := value.GetUnit(reading.Use); ok && value.TransformAvailable(reading.useUnit, u) {
				reading.useUnit = u
			}
		}
	}
	return nil
}

func (s *Stations) initCalculatedValue(ctx context.Context) error {
	parent := SensorsFromContext(ctx)
	calculation := CalculatedValueFromContext(ctx)

	// Ensure ID is set and the Source entries have the same prefix
	prefix := parent.ID + "."
	calculation.ID = strings.ToLower(prefix + ctx.Value("ReadingId").(string))

	calculation.sensors = parent

	for i, src := range calculation.Source {
		var s string
		switch {
		// Keywords not to expand.
		// "current" the metrics current value
		case src == "current":
			s = src

		// If a calculation source has a "." then it's an absolute reference and
		// not local to the reading
		case strings.Contains(src, "."):
			s = src

		default:
			s = prefix + src
		}
		calculation.Source[i] = strings.ToLower(s)
	}

	// Lookup the Calculator to use
	calc, err := value.GetCalculator(calculation.Type)
	if err != nil {
		return err
	}
	calculation.calculator = calc

	// If we have Use set then try to convert to that unit
	if calculation.Use != "" {
		to, ok := value.GetUnit(calculation.Use)
		if ok {
			calculation.calculator = calculation.calculator.As(to)
		}
	}

	return nil
}
