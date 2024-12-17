package astro

import (
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
)

type Astro struct {
	Angle Angle // Angle functions
	//Equatorial Equatorial // Equatorial functions
}

// Coord returns the Meeus coordinate on Earth.
//
// Note: longitude is positive-West here, the opposite to normal geography
func (_ Astro) Coord(lat, lon float64) *globe.Coord {
	return &globe.Coord{
		Lat: unit.AngleFromDeg(lat),
		Lon: unit.AngleFromDeg(lon),
	}
}

//type Equatorial struct{}
//
//func (_ Equatorial) FromEq(e coord2.Equatorial) coord.Equatorial {
//	return coord.NewFromEq(e)
//}
