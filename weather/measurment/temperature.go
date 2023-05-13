package measurment

import (
	"errors"
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
	Fahrenheit   value.Unit
	Celsius      value.Unit
	Kelvin       value.Unit
	notTempError = errors.New("value not a Temperature")
)

func IsTemperature(v value.Value) bool {
	u := v.Unit()
	return u == Fahrenheit || u == Celsius || u == Kelvin
}

// AssertTemperature returns an error if the value is not a temperature value, or it's value is invalid
func AssertTemperature(v value.Value) error {
	// If it is an error then call BoundsError which will be nil unless the temperature is invalid
	if IsTemperature(v) {
		return v.BoundsError()
	}
	return notTempError
}

func IsTemperatureErr(e error) bool {
	return e == notTempError
}

// TemperatureRelativeHumidityFunc is a common function for values based on temperature and relative humidity
type TemperatureRelativeHumidityFunc func(temp, relHumidity value.Value) (value.Value, error)

// TemperatureRelativeHumidityCalculation wraps a TemperatureRelativeHumidityFunc ensuring that the values passed
// are of the correct units.
//
// Specifically that temp is a Temperature and relHumidity is RelativeHumidity.
// It also takes a unit which is the temperature unit required by the underlying function.
func TemperatureRelativeHumidityCalculation(temp, relHumidity value.Value, unit value.Unit, f TemperatureRelativeHumidityFunc) (value.Value, error) {
	if err := AssertTemperature(temp); err != nil {
		return value.Value{}, err
	}
	if err := AssertRelativeHumidity(relHumidity); err != nil {
		return value.Value{}, err
	}

	// Convert temp into unit
	t1, err := temp.As(unit)
	if err != nil {
		return value.Value{}, err
	}

	return f(t1, relHumidity)
}
