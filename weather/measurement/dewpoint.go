package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

// GetDewPoint returns a Value representing the Dewpoint based on Temperature and RelativeHumidity.
// The returned value is in Celsius.
func GetDewPoint(temp value.Value, relHumidity value.Value) (value.Value, error) {
	temp, err := temp.As(Celsius)
	if err != nil {
		return value.Value{}, err
	}

	relHumidity, err = relHumidity.As(RelativeHumidity)
	if err != nil {
		return value.Value{}, err
	}

	t0, rh := temp.Float(), relHumidity.Float()

	b, c := 17.368, 238.88
	if t0 <= 0 {
		b, c = 17.966, 247.15
	}
	pa := math.Log(rh / 100.0 * math.Exp(b*t0/(c+t0)))
	dp := c * pa / (b - pa)
	return Celsius.Value(dp), nil
}
