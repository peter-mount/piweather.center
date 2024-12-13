package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	// Generic Voltage
	Volt = value.NewUnit("Volt", "Volt", "V", 3, nil)
	MilliVolt = value.NewUnit("MilliVolt", "MilliVolt", " mV", 0, nil)
	MicroVolt = value.NewUnit("MicroVolt", "MicroVolt", " µV", 0, nil)
	DecibelVolt = value.NewUnit("DecibelVolt", "Voltage dBV", " dBV", 1, nil)
	DecibelVoltU = value.NewUnit("DecibelVoltU", "Voltage dBu", " dBu", 1, nil)

	value.NewBasicBiTransform(MilliVolt, Volt, Milli)
	value.NewBasicBiTransform(MicroVolt, Volt, Micro)

	value.NewTransform(Volt, DecibelVolt, vRmsToDbv)
	value.NewTransform(DecibelVolt, Volt, dBvToVrms)

	value.NewTransform(Volt, DecibelVoltU, vRmsToDbu)
	value.NewTransform(DecibelVoltU, Volt, dBuToVrms)

	Voltage = value.NewGroup("Voltage", Volt, MilliVolt, MicroVolt, DecibelVolt, DecibelVoltU)
}

var (
	// Voltage
	Voltage   *value.Group
	Volt      *value.Unit
	MilliVolt *value.Unit
	MicroVolt *value.Unit
	// DecibelVolt Decibel Volt referenced to 1V
	DecibelVolt *value.Unit
	// DecibelVoltU Decibel Volt referenced to 0.775V
	DecibelVoltU *value.Unit
)

// VdBV to Vrms - 10 ^ (Vdbv / 20.0)
//
// src: https://www.everythingrf.com/community/what-is-dbv
func dBvToVrms(dBv float64) (float64, error) {
	return math.Pow(10.0, dBv/20.0), nil
}

// V(dBV) = 20 * log10 ( V(rms) / 1V)
//
// src: https://www.everythingrf.com/community/what-is-dbv
func vRmsToDbv(v float64) (float64, error) {
	return 20.0 * math.Log10(v/1.0), nil
}

func vRmsToDbu(v float64) (float64, error) {
	return 20.0 * math.Log10(v/0.775), nil
}

func dBuToVrms(dBv float64) (float64, error) {
	return math.Pow(10.0, dBv/20.0) * 0.775, nil
}
