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

	// This is an approximate conversion of lux to W/m2 based on average wavelength of sunlight
	wm2ToLux = 0.0079

	// Mass conversion constants:
	poundToGram = 453.59237

	// Pressure conversion constants:
	standardGravity = 9.80665
	mercuryDensity  = 13.5951

	// Volume conversion constants:
	cubicFootToCubicMeter = footToM * footToM * footToM
)
