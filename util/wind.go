package util

import "math"

var (
	windCompassDirection = []string{
		"N", "NNE", "NE", "ENE",
		"E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW",
		"W", "WNW", "NW", "NNW",
	}
)

func WindCompassDirection(d float64) string {
	for d < 0 {
		d += 360
	}
	for d >= 360 {
		d -= 360
	}
	e := int(math.Floor((d+11.25)/22.5)) % 16
	return windCompassDirection[e]
}
