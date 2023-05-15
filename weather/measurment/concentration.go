package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	PartsPerMillion = value.NewLowerBoundUnit("Parts per million", " ppm", value.Dp3, 0)
	MicrogramsPerCubicMeter = value.NewLowerBoundUnit("Micrograms Per Cubic Meter", " µg/m³", value.Dp3, 0)

	// 1m³ of water has a mass of 1000kg, hence 1 ppm == 1000 µg/m³
	value.NewBasicBiTransform(PartsPerMillion, MicrogramsPerCubicMeter, 1000.0)
}

var (
	PartsPerMillion         value.Unit
	MicrogramsPerCubicMeter value.Unit
)
