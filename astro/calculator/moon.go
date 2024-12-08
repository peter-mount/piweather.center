package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/moonposition"
	"github.com/soniakeys/meeus/v3/nutation"
)

var (
// obliquity = coord.NewObliquity(unit.AngleFromDeg(23.4392911))
)

func (c *calculator) CalculateMoon(t value.Time) (api.EphemerisResult, error) {
	jd := julian.FromTime(t.Time())

	// Δψ nutation in longitude
	// Δε nutation in obliquity
	_, Δε := nutation.Nutation(jd.Float())
	obliquity := coord.NewObliquity(nutation.MeanObliquityLaskar(jd.Float()) + Δε)

	lon, lat, R := moonposition.Position(jd.Float())

	return api.NewEphemerisResult("moon", t).
			SetEcliptic2(lat, lon, obliquity).
			SetDistance(measurement.Kilometers.Value(R)),
		nil
}
