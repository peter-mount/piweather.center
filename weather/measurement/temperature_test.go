package measurement

import (
	"testing"
)

func Test_temperature_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		newConversionTest(Celsius.Value(0), Fahrenheit.Value(32.0), false),
		newConversionTest(Celsius.Value(10), Fahrenheit.Value(50.0), false),
		newConversionTest(Fahrenheit.Value(32), Celsius.Value(0.0), false),
		newConversionTest(Fahrenheit.Value(50), Celsius.Value(10.0), false),
		newConversionTest(Celsius.Value(0), Kelvin.Value(273.15), false),
		newConversionTest(Celsius.Value(10), Kelvin.Value(283.15), false),
		newConversionTest(Kelvin.Value(273.15), Celsius.Value(0.0), false),
		newConversionTest(Kelvin.Value(283.15), Celsius.Value(10.0), false),
		newConversionTest(Kelvin.Value(0), Celsius.Value(-273.15), false),
		// This is invalid as cannot be colder than Absolute Zero so expect an error
		newConversionTest(Kelvin.Value(-1), Celsius.Value(-274.15), true),
		newConversionTest(Celsius.Value(-274.15), Kelvin.Value(-1), true),
		newConversionTest(Kelvin.Value(0), Fahrenheit.Value(-459.67), true),
		newConversionTest(Fahrenheit.Value(-459.67), Kelvin.Value(0), true),
	})
}
