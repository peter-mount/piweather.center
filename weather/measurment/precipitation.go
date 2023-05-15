package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	MillimetersPerHour = value.NewLowerBoundUnit("Millimeters per hour", " mm/h", value.Dp3, 0)
	InchesPerHour = value.NewLowerBoundUnit("Inches per hour", " in/h", value.Dp3, 0)

	value.NewBasicBiTransform(InchesPerHour, MillimetersPerHour, mmToM/inToM)
}

var (
	MillimetersPerHour value.Unit
	InchesPerHour      value.Unit
)
