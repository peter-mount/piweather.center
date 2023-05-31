package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

// WindChill temperature is only defined for temperatures at or below 50F and wind speeds above 3mph
func WindChill(temp, windSpeed value.Value) (value.Value, error) {
	temp, err := temp.As(Fahrenheit)
	if err != nil {
		return value.Value{}, err
	}

	windSpeed, err = windSpeed.As(MetersPerSecond)
	if err != nil {
		return value.Value{}, err
	}

	T := temp.Float()
	V := math.Pow(windSpeed.Float(), 0.16)

	return Fahrenheit.Value(35.74 + (0.6215 * T) - (35.75 * V) + (0.4275 * T * V)), nil
}
