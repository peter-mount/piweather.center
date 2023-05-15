package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	PressurePA = value.NewLowerBoundUnit("Pressure Pascals", " Pa", value.Dp1, 0)
	PressureHPA = value.NewLowerBoundUnit("Pressure HPascals", " hPa", value.Dp1, 0)
	PressurePSI = value.NewLowerBoundUnit("Pressure Pounds per Square Inch", " psi", value.Dp1, 0)
	PressureInHg = value.NewLowerBoundUnit("Pressure Inches Mercury", " inHg", value.Dp1, 0)
	PressureMmHg = value.NewLowerBoundUnit("Pressure mm Mercury", " mmHg", value.Dp1, 0)
	PressureBar = value.NewLowerBoundUnit("Pressure Bar", " bar", value.Dp1, 0)
	PressureCBar = value.NewLowerBoundUnit("Pressure CentiBar", " cbar", value.Dp1, 0)
	PressureMBar = value.NewLowerBoundUnit("Pressure MilliBar", " mbar", value.Dp1, 0)

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
	PressurePA   value.Unit
	PressureHPA  value.Unit
	PressurePSI  value.Unit
	PressureInHg value.Unit
	PressureMmHg value.Unit
	PressureBar  value.Unit
	PressureCBar value.Unit
	PressureMBar value.Unit
)
