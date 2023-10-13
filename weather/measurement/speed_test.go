package measurement

import (
	"testing"
)

func Test_speed(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// MetersPerSecond
		// =========================
		{MetersPerSecond.Value(1), KilometersPerHour.Value(3.6), false},
		{MetersPerSecond.Value(1), Knots.Value(1.9438444), false},
		{MetersPerSecond.Value(1), MilesPerHour.Value(2.236936), false},
		// =========================
		// KilometersPerHour
		// =========================
		{KilometersPerHour.Value(1), MetersPerSecond.Value(0.27777777), false},
		{KilometersPerHour.Value(1), Knots.Value(0.539956), false},
		{KilometersPerHour.Value(1), MilesPerHour.Value(0.62137111), false},
		// =========================
		// Knots
		// =========================
		{Knots.Value(1), MetersPerSecond.Value(0.5144444444), false},
		{Knots.Value(1), MilesPerHour.Value(1.1507794480), false},
		{Knots.Value(1), KilometersPerHour.Value(1.852), false},
		// =========================
		// MilesPerHour
		// =========================
		{MilesPerHour.Value(1), MetersPerSecond.Value(0.4470400), false},
		{MilesPerHour.Value(1), KilometersPerHour.Value(1.609344), false},
		{MilesPerHour.Value(1), Knots.Value(0.8690), false},
		{MilesPerHour.Value(1), FeetPerSecond.Value(1.4667), false},
		// =========================
		// FeetPerSecond
		// =========================
		{FeetPerSecond.Value(1), MetersPerSecond.Value(0.30488), false},
		{FeetPerSecond.Value(1), KilometersPerHour.Value(1.0973), false},
		{FeetPerSecond.Value(1), Knots.Value(0.5924839), false},
		{FeetPerSecond.Value(1), MilesPerHour.Value(0.6818185), false},
	})
}
