package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	GramsPerCubicMeter = value.NewLowerBoundUnit("GramsPerCubicMeter", "Grams Per Cubic Meter", "g/m³", 3, 0)
	PoundsPerCubitFoot = value.NewLowerBoundUnit("PoundsPerCubitFoot", "Pounds Per Cubit Foot", "lbs/ft³", 3, 0)
	value.NewBasicBiTransform(PoundsPerCubitFoot, GramsPerCubicMeter, cubicFootToCubicMeter/poundToGram)
	Density = value.NewGroup("Density", GramsPerCubicMeter, PoundsPerCubitFoot)
}

var (
	Density            *value.Group
	GramsPerCubicMeter *value.Unit
	PoundsPerCubitFoot *value.Unit
)
