package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	MillimetersPerHour = value.NewLowerBoundUnit("MillimetersPerHour", "Precipitation", "Millimeters per hour", " mm/h", 3, 0)
	InchesPerHour = value.NewLowerBoundUnit("InchesPerHour", "Precipitation", "Inches per hour", " in/h", 3, 0)

	value.NewBasicBiTransform(InchesPerHour, MillimetersPerHour, mmToM/inToM)
}

var (
	MillimetersPerHour *value.Unit
	InchesPerHour      *value.Unit
)
