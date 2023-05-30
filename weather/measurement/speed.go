package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	MetersPerSecond = value.NewLowerBoundUnit("MetersPerSecond", "Meters Per Second", " m/s", 1, 0)
	KilometersPerHour = value.NewLowerBoundUnit("KilometersPerHour", "Kilometers Per Hour", " km/h", 1, 0)
	MilesPerHour = value.NewLowerBoundUnit("MilesPerHour", "Miles Per Hour", " mph", 1, 0)
	FeetPerSecond = value.NewLowerBoundUnit("FeetPerSecond", "Feet Per Second", " ft/s", 1, 0)
	Knots = value.NewLowerBoundUnit("Knots", "Knots", " kn", 1, 0)
	BeaufortScale = value.NewBoundedUnit("BeaufortScale", "Beaufort Scale", "", 0, 0, 12)

	// Transforms between mps and each unit - this registers both directions
	value.NewBasicBiTransform(MetersPerSecond, KilometersPerHour, mpsToKph)
	value.NewBasicBiTransform(MetersPerSecond, MilesPerHour, mpsToMph)
	value.NewBasicBiTransform(MetersPerSecond, FeetPerSecond, feetPerSecond)
	value.NewBasicBiTransform(MetersPerSecond, Knots, mpsToKnots)
	value.NewBiTransform(MetersPerSecond, BeaufortScale, mpsToBeaufort, beaufortToMps)

	// Conversions between units other than MetersPerSecond. These convert to MetersPerSecond first then to the final one
	Speed = value.NewGroup("Speed", MetersPerSecond, FeetPerSecond, KilometersPerHour, Knots, MilesPerHour, BeaufortScale)
}

var (
	Speed             *value.Group
	FeetPerSecond     *value.Unit
	KilometersPerHour *value.Unit
	Knots             *value.Unit
	MetersPerSecond   *value.Unit
	MilesPerHour      *value.Unit

	// BeaufortScale Beaufort wind force scale https://www.metoffice.gov.uk/weather/guides/coast-and-sea/beaufort-scale
	BeaufortScale *value.Unit

	// Max wind speed in m/s for each beaufort scale 0..11. 12 is not included as it's unbounded
	beaufortMaxSpeed = []float64{1, 2, 4, 6, 9, 11, 14, 17, 21, 25, 29, 33}
	// Mean Wind Speed in m/s for each beaufort scale. Last entry is undefined so we return the lower limit
	beaufortSpeed = []float64{0, 1, 3, 5, 7, 10, 12, 15, 19, 23, 27, 31, 33}
)

// mpsToBeaufort returns the beaufort scale number based on a series of max speeds in m/s.
// Hence, anything under 1m/s is 0, <2 is 1 etc.
// If none of the entries matched then the length of the array is returned, representing
// beaufort scale 12 - or a Hurricane
func mpsToBeaufort(f float64) (float64, error) {
	for i, max := range beaufortMaxSpeed {
		if value.LessThan(f, max) {
			return float64(i), nil
		}
	}
	return 12, nil
}

func beaufortToMps(f float64) (float64, error) {
	bf := int(f)
	if bf < 0 {
		bf = 0
	}
	if bf >= len(beaufortSpeed) {
		bf = len(beaufortSpeed) - 1
	}
	return beaufortSpeed[bf], nil
}
