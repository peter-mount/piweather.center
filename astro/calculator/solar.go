package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sun"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
)

// Deprecated
func (c *calculator) SolarAltitudeCalculator() value.Calculator {
	return func(t value.Time, _ ...value.Value) (value.Value, error) {
		_, h, err := c.SolarHZ(t)
		if err != nil {
			return value.Value{}, err
		}
		return measurement.Degree.Value(h.Deg()), nil
	}
}

// Deprecated
func (c *calculator) SolarAzimuthCalculator() value.Calculator {
	return func(t value.Time, _ ...value.Value) (value.Value, error) {
		A, _, err := c.SolarHZ(t)
		if err != nil {
			return value.Value{}, err
		}
		return measurement.Degree.Value(A.Deg()), nil
	}
}

// Deprecated
func (c *calculator) SolarHZ(t value.Time) (unit.Angle, unit.Angle, error) {
	earth, err := c.Planet(planetposition.Earth)
	if err != nil {
		return 0, 0, err
	}

	jd := julian.FromTime(t.Time())
	loc := t.Location()
	A, h := sun.ApparentHzVSOP87(jd, loc.Lat, loc.Lon, earth)
	return A, h, nil
}

func (c *calculator) CalculateSun(t value.Time) (api.EphemerisResult, error) {
	earth, err := c.Planet(planetposition.Earth)
	if err != nil {
		return nil, err
	}

	jd := julian.FromTime(t.Time())

	ra, dec, R := solar.ApparentEquatorialVSOP87(earth, jd.Float())

	return api.NewEphemerisResult("sun", t).
			SetEquatorial(ra, dec).
			SetDistance(measurement.AU.Value(R)),
		nil
}
