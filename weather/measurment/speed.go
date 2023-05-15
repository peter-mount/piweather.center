package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	MetersPerSecond = value.NewLowerBoundUnit("Meters Per Second", " m/s", value.Dp1, 0)
	KilometersPerHour = value.NewLowerBoundUnit("Kilometers Per Hour", " kph", value.Dp1, 0)
	MilesPerHour = value.NewLowerBoundUnit("Miles Per Hour", " mph", value.Dp1, 0)
	FeetPerSecond = value.NewLowerBoundUnit("Feet Per Second", " fps", value.Dp1, 0)
	Knots = value.NewLowerBoundUnit("Knots", " knots", value.Dp1, 0)

	// Transforms between mps and each unit - this registers both directions
	value.NewBasicBiTransform(MetersPerSecond, KilometersPerHour, mpsToKph)
	value.NewBasicBiTransform(MetersPerSecond, MilesPerHour, mpsToMph)
	value.NewBasicBiTransform(MetersPerSecond, FeetPerSecond, feetPerSecond)
	value.NewBasicBiTransform(MetersPerSecond, Knots, mpsToKnots)

	// Conversions between units other than MetersPerSecond. These convert to MetersPerSecond first then to the final one
	value.NewTransformations(MetersPerSecond, FeetPerSecond, KilometersPerHour, Knots, MilesPerHour)
}

var (
	FeetPerSecond     value.Unit
	KilometersPerHour value.Unit
	Knots             value.Unit
	MetersPerSecond   value.Unit
	MilesPerHour      value.Unit
)
