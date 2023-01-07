package sun

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
)

func ApparentEquatorial(jd julian.Day) coord.Equatorial {
	a, d := solar.ApparentEquatorial(jd.JD())
	return coord.New(a, d)
}

func ApparentEquatorialVSOP87(jd julian.Day) coord.Equatorial {
	e, err := planetposition.LoadPlanet(planetposition.Earth)
	if err != nil {
		panic(err)
	}

	a, d, _ := solar.ApparentEquatorialVSOP87(e, jd.JD())
	return coord.New(a, d)
}
