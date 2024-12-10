package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sun"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/nutation"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
)

// SolarAltitudeCalculator
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

// SolarAzimuthCalculator
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

// SolarHZ
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

	jde := julian.FromTime(t.Time()).Float()

	// Based on solar.ApparentEquatorialVSOP87 but, like CalculatePlanet we need additional information,
	// but we also cut out an eclToEq conversion
	//ra, dec, R := solar.ApparentEquatorialVSOP87(earth, jd.Float())
	s, β, R := solar.TrueVSOP87(earth, jde)

	// Obliquity
	Δψ, Δε := nutation.Nutation(jde)
	a := aberration(R)
	λ := s + Δψ + a
	ε := nutation.MeanObliquity(jde) + Δε

	// Extra parameters here

	return api.NewEphemerisResult("sun", t).
			SetObliquity(ε).
			SetEcliptic(λ, β).
			SetDistance(measurement.AU.Value(R)).
			SetSemiDiameter(measurement.AngleRoundDown(measurement.Degree.Value(SemiDiameter(station.EphemerisTargetSun, R).Deg()))).
			SetLightTime(measurement.DurationRoundDown(measurement.DurationDay.Value(base.LightTime(R)))),
		nil
}
