package sun

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sidereal"
	coord2 "github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
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

// ApparentHzVSOP87 calculates the azimuth and elevation of the sun
//
//	jd: Julian Day
//	φ: latitude of observer on Earth
//	ψ: longitude of observer on Earth
//	st: sidereal time at Greenwich at time of observation.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
//
// Results:
//
//	A: azimuth of observed point, measured westward from the South.
//	h: elevation, or height of observed point above horizon.
func ApparentHzVSOP87(jd julian.Day, φ, ψ unit.Angle, posEarth *planetposition.V87Planet) (unit.Angle, unit.Angle) {
	st := sidereal.FromJD(jd)

	α, δ, _ := solar.ApparentEquatorialVSOP87(posEarth, jd.JD())
	A, h := coord2.EqToHz(α, δ, φ, ψ, st)
	return A, h
}
