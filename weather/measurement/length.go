package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Meters = value.NewLowerBoundUnit("Meters", "Meters", " m", 3, 0)
	Kilometers = value.NewLowerBoundUnit("Kilometers", "Kilometers", " km", 3, 0)
	CentiMeters = value.NewLowerBoundUnit("CentiMeters", "CentiMeters", " cm", 3, 0)
	MilliMeters = value.NewLowerBoundUnit("MilliMeters", "MilliMeters", " mm", 3, 0)
	Inches = value.NewLowerBoundUnit("Inches", "Inches", " in", 3, 0)
	Feet = value.NewLowerBoundUnit("Feet", "Feet", " ft", 3, 0)
	Yard = value.NewLowerBoundUnit("Yard", "Yard", " yd", 3, 0)
	Miles = value.NewLowerBoundUnit("Miles", "Miles", " mi", 3, 0)

	value.NewBasicBiTransform(Meters, Kilometers, 1/kmToM)
	value.NewBasicBiTransform(Meters, CentiMeters, 1/cmToM)
	value.NewBasicBiTransform(Meters, MilliMeters, 1/mmToM)
	value.NewBasicBiTransform(Meters, Inches, 1/inToM)
	value.NewBasicBiTransform(Meters, Feet, 1/footToM)
	value.NewBasicBiTransform(Meters, Yard, 1/yardToM)
	value.NewBasicBiTransform(Meters, Miles, 1/mileToM)

	Length = value.NewGroup("Length", Meters, Kilometers, CentiMeters, MilliMeters, Inches, Feet, Yard, Miles)
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
)
