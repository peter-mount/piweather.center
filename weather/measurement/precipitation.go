package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	MillimetersPerHour = value.NewLowerBoundUnit("MillimetersPerHour", "Millimeters per hour", " mm/h", 3, 0)
	InchesPerHour = value.NewLowerBoundUnit("InchesPerHour", "Inches per hour", " in/h", 3, 0)

	value.NewBasicBiTransform(InchesPerHour, MillimetersPerHour, mmToM/inToM)

	Precipitation = value.NewGroup("Precipitation", MillimetersPerHour, InchesPerHour)

}

var (
	Precipitation      *value.Group
	MillimetersPerHour *value.Unit
	InchesPerHour      *value.Unit
)
