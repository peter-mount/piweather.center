package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/moonposition"
	"github.com/soniakeys/meeus/v3/nutation"
	"github.com/soniakeys/unit"
)

var (
// obliquity = coord.NewObliquity(unit.AngleFromDeg(23.4392911))
)

func (c *calculator) CalculateMoon(t value.Time) (api.EphemerisResult, error) {
	jd := julian.FromTime(t.Time())

	λ, β, R := moonposition.Position(jd.Float())

	// Δψ nutation in longitude
	// Δε nutation in obliquity
	Δψ, Δε := nutation.Nutation(jd.Float())
	a := aberration(R)
	λ = λ + Δψ + a
	ε := nutation.MeanObliquity(jd.Float()) + Δε

	// Light Time from earth in days, but we need to convert R (km) to AU first
	rAU := measurement.Kilometers.Value(R).AsOrInvalid(measurement.AU).Float()
	τEarth := base.LightTime(rAU)

	return api.NewEphemerisResult("moon", t).
			SetObliquity(ε).
			SetEcliptic(λ, β).
			SetDistance(measurement.Kilometers.Value(R)).
			SetSemiDiameter(measurement.AngleRoundDown(measurement.Degree.Value(SemiDiameter(station.EphemerisTargetMoon, rAU).Deg()))).
			// Light Time is in Days but DurationRoundDown makes it more useful
			SetLightTime(measurement.DurationRoundDown(measurement.DurationDay.Value(τEarth))),
		nil
}

// Low precision formula.  The high precision formula is not implemented
// because the low precision formula already gives position results to the
// accuracy given on p. 165.  The high precision formula the represents lots
// of typing with associated chance of typos, and no way to test the result.
func aberration(R float64) unit.Angle {
	// (25.10) p. 167
	return unit.AngleFromSec(-20.4898).Div(R)
}
