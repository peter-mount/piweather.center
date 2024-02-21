package astro

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/soniakeys/unit"
)

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
