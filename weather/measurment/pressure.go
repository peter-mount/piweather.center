package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	PressurePA = value.NewLowerBoundUnit("PressurePA", "Pressure Pascals", " Pa", 1, 0)
	PressureHPA = value.NewLowerBoundUnit("PressureHPA", "Pressure Hecto Pascals", " hPa", 1, 0)
	PressureKPA = value.NewLowerBoundUnit("PressureKPA", "Pressure Kilo Pascals", " kPa", 1, 0)
	PressurePSI = value.NewLowerBoundUnit("PressurePSI", "Pressure Pounds per Square Inch", " psi", 1, 0)
	PressureInHg = value.NewLowerBoundUnit("PressureInHg", "Pressure Inches Mercury", " inHg", 1, 0)
	PressureMmHg = value.NewLowerBoundUnit("PressureMmHg", "Pressure mm Mercury", " mmHg", 1, 0)
	PressureBar = value.NewLowerBoundUnit("PressureBar", "Pressure Bar", " bar", 1, 0)
	PressureCBar = value.NewLowerBoundUnit("PressureCBar", "Pressure CentiBar", " cbar", 1, 0)
	PressureMBar = value.NewLowerBoundUnit("PressureMBar", "Pressure MilliBar", " mbar", 1, 0)

	// Transforms from base unit PressurePA
	value.NewBasicBiTransform(PressurePA, PressureHPA, 1.0/100.0)
	value.NewBasicBiTransform(PressurePA, PressureKPA, 1.0/1000.0)
	value.NewBasicBiTransform(PressurePA, PressurePSI, 1.0/6894.757)
	value.NewBasicBiTransform(PressurePA, PressureInHg, 1.0/(inToM*1000.0*standardGravity*mercuryDensity))
	value.NewBasicBiTransform(PressurePA, PressureMmHg, 1.0/(mmToM*1000.0*standardGravity*mercuryDensity))
	value.NewBasicBiTransform(PressurePA, PressureBar, 1.0/100000.0)
	value.NewBasicBiTransform(PressurePA, PressureCBar, 1.0/1000.0)
	value.NewBasicBiTransform(PressurePA, PressureMBar, 1.0/100.0)

	// Conversions between units other than PressurePA. These convert to PressurePA first then to the final one
	Pressure = value.NewGroup("Pressure", PressurePA, PressureHPA, PressureKPA, PressurePSI, PressureInHg, PressureMmHg, PressureBar, PressureCBar, PressureMBar)
}

var (
	Pressure     *value.Group
	PressurePA   *value.Unit
	PressureHPA  *value.Unit
	PressureKPA  *value.Unit
	PressurePSI  *value.Unit
	PressureInHg *value.Unit
	PressureMmHg *value.Unit
	PressureBar  *value.Unit
	PressureCBar *value.Unit
	PressureMBar *value.Unit
)
