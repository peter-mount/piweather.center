package measurement

const (
	// Distance conversion constants:
	cmToM           = 0.01
	inToM           = 0.0254
	kmToM           = 1000.0
	mmToM           = 0.001
	nauticalMileToM = 1852.0
	footToM         = inToM * 12.0
	yardToM         = footToM * 3.0
	mileToM         = yardToM * 1760.0
	feetPerSecond   = 1.0 / footToM
	mpsToKph        = hrsToSec / kmToM
	mpsToMph        = hrsToSec / mileToM
	mpsToKnots      = hrsToSec / nauticalMileToM

	// Duration conversion constants:
	hrsToSec = 60.0 * 60.0
	dayToSec = hrsToSec * 24.0

	// Illuminance conversion constants:
	kluxToLux = 1000.0
	fcToLux   = 10.7639
	wm2ToLux  = 0.0079

	// Mass conversion constants:
	poundToGram = 453.59237

	// Pressure conversion constants:
	standardGravity = 9.80665
	mercuryDensity  = 13.5951

	// Volume conversion constants:
	cubicFootToCubicMeter = footToM * footToM * footToM

	Quetta = 1e30
	Ronna  = 1e27
	Yotta  = 1e24
	Zetta  = 1e21
	Exa    = 1e18
	Peta   = 1e15
	Tera   = 1e12
	Giga   = 1e9
	Mega   = 1e6
	Kilo   = 1e3
	Hecto  = 1e2
	Deca   = 1e1
	Deci   = 1e-1
	Centi  = 1e-2
	Milli  = 1e-3
	Micro  = 1e-6
	Nano   = 1e-9
	Pico   = 1e-12
	Femto  = 1e-15
	Atto   = 1e-18
	Zepto  = 1e-21
	Yocto  = 1e-24
	Ronto  = 1e-27
	Quecto = 1e-30
)
