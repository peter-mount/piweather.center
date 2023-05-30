package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	// Register the Calculator which enforces
	value.NewCalculator("dewPoint", TemperatureRelativeHumidityCalculator(GetDewPoint))
}

// GetDewPoint returns a Value representing the Dewpoint based on Temperature and RelativeHumidity.
// The returned value is in Celsius.
func GetDewPoint(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return TemperatureRelativeHumidityCalculation(temp, relHumidity, Celsius, getDewPoint)
}

// temp must be Celsius
func getDewPoint(temp value.Value, relHumidity value.Value) (value.Value, error) {
	t0, rh := temp.Float(), relHumidity.Float()

	b, c := 17.368, 238.88
	if t0 <= 0 {
		b, c = 17.966, 247.15
	}
	pa := math.Log(rh / 100.0 * math.Exp(b*t0/(c+t0)))
	dp := c * pa / (b - pa)
	return Celsius.Value(dp), nil
}
