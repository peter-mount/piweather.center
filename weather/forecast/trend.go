package forecast

import (
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Trend int8

const (
	// The drop/rise limit in mbar
	delta = 1.6
)

const (
	// TrendFalling when a drop of 1.6 mbar in 3 hours
	TrendFalling Trend = iota - 1
	// TrendSteady when no drop or rise of 1.6 mbar in 3 hours
	TrendSteady
	// TrendRising when rise of 1.6 mbar in 3 hours
	TrendRising
)

// GetTrend returns the trend between two pressure values.
// Ideally, p0 should be the pressure from 3 hours before p1
func GetTrend(p0, p1 value.Value) (Trend, error) {
	m0, err := p0.As(measurement.PressureMBar)
	if err != nil {
		return 0, err
	}

	m1, err := p1.As(measurement.PressureMBar)
	if err != nil {
		return 0, err
	}

	d := m0.Float() - m1.Float()

	switch {
	case value.GreaterThanEqual(d, delta):
		return TrendFalling, nil
	case value.LessThanEqual(d, -delta):
		return TrendRising, nil
	default:
		return TrendSteady, nil
	}
}
