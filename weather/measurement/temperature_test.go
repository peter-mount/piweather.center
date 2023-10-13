package measurement

import (
	"testing"
)

func Test_temperature_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		{Celsius.Value(0), Fahrenheit.Value(32.0), false},
		{Celsius.Value(10), Fahrenheit.Value(50.0), false},
		{Fahrenheit.Value(32), Celsius.Value(0.0), false},
		{Fahrenheit.Value(50), Celsius.Value(10.0), false},
		{Celsius.Value(0), Kelvin.Value(273.15), false},
		{Celsius.Value(10), Kelvin.Value(283.15), false},
		{Kelvin.Value(273.15), Celsius.Value(0.0), false},
		{Kelvin.Value(283.15), Celsius.Value(10.0), false},
		{Kelvin.Value(0), Celsius.Value(-273.15), false},
		// This is invalid as cannot be colder than Absolute Zero so expect an error
		{Kelvin.Value(-1), Celsius.Value(-274.15), true},
		{Celsius.Value(-274.15), Kelvin.Value(-1), true},
		{Kelvin.Value(0), Fahrenheit.Value(-460.67), true},
		{Fahrenheit.Value(-460.67), Kelvin.Value(0), true},
	})
}
