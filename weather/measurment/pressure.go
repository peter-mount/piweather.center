package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	PressurePA = value.NewLowerBoundUnit("PressurePA", "Pressure", "Pressure Pascals", " Pa", 1, 0)
	PressureHPA = value.NewLowerBoundUnit("PressureHPA", "Pressure", "Pressure HPascals", " hPa", 1, 0)
	PressurePSI = value.NewLowerBoundUnit("PressurePSI", "Pressure", "Pressure Pounds per Square Inch", " psi", 1, 0)
	PressureInHg = value.NewLowerBoundUnit("PressureInHg", "Pressure", "Pressure Inches Mercury", " inHg", 1, 0)
	PressureMmHg = value.NewLowerBoundUnit("PressureMmHg", "Pressure", "Pressure mm Mercury", " mmHg", 1, 0)
	PressureBar = value.NewLowerBoundUnit("PressureBar", "Pressure", "Pressure Bar", " bar", 1, 0)
	PressureCBar = value.NewLowerBoundUnit("PressureCBar", "Pressure", "Pressure CentiBar", " cbar", 1, 0)
	PressureMBar = value.NewLowerBoundUnit("PressureMBar", "Pressure", "Pressure MilliBar", " mbar", 1, 0)

	// Transforms from base unit PressurePA
	value.NewBasicBiTransform(PressurePA, PressureHPA, 1.0/1000.0)
	value.NewBasicBiTransform(PressurePA, PressurePSI, 1.0/6894.757)
	value.NewBasicBiTransform(PressurePA, PressureInHg, 1.0/(inToM*1000.0*standardGravity*mercuryDensity))
	value.NewBasicBiTransform(PressurePA, PressureMmHg, 1.0/(mmToM*1000.0*standardGravity*mercuryDensity))
	value.NewBasicBiTransform(PressurePA, PressureBar, 1.0/100000.0)
	value.NewBasicBiTransform(PressurePA, PressureCBar, 1.0/1000.0)
	value.NewBasicBiTransform(PressurePA, PressureMBar, 1.0/100.0)

	// Conversions between units other than PressurePA. These convert to PressurePA first then to the final one
	value.NewTransformations(PressurePA, PressureHPA, PressurePSI, PressureInHg, PressureMmHg, PressureBar, PressureCBar, PressureMBar)
}

var (
	PressurePA   *value.Unit
	PressureHPA  *value.Unit
	PressurePSI  *value.Unit
	PressureInHg *value.Unit
	PressureMmHg *value.Unit
	PressureBar  *value.Unit
	PressureCBar *value.Unit
	PressureMBar *value.Unit
)
