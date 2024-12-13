package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

// PressureMeanSeaLevel calculates the pressure in HPA at mean sea level
func PressureMeanSeaLevel(pressure, temperature, altitude value.Value) (value.Value, error) {
	pr, err := pressure.As(PressureHPA)
	if err != nil {
		return value.Value{}, err
	}

	temp, err := temperature.As(Kelvin)
	if err != nil {
		return value.Value{}, err
	}

	alt, err := altitude.As(Meters)
	if err != nil {
		return value.Value{}, err
	}

	h := 0.0065 * alt.Float()

	p := math.Pow(1-(h/(temp.Float()+h)), -5.257)

	return pr.Multiply(PressureHPA.Value(p))
}
