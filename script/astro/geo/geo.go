package geo

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
)

func init() {
	packages.RegisterPackage(&Geo{})
}

type Geo struct{}

// LatLong returns the Geographic coordinate on Earth.
//
// Note: longitude is positive-East here, the opposite to normal geography
func (_ Geo) LatLong(lat, lon, altitude float64) *coord.LatLong {
	return &coord.LatLong{
		Latitude:  unit.AngleFromDeg(lat),
		Longitude: unit.AngleFromDeg(lon),
		Altitude:  altitude,
	}
}

// Coord returns the Meeus coordinate on Earth.
//
// Note: longitude is positive-West here, the opposite to normal geography
func (_ Geo) Coord(lat, lon float64) *globe.Coord {
	return &globe.Coord{
		Lat: unit.AngleFromDeg(lat),
		Lon: unit.AngleFromDeg(lon),
	}
}
