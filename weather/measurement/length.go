package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Meters = value.NewLowerBoundUnit("Meters", "Meters", " m", 1, 0)
	Kilometers = value.NewLowerBoundUnit("Kilometers", "Kilometers", " km", 1, 0)
	CentiMeters = value.NewLowerBoundUnit("CentiMeters", "CentiMeters", " cm", 1, 0)
	MilliMeters = value.NewLowerBoundUnit("MilliMeters", "MilliMeters", " mm", 1, 0)
	Inches = value.NewLowerBoundUnit("Inches", "Inches", " in", 3, 0)
	Feet = value.NewLowerBoundUnit("Feet", "Feet", " ft", 3, 0)
	Yard = value.NewLowerBoundUnit("Yard", "Yard", " yd", 3, 0)
	Miles = value.NewLowerBoundUnit("Miles", "Miles", " mi", 3, 0)
	AU = value.NewLowerBoundUnit("AU", "Astronomical Unit", " au", 3, 0)
	LightYear = value.NewLowerBoundUnit("LightYear", "Lightyear", " ly", 3, 0)

	// Base unit is Meters but as our constants are all ToM then use Meters as the destination
	value.NewBasicBiTransform(Kilometers, Meters, kmToM)
	value.NewBasicBiTransform(CentiMeters, Meters, cmToM)
	value.NewBasicBiTransform(MilliMeters, Meters, mmToM)
	value.NewBasicBiTransform(Inches, Meters, inToM)
	value.NewBasicBiTransform(Feet, Meters, footToM)
	value.NewBasicBiTransform(Yard, Meters, yardToM)
	value.NewBasicBiTransform(Miles, Meters, mileToM)
	value.NewBasicBiTransform(AU, Meters, auToM)
	value.NewBasicBiTransform(LightYear, Meters, lyToKm*kmToM)

	value.NewBasicBiTransform(LightYear, Kilometers, lyToKm)

	Length = value.NewGroup("Length", Meters, Kilometers, CentiMeters, MilliMeters, Inches, Feet, Yard, Miles, AU, LightYear)
}

var (
	Length      *value.Group
	Meters      *value.Unit
	Kilometers  *value.Unit
	CentiMeters *value.Unit
	MilliMeters *value.Unit
	Inches      *value.Unit
	Feet        *value.Unit
	Yard        *value.Unit
	Miles       *value.Unit
	AU          *value.Unit
	LightYear   *value.Unit
)
