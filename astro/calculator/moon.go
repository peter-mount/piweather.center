package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
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

	return api.NewEphemerisResult("moon", t).
			SetObliquity(ε).
			SetEcliptic(β, λ).
			SetDistance(measurement.Kilometers.Value(R)),
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
