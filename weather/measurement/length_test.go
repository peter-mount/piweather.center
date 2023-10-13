package measurement

import (
	"testing"
)

func Test_length_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// Metric
		// =========================
		{Kilometers.Value(1), Meters.Value(1000), false},
		{CentiMeters.Value(1), Meters.Value(0.01), false},
		{MilliMeters.Value(1), Meters.Value(0.001), false},
		// =========================
		// Inches
		// =========================
		{Inches.Value(1), Meters.Value(0.0254), false},
		{Inches.Value(1), CentiMeters.Value(2.54), false},
		{Inches.Value(1), MilliMeters.Value(25.4), false},
		{Inches.Value(12), Feet.Value(1), false},
		{Inches.Value(36), Yard.Value(1), false},
		{Inches.Value(63360), Miles.Value(1), false},
		// =========================
		// Feet
		// =========================
		{Feet.Value(1), Meters.Value(0.3048), false},
		{Feet.Value(1), CentiMeters.Value(30.48), false},
		{Feet.Value(1), MilliMeters.Value(304.8), false},
		{Feet.Value(3), Yard.Value(1), false},
		{Feet.Value(5280), Miles.Value(1), false},
		// =========================
		// Yard
		// =========================
		{Yard.Value(1), Meters.Value(0.9144), false},
		{Yard.Value(1), CentiMeters.Value(91.44), false},
		{Yard.Value(1), MilliMeters.Value(914.4), false},
		{Yard.Value(1760), Miles.Value(1), false},
		// =========================
		// Miles
		// =========================
		{Miles.Value(1), Kilometers.Value(1.609344), false},
		{Miles.Value(1), Meters.Value(1609.344), false},
		{Miles.Value(1), CentiMeters.Value(160934.4), false},
		{Miles.Value(1), MilliMeters.Value(1609344), false},
	})
}
