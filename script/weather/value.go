package weather

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"time"
)

func init() {
	packages.Register("value", &Value{})
}

type Value struct{}

func (_ Value) PlainTime(t time.Time) value.Time {
	return value.PlainTime(t)
}

func (_ Value) BasicTime(t time.Time, loc *globe.Coord, alt float64) value.Time {
	return value.BasicTime(t, loc, alt)
}

func (_ Value) Radian() *value.Unit { return measurement.Radian }

func (_ Value) Degree() *value.Unit { return measurement.Degree }

func (_ Value) ArcMinute() *value.Unit { return measurement.ArcMinute }

func (_ Value) ArcSecond() *value.Unit { return measurement.ArcSecond }

func (_ Value) Gradian() *value.Unit { return measurement.Gradian }

func (_ Value) HourAngle() *value.Unit { return measurement.HourAngle }

func (_ Value) Turn() *value.Unit { return measurement.Turn }

func (_ Value) PartsPerMillion() *value.Unit { return measurement.PartsPerMillion }

func (_ Value) MicrogramsPerCubicMeter() *value.Unit { return measurement.MicrogramsPerCubicMeter }

func (_ Value) GramsPerCubicMeter() *value.Unit { return measurement.GramsPerCubicMeter }

func (_ Value) PoundsPerCubitFoot() *value.Unit { return measurement.PoundsPerCubitFoot }

func (_ Value) RelativeHumidity() *value.Unit { return measurement.RelativeHumidity }

func (_ Value) Lux() *value.Unit { return measurement.Lux }

func (_ Value) FootCandles() *value.Unit { return measurement.FootCandles }

func (_ Value) KiloFootCandles() *value.Unit { return measurement.KiloFootCandles }

func (_ Value) KiloLux() *value.Unit { return measurement.KiloLux }

func (_ Value) WattsPerSquareMeter() *value.Unit { return measurement.WattsPerSquareMeter }

func (_ Value) KiloWattsPerSquareMeter() *value.Unit { return measurement.KiloWattsPerSquareMeter }

func (_ Value) Meters() *value.Unit { return measurement.Meters }

func (_ Value) Kilometers() *value.Unit { return measurement.Kilometers }

func (_ Value) CentiMeters() *value.Unit { return measurement.CentiMeters }

func (_ Value) MilliMeters() *value.Unit { return measurement.MilliMeters }

func (_ Value) Inches() *value.Unit { return measurement.Inches }

func (_ Value) Feet() *value.Unit { return measurement.Feet }

func (_ Value) Yard() *value.Unit { return measurement.Yard }

func (_ Value) Miles() *value.Unit { return measurement.Miles }

func (_ Value) MillimetersPerHour() *value.Unit { return measurement.MillimetersPerHour }

func (_ Value) InchesPerHour() *value.Unit { return measurement.InchesPerHour }

func (_ Value) PressurePA() *value.Unit { return measurement.PressurePA }

func (_ Value) PressureHPA() *value.Unit { return measurement.PressureHPA }

func (_ Value) PressureKPA() *value.Unit { return measurement.PressureKPA }

func (_ Value) PressurePSI() *value.Unit { return measurement.PressurePSI }

func (_ Value) PressureInHg() *value.Unit { return measurement.PressureInHg }

func (_ Value) PressureMmHg() *value.Unit { return measurement.PressureMmHg }

func (_ Value) PressureBar() *value.Unit { return measurement.PressureBar }

func (_ Value) PressureCBar() *value.Unit { return measurement.PressureCBar }

func (_ Value) PressureMBar() *value.Unit { return measurement.PressureMBar }

func (_ Value) CountPerMinute() *value.Unit { return measurement.CountPerMinute }

func (_ Value) MetersPerSecond() *value.Unit { return measurement.MetersPerSecond }

func (_ Value) KilometersPerHour() *value.Unit { return measurement.KilometersPerHour }

func (_ Value) MilesPerHour() *value.Unit { return measurement.MilesPerHour }

func (_ Value) FeetPerSecond() *value.Unit { return measurement.FeetPerSecond }

func (_ Value) Knots() *value.Unit { return measurement.Knots }

func (_ Value) BeaufortScale() *value.Unit { return measurement.BeaufortScale }

func (_ Value) Celsius() *value.Unit { return measurement.Celsius }

func (_ Value) Fahrenheit() *value.Unit { return measurement.Fahrenheit }

func (_ Value) Kelvin() *value.Unit { return measurement.Kelvin }

func (_ Value) ModifiedJD() *value.Unit { return measurement.ModifiedJD }

func (_ Value) JulianDate() *value.Unit { return measurement.JulianDate }

func (_ Value) MarsSolDate() *value.Unit { return measurement.MarsSolDate }

func (_ Value) RataDie() *value.Unit { return measurement.RataDie }

func (_ Value) Volt() *value.Unit { return measurement.Volt }

func (_ Value) MilliVolt() *value.Unit { return measurement.MilliVolt }

func (_ Value) MicroVolt() *value.Unit { return measurement.MicroVolt }

func (_ Value) DecibelVolt() *value.Unit { return measurement.DecibelVolt }

func (_ Value) DecibelVoltU() *value.Unit { return measurement.DecibelVoltU }
