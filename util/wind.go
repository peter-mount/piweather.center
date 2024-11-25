package util

import (
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

type WindCompass uint8

func (w WindCompass) String() string {
	return windCompassDirection[w]
}

// Int returns the value as an index between 0..15
func (w WindCompass) Int() int {
	return int(w)
}

// Degree returns the WindCompass in degrees, N=0 increasing clockwise
func (w WindCompass) Degree() float64 {
	return float64(w) * 22.5
}

// Add returns WindCompass with an increment (0..15) added to it.
// To subtract pass a negative value
func (w WindCompass) Add(i int) WindCompass {
	return WindCompass((int(w) + i) % 16)
}

// AddDegree returns a WindCompass with the appropriate number of degrees added.
// Clockwise is positive, CounterClockwise is negative.
func (w WindCompass) AddDegree(i float64) WindCompass {
	return WindCompassDirection(w.Degree() + i)
}

const (
	WindN = iota
	WindNNE
	WindNE
	WindENE
	WindE
	WindESE
	WindSE
	WindSSE
	WindS
	WindSSW
	WindSW
	WindWSW
	WindW
	WindWNW
	WindNW
	WindNNW
)

var (
	windCompassDirection = []string{
		"N", "NNE", "NE", "ENE",
		"E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW",
		"W", "WNW", "NW", "NNW",
	}
)

func WindCompassDirection(d float64) WindCompass {
	for d < 0 {
		d += 360
	}
	for d >= 360 {
		d -= 360
	}
	return WindCompass(int(math.Floor((d+11.25)/22.5)) % 16)
}

func WindCompassDirectionDegrees(d value.Value) (WindCompass, error) {
	v, err := d.As(measurement.Degree)
	if err == nil {
		err = v.BoundsError()
	}
	if err != nil {
		return 0, err
	}

	return WindCompassDirection(v.Float()), nil
}
