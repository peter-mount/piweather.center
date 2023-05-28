package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Lux = value.NewLowerBoundUnit("Lux", "Lux", " lx", 1, 0)
	FootCandles = value.NewLowerBoundUnit("FootCandles", "Foot Candles", " fc", 1, 0)
	KiloFootCandles = value.NewLowerBoundUnit("KiloFootCandles", "Kilo Foot Candles", " kfc", 1, 0)
	KiloLux = value.NewLowerBoundUnit("KiloLux", "KiloLux", " klx", 1, 0)
	WattsPerSquareMeter = value.NewLowerBoundUnit("WattsPerSquareMeter", "Watts Per Square Meter", " W/m²", 1, 0)

	// Transforms from base unit Lux
	value.NewBasicBiTransform(Lux, FootCandles, 1.0/fcToLux)
	value.NewBasicBiTransform(Lux, KiloFootCandles, 1.0/fcToLux/1000.0)
	value.NewBasicBiTransform(Lux, KiloLux, 1.0/kluxToLux)
	value.NewBasicBiTransform(Lux, WattsPerSquareMeter, wm2ToLux)

	Illuminance = value.NewGroup("Illuminance", Lux, FootCandles, KiloFootCandles, KiloLux, WattsPerSquareMeter)
}

var (
	Illuminance         *value.Group
	Lux                 *value.Unit
	FootCandles         *value.Unit
	KiloFootCandles     *value.Unit
	KiloLux             *value.Unit
	WattsPerSquareMeter *value.Unit
)
