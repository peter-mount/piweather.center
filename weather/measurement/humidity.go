package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	RelativeHumidity = value.NewBoundedUnit("RelativeHumidity", "Relative Humidity", "%", 0, 0, 100)
	Humidity = value.NewGroup("Humidity", RelativeHumidity)

	value.NewCalculator("absoluteHumidity", TemperatureRelativeHumidityCalculator(GetAbsoluteHumidity))

}

var (
	Humidity         *value.Group
	RelativeHumidity *value.Unit
)

// GetAbsoluteHumidity Get the absolute humidity (amount of water vapor in the air) in metric.
func GetAbsoluteHumidity(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return TemperatureRelativeHumidityCalculation(temp, relHumidity, Celsius, getAbsoluteHumidity)
}

func getAbsoluteHumidity(temp value.Value, relHumidity value.Value) (value.Value, error) {
	t := temp.Float()
	return GramsPerCubicMeter.Value((6.112 * math.Exp((17.67*t)/(t+243.5)) * relHumidity.Float() * 2.1674) / (273.15 + t)), nil
}
