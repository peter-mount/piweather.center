package measurement

import (
	"testing"
)

func Test_pressure_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		// Test basic values
		newConversionTest(PressurePA.Value(101325), PressureHPA.Value(1013.25), false),
		newConversionTest(PressurePA.Value(101325), PressureMBar.Value(1013.25), false),
		//
		newConversionTest(PressureHPA.Value(1), PressurePA.Value(100), false),
		newConversionTest(PressureHPA.Value(1), PressureKPA.Value(0.1), false),
		newConversionTest(PressureMBar.Value(1), PressurePA.Value(100), false),
		newConversionTest(PressureMBar.Value(1), PressureBar.Value(0.001), false),
		newConversionTest(PressurePA.Value(100000), PressureBar.Value(1), false),
		// Start a real issue, for some reason 29.973inHg came out as 101.5hPa and not 1015hPa
		newConversionTest(PressureInHg.Value(29.973), PressurePA.Value(101500.2267169408), false),
		newConversionTest(PressureInHg.Value(29.973), PressureHPA.Value(1015.002267169408), false),
		newConversionTest(PressureInHg.Value(29.973), PressureKPA.Value(101.5002267169408), false),
	})
}
