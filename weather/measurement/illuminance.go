package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Lux = value.NewLowerBoundUnit("Lux", "Lux", " lx", 1, 0)
	FootCandles = value.NewLowerBoundUnit("FootCandles", "Foot Candles", " fc", 1, 0)
	KiloFootCandles = value.NewLowerBoundUnit("KiloFootCandles", "Kilo Foot Candles", " kfc", 1, 0)
	KiloLux = value.NewLowerBoundUnit("KiloLux", "KiloLux", " klx", 1, 0)
	WattsPerSquareMeter = value.NewLowerBoundUnit("WattsPerSquareMeter", "Watts Per Square Meter", " W/m²", 1, 0)
	KiloWattsPerSquareMeter = value.NewLowerBoundUnit("KiloWattsPerSquareMeter", "KiloWatts Per Square Meter", " kW/m²", 2, 0)

	// Transforms from base unit Lux
	value.NewBasicBiTransform(FootCandles, Lux, fcToLux)
	value.NewBasicBiTransform(KiloFootCandles, Lux, fcToLux*1000.0)
	value.NewBasicBiTransform(KiloLux, Lux, kluxToLux)
	value.NewBasicBiTransform(Lux, WattsPerSquareMeter, wm2ToLux)
	value.NewBasicBiTransform(Lux, KiloWattsPerSquareMeter, wm2ToLux/Kilo)

	// W/m² -> kW/m² for speed
	value.NewBasicBiTransform(KiloWattsPerSquareMeter, WattsPerSquareMeter, Kilo)

	Illuminance = value.NewGroup("Illuminance", Lux, FootCandles, KiloFootCandles, KiloLux, WattsPerSquareMeter, KiloWattsPerSquareMeter)
}

var (
	Illuminance             *value.Group
	Lux                     *value.Unit
	FootCandles             *value.Unit
	KiloFootCandles         *value.Unit
	KiloLux                 *value.Unit
	WattsPerSquareMeter     *value.Unit
	KiloWattsPerSquareMeter *value.Unit
)
