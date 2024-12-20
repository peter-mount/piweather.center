package calculator

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/weather/value"
)

func (c *calculator) SolarSystem(t value.Time) (api.EphemerisDay, error) {
	ephem := api.NewEphemerisDay("", t)

	e, err := c.CalculateSun(t)
	if err != nil {
		return nil, err
	}
	ephem.Add(e)

	e, err = c.CalculateMoon(t)
	if err != nil {
		return nil, err
	}
	ephem.Add(e)

	for p := station.EphemerisTargetMercury; p <= station.EphemerisTargetNeptune; p++ {
		if p != station.EphemerisTargetEarth {
			e, err = c.CalculatePlanet(p, t)
			if err != nil {
				return nil, err
			}
			ephem.Add(e)
		}
	}

	return ephem, nil
}
