package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/moonposition"
)

func (c *calculator) CalculateMoon(t value.Time) (api.EphemerisResult, error) {
	//earth, err := c.Planet(planetposition.Earth)
	//if err != nil {
	//	return nil, err
	//}

	jd := julian.FromTime(t.Time())

	lon, lat, R := moonposition.Position(jd.Float())

	ra, dec := lon.RA(), lat

	return api.NewEphemerisResult("moon", t).
			SetEquatorial(ra, dec).
			SetDistance(measurement.Kilometers.Value(R)),
		nil
}
