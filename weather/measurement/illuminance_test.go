package measurement

import "testing"

func Test_illuminance(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// Lux
		// =========================
		newConversionTest(Lux.Value(1), FootCandles.Value(0.0929031299064), false),
		newConversionTest(Lux.Value(1), KiloFootCandles.Value(0.0000929031299064), false),
		newConversionTest(Lux.Value(1), KiloLux.Value(0.001), false),
		newConversionTest(Lux.Value(1), WattsPerSquareMeter.Value(0.0079), false),
		newConversionTest(Lux.Value(1), KiloWattsPerSquareMeter.Value(0.0000079), false),
		// =========================
		// KiloLux
		// =========================
		newConversionTest(KiloLux.Value(1), FootCandles.Value(92.9031299064), false),
		newConversionTest(KiloLux.Value(1), KiloFootCandles.Value(0.0929031299064), false),
		newConversionTest(KiloLux.Value(1), Lux.Value(1000), false),
		newConversionTest(KiloLux.Value(1), WattsPerSquareMeter.Value(7.9), false),
		newConversionTest(KiloLux.Value(1), KiloWattsPerSquareMeter.Value(0.0079), false),
		// =========================
		// FootCandles
		// =========================
		newConversionTest(FootCandles.Value(1), Lux.Value(10.7639), false),
		newConversionTest(FootCandles.Value(1), KiloFootCandles.Value(.001), false),
		newConversionTest(FootCandles.Value(1), WattsPerSquareMeter.Value(0.08503481), false),
		newConversionTest(FootCandles.Value(1), KiloWattsPerSquareMeter.Value(0.00008503481), false),
		// =========================
		// KiloFootCandles
		// =========================
		newConversionTest(KiloFootCandles.Value(1), Lux.Value(10763.9), false),
		newConversionTest(KiloFootCandles.Value(1), FootCandles.Value(1000), false),
		// =========================
		// WattsPerSquareMeter
		// =========================
		newConversionTest(WattsPerSquareMeter.Value(1), Lux.Value(126.582278481), false),
		newConversionTest(WattsPerSquareMeter.Value(1), KiloWattsPerSquareMeter.Value(.001), false),
		newConversionTest(WattsPerSquareMeter.Value(1), FootCandles.Value(11.759889), false),
		newConversionTest(WattsPerSquareMeter.Value(1), KiloFootCandles.Value(0.011759), false),
		// =========================
		// KiloWattsPerSquareMeter
		// =========================
		newConversionTest(KiloWattsPerSquareMeter.Value(1), Lux.Value(126582.278481), false),
		newConversionTest(KiloWattsPerSquareMeter.Value(1), WattsPerSquareMeter.Value(1000), false),
		newConversionTest(KiloWattsPerSquareMeter.Value(1), FootCandles.Value(11759.8899), false),
		newConversionTest(KiloWattsPerSquareMeter.Value(1), KiloFootCandles.Value(11.7598899), false),
	})
}
