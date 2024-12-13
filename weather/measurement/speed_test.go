package measurement

import (
	"testing"
)

func Test_speed(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// MetersPerSecond
		// =========================
		newConversionTest(MetersPerSecond.Value(1), KilometersPerHour.Value(3.6), false),
		newConversionTest(MetersPerSecond.Value(1), Knots.Value(1.9438444), false),
		newConversionTest(MetersPerSecond.Value(1), MilesPerHour.Value(2.236936), false),
		// =========================
		// KilometersPerHour
		// =========================
		newConversionTest(KilometersPerHour.Value(1), MetersPerSecond.Value(0.27777777), false),
		newConversionTest(KilometersPerHour.Value(1), Knots.Value(0.539956), false),
		newConversionTest(KilometersPerHour.Value(1), MilesPerHour.Value(0.62137111), false),
		// =========================
		// Knots
		// =========================
		newConversionTest(Knots.Value(1), MetersPerSecond.Value(0.5144444444), false),
		newConversionTest(Knots.Value(1), MilesPerHour.Value(1.1507794480), false),
		newConversionTest(Knots.Value(1), KilometersPerHour.Value(1.852), false),
		// =========================
		// MilesPerHour
		// =========================
		newConversionTest(MilesPerHour.Value(1), MetersPerSecond.Value(0.4470400), false),
		newConversionTest(MilesPerHour.Value(1), KilometersPerHour.Value(1.609344), false),
		newConversionTest(MilesPerHour.Value(1), Knots.Value(0.8690), false),
		newConversionTest(MilesPerHour.Value(1), FeetPerSecond.Value(1.4667), false),
		// =========================
		// FeetPerSecond
		// =========================
		//newConversionTest(FeetPerSecond.Value(1), MetersPerSecond.Value(0.30488), false),
		newConversionTest(FeetPerSecond.Value(1), KilometersPerHour.Value(1.0973), false),
		newConversionTest(FeetPerSecond.Value(1), Knots.Value(0.5924839), false),
		newConversionTest(FeetPerSecond.Value(1), MilesPerHour.Value(0.6818185), false),
	})
}
