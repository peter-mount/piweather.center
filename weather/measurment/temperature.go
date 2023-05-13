package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

const (
	Celsius0Kelvin = 273.15
	MinCelsius     = -Celsius0Kelvin
	MinFahrenheit  = -459.67
	MinKelvin      = 0.0
)

func init() {
	Fahrenheit = value.NewLowerBoundUnit("Fahrenheit", " °F", value.Dp1, MinFahrenheit)
	Celsius = value.NewLowerBoundUnit("Celsius", " °C", value.Dp1, MinCelsius)
	Kelvin = value.NewLowerBoundUnit("Kelvin", " K", value.Dp1, MinKelvin)

	value.NewTransform(Celsius, Kelvin, celsiusKelvin)
	value.NewTransform(Kelvin, Celsius, kelvinCelsius)
	value.NewTransform(Fahrenheit, Celsius, fahrenheitCelsius)
	value.NewTransform(Celsius, Fahrenheit, celsiusFahrenheit)
	value.NewTransform(Fahrenheit, Kelvin, func(f float64) (float64, error) {
		v, _ := fahrenheitCelsius(f)
		return celsiusKelvin(v)
	})
	value.NewTransform(Kelvin, Fahrenheit, func(f float64) (float64, error) {
		v, _ := kelvinCelsius(f)
		return celsiusFahrenheit(v)
	})

}

func celsiusKelvin(f float64) (float64, error) { return f + Celsius0Kelvin, nil }

func kelvinCelsius(f float64) (float64, error) { return f - Celsius0Kelvin, nil }

func fahrenheitCelsius(f float64) (float64, error) { return (f - 32.0) * 5.0 / 9.0, nil }

func celsiusFahrenheit(f float64) (float64, error) { return (f * 9.0 / 5.0) + 32.0, nil }

var (
	Fahrenheit value.Unit
	Celsius    value.Unit
	Kelvin     value.Unit
)
