package measurement

import "testing"

func Test_precipitation(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// MillimetersPerHour
		// =========================
		newConversionTest(MillimetersPerHour.Value(1), InchesPerHour.Value(25.4), false),
		// =========================
		// InchesPerHour
		// =========================
		newConversionTest(InchesPerHour.Value(1), MillimetersPerHour.Value(0.0393700), false),
	})
}
