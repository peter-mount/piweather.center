package astro

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sun"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/unit"
)

type Sun struct{}

func (_ Sun) ApparentEquatorial(jd julian.Day) coord.Equatorial {
	return sun.ApparentEquatorial(jd)
}

func (_ Sun) ApparentEquatorialVSOP87(jd julian.Day) (coord.Equatorial, error) {
	return sun.ApparentEquatorialVSOP87(jd)
}

func (_ Sun) ApparentHzVSOP87(jd julian.Day, φ, ψ unit.Angle, posEarth *planetposition.V87Planet) (unit.Angle, unit.Angle) {
	return sun.ApparentHzVSOP87(jd, φ, ψ, posEarth)
}
