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
}

func celsiusKelvin(f float64) (float64, error) { return f + Celsius0Kelvin, nil }

func kelvinCelsius(f float64) (float64, error) { return f - Celsius0Kelvin, nil }

func fahrenheitCelsius(f float64) (float64, error) { return (f - 32.0) * 5.0 / 9.0, nil }

func celsiusFahrenheit(f float64) (float64, error) { return (f * 9.0 / 5.0) + 32.0, nil }

// TemperatureRelativeHumidityFunc is a common function for values based on temperature and relative humidity
type TemperatureRelativeHumidityFunc func(temp, relHumidity value.Value) (value.Value, error)

// TemperatureRelativeHumidityCalculation wraps a TemperatureRelativeHumidityFunc ensuring that the values passed
// are of the correct units.
//
// Specifically that temp is a Temperature and relHumidity is RelativeHumidity.
// It also takes a unit which is the temperature unit required by the underlying function.
//
// Examples of this function in use are DewPoint and HeatIndex values which are a function of Temperature and Relative Humidity.
//
// This will return an error if:
//   - temp, relHumidity or unit are not of the appropriate Units
//   - the values of temp or relHumidity are invalid
//   - the temperature to be passed to the function when transformed to unit is invalid
//   - the function returns an error
//   - the final result is invalid
func TemperatureRelativeHumidityCalculation(temp, relHumidity value.Value, unit *value.Unit, f TemperatureRelativeHumidityFunc) (value.Value, error) {
	if err := Temperature.AssertValue(temp); err != nil {
		return value.Value{}, err
	}
	if err := RelativeHumidity.AssertValue(relHumidity); err != nil {
		return value.Value{}, err
	}

	// Convert temp into unit
	t1, err := temp.As(unit)
	if err != nil {
		return value.Value{}, err
	}
	if !t1.IsValid() {
		return value.Value{}, t1.BoundsError()
	}

	r, err := f(t1, relHumidity)
	if err != nil {
		return value.Value{}, err
	}
	return r, r.BoundsError()
}

// TemperatureRelativeHumidityCalculator returns a Calculator that expects 2 parameters,
// Temperature and RelativeHumidity
func TemperatureRelativeHumidityCalculator(f TemperatureRelativeHumidityFunc) value.Calculator {
	return value.AssertCalculator(func(_ value.Time, v ...value.Value) (value.Value, error) {
		return f(v[0], v[1])
	},
		Temperature.AssertValue,
		RelativeHumidity.AssertValue)
}
