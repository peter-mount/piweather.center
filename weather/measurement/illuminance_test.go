package measurement

import "testing"

func Test_illuminance(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// Lux
		// =========================
		{Lux.Value(1), FootCandles.Value(0.0929031299064), false},
		{Lux.Value(1), KiloFootCandles.Value(0.0000929031299064), false},
		{Lux.Value(1), KiloLux.Value(0.001), false},
		{Lux.Value(1), WattsPerSquareMeter.Value(0.0079), false},
		{Lux.Value(1), KiloWattsPerSquareMeter.Value(0.0000079), false},
		// =========================
		// KiloLux
		// =========================
		{KiloLux.Value(1), FootCandles.Value(92.9031299064), false},
		{KiloLux.Value(1), KiloFootCandles.Value(0.0929031299064), false},
		{KiloLux.Value(1), Lux.Value(1000), false},
		{KiloLux.Value(1), WattsPerSquareMeter.Value(7.9), false},
		{KiloLux.Value(1), KiloWattsPerSquareMeter.Value(0.0079), false},
		// =========================
		// FootCandles
		// =========================
		{FootCandles.Value(1), Lux.Value(10.7639), false},
		{FootCandles.Value(1), KiloFootCandles.Value(.001), false},
		{FootCandles.Value(1), WattsPerSquareMeter.Value(0.08503481), false},
		{FootCandles.Value(1), KiloWattsPerSquareMeter.Value(0.00008503481), false},
		// =========================
		// KiloFootCandles
		// =========================
		{KiloFootCandles.Value(1), Lux.Value(10763.9), false},
		{KiloFootCandles.Value(1), FootCandles.Value(1000), false},
		// =========================
		// WattsPerSquareMeter
		// =========================
		{WattsPerSquareMeter.Value(1), Lux.Value(126.582278481), false},
		{WattsPerSquareMeter.Value(1), KiloWattsPerSquareMeter.Value(.001), false},
		{WattsPerSquareMeter.Value(1), FootCandles.Value(11.759889), false},
		{WattsPerSquareMeter.Value(1), KiloFootCandles.Value(0.011759), false},
		// =========================
		// KiloWattsPerSquareMeter
		// =========================
		{KiloWattsPerSquareMeter.Value(1), Lux.Value(126582.278481), false},
		{KiloWattsPerSquareMeter.Value(1), WattsPerSquareMeter.Value(1000), false},
		{KiloWattsPerSquareMeter.Value(1), FootCandles.Value(11759.8899), false},
		{KiloWattsPerSquareMeter.Value(1), KiloFootCandles.Value(11.7598899), false},
	})
}
