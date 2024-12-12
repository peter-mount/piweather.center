package measurement

import (
	"testing"
)

func Test_length_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// Metric
		// =========================
		newConversionTest(Kilometers.Value(1), Meters.Value(1000), false),
		newConversionTest(CentiMeters.Value(1), Meters.Value(0.01), false),
		newConversionTest(MilliMeters.Value(1), Meters.Value(0.001), false),
		// =========================
		// Inches
		// =========================
		newConversionTest(Inches.Value(1), Meters.Value(0.0254), false),
		newConversionTest(Inches.Value(1), CentiMeters.Value(2.54), false),
		newConversionTest(Inches.Value(1), MilliMeters.Value(25.4), false),
		newConversionTest(Inches.Value(12), Feet.Value(1), false),
		newConversionTest(Inches.Value(36), Yard.Value(1), false),
		newConversionTest(Inches.Value(63360), Miles.Value(1), false),
		// =========================
		// Feet
		// =========================
		newConversionTest(Feet.Value(1), Meters.Value(0.3048), false),
		newConversionTest(Feet.Value(1), CentiMeters.Value(30.48), false),
		newConversionTest(Feet.Value(1), MilliMeters.Value(304.8), false),
		newConversionTest(Feet.Value(3), Yard.Value(1), false),
		newConversionTest(Feet.Value(5280), Miles.Value(1), false),
		// =========================
		// Yard
		// =========================
		newConversionTest(Yard.Value(1), Meters.Value(0.9144), false),
		newConversionTest(Yard.Value(1), CentiMeters.Value(91.44), false),
		newConversionTest(Yard.Value(1), MilliMeters.Value(914.4), false),
		newConversionTest(Yard.Value(1760), Miles.Value(1), false),
		// =========================
		// Miles
		// =========================
		newConversionTest(Miles.Value(1), Kilometers.Value(1.609344), false),
		newConversionTest(Miles.Value(1), Meters.Value(1609.344), false),
		newConversionTest(Miles.Value(1), CentiMeters.Value(160934.4), false),
		newConversionTest(Miles.Value(1), MilliMeters.Value(1609344), false),
	})
}
