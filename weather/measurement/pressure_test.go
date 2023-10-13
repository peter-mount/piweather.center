package measurement

import (
	"testing"
)

func Test_pressure_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		// Test basic values
		{PressurePA.Value(101325), PressureHPA.Value(1013.25), false},
		{PressurePA.Value(101325), PressureMBar.Value(1013.25), false},
		//
		{PressureHPA.Value(1), PressurePA.Value(100), false},
		{PressureHPA.Value(1), PressureKPA.Value(0.1), false},
		{PressureMBar.Value(1), PressurePA.Value(100), false},
		{PressureMBar.Value(1), PressureBar.Value(0.001), false},
		{PressurePA.Value(100000), PressureBar.Value(1), false},
		// From a real issue, for some reason 29.973inHg came out as 101.5hPa and not 1015hPa
		{PressureInHg.Value(29.973), PressurePA.Value(101500.2267169408), false},
		{PressureInHg.Value(29.973), PressureHPA.Value(1015.002267169408), false},
		{PressureInHg.Value(29.973), PressureKPA.Value(101.5002267169408), false},
	})
}
