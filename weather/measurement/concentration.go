package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	PartsPerMillion = value.NewLowerBoundUnit("PartsPerMillion", "Parts per million", " ppm", 0, 0)
	MicrogramsPerCubicMeter = value.NewLowerBoundUnit("MicrogramsPerCubicMeter", "Micrograms Per Cubic Meter", " µg/m³", 0, 0)

	// 1m³ of water has a mass of 1000kg, hence 1 ppm == 1000 µg/m³
	value.NewBasicBiTransform(PartsPerMillion, MicrogramsPerCubicMeter, 1000.0)

	Concentration = value.NewGroup("Concentration", PartsPerMillion, MicrogramsPerCubicMeter)
}

var (
	Concentration           *value.Group
	PartsPerMillion         *value.Unit
	MicrogramsPerCubicMeter *value.Unit
)
