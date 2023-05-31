package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	RelativeHumidity = value.NewBoundedUnit("RelativeHumidity", "Relative Humidity", "%", 0, 0, 100)
	Humidity = value.NewGroup("Humidity", RelativeHumidity)

	value.NewCalculator("absoluteHumidity", value.AssertCalculator(value.Calculator2arg(GetAbsoluteHumidity), Temperature.AssertValue, RelativeHumidity.AssertValue))
}

var (
	Humidity         *value.Group
	RelativeHumidity *value.Unit
)

// GetAbsoluteHumidity Get the absolute humidity (amount of water vapor in the air) in metric.
func GetAbsoluteHumidity(temp value.Value, relHumidity value.Value) (value.Value, error) {
	temp, err := temp.As(Celsius)
	if err != nil {
		return value.Value{}, err
	}

	relHumidity, err = relHumidity.As(RelativeHumidity)
	if err != nil {
		return value.Value{}, err
	}

	t := temp.Float()
	return GramsPerCubicMeter.Value((6.112 * math.Exp((17.67*t)/(t+243.5)) * relHumidity.Float() * 2.1674) / (273.15 + t)), nil
}
