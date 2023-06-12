package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

const (
	Celsius0Kelvin = 273.15
	MinCelsius     = -Celsius0Kelvin
	MinFahrenheit  = -459.67
	MinKelvin      = 0.0
)

var (
	Temperature *value.Group
	Celsius     *value.Unit
	Fahrenheit  *value.Unit
	Kelvin      *value.Unit
)

func init() {
	Fahrenheit = value.NewLowerBoundUnit("Fahrenheit", "Fahrenheit", " °F", 1, MinFahrenheit)
	Celsius = value.NewLowerBoundUnit("Celsius", "Celsius", " °C", 1, MinCelsius)
	Kelvin = value.NewLowerBoundUnit("Kelvin", "Kelvin", " K", 1, MinKelvin)

	value.NewTransform(Celsius, Kelvin, celsiusKelvin)
	value.NewTransform(Kelvin, Celsius, kelvinCelsius)
	value.NewTransform(Fahrenheit, Celsius, fahrenheitCelsius)
	value.NewTransform(Celsius, Fahrenheit, celsiusFahrenheit)
	value.NewTransform(Fahrenheit, Kelvin, value.Of(fahrenheitCelsius, celsiusKelvin))
	value.NewTransform(Kelvin, Fahrenheit, value.Of(kelvinCelsius, celsiusFahrenheit))

	Temperature = value.NewGroup("Temperature", Celsius, Fahrenheit, Kelvin)

	value.NewCalculator("dewPoint", value.AssertCalculator(value.Calculator2arg(GetDewPoint), Temperature.AssertValue, RelativeHumidity.AssertValue))
	value.NewCalculator("heatIndex", value.AssertCalculator(value.Calculator2arg(HeatIndex), Temperature.AssertValue, RelativeHumidity.AssertValue))
	value.NewCalculator("windChill", value.AssertCalculator(value.Calculator2arg(WindChill), Temperature.AssertValue, Speed.AssertValue))
	value.NewCalculator("feelsLike", value.AssertCalculator(value.Calculator3arg(FeelsLike), Temperature.AssertValue, RelativeHumidity.AssertValue, MetersPerSecond.AssertValue))
}

func celsiusKelvin(f float64) (float64, error) { return f + Celsius0Kelvin, nil }

func kelvinCelsius(f float64) (float64, error) { return f - Celsius0Kelvin, nil }

func fahrenheitCelsius(f float64) (float64, error) { return (f - 32.0) * 5.0 / 9.0, nil }

func celsiusFahrenheit(f float64) (float64, error) { return (f * 9.0 / 5.0) + 32.0, nil }
