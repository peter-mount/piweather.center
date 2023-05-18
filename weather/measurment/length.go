package measurment

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Meters = value.NewLowerBoundUnit("Meters", "Length", "Meters", " m", 3, 0)
	Kilometers = value.NewLowerBoundUnit("Kilometers", "Length", "Kilometers", " km", 3, 0)
	CentiMeters = value.NewLowerBoundUnit("CentiMeters", "Length", "CentiMeters", " cm", 3, 0)
	MilliMeters = value.NewLowerBoundUnit("MilliMeters", "Length", "MilliMeters", " mm", 3, 0)
	Inches = value.NewLowerBoundUnit("Inches", "Length", "Inches", " in", 3, 0)
	Feet = value.NewLowerBoundUnit("Feet", "Length", "Feet", " ft", 3, 0)
	Yard = value.NewLowerBoundUnit("Yard", "Length", "Yard", " yd", 3, 0)
	Miles = value.NewLowerBoundUnit("Miles", "Length", "Miles", " mi", 3, 0)

	value.NewBasicBiTransform(Meters, Kilometers, 1/kmToM)
	value.NewBasicBiTransform(Meters, CentiMeters, 1/cmToM)
	value.NewBasicBiTransform(Meters, MilliMeters, 1/mmToM)
	value.NewBasicBiTransform(Meters, Inches, 1/inToM)
	value.NewBasicBiTransform(Meters, Feet, 1/footToM)
	value.NewBasicBiTransform(Meters, Yard, 1/yardToM)
	value.NewBasicBiTransform(Meters, Miles, 1/mileToM)

	value.NewTransformations(Meters, Kilometers, CentiMeters, MilliMeters, Inches, Feet, Yard, Miles)
}

var (
	Meters      *value.Unit
	Kilometers  *value.Unit
	CentiMeters *value.Unit
	MilliMeters *value.Unit
	Inches      *value.Unit
	Feet        *value.Unit
	Yard        *value.Unit
	Miles       *value.Unit
)
