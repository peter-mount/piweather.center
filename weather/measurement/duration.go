package measurement

import "github.com/peter-mount/piweather.center/weather/value"

var (
	// Duration a group of units that represent a duration in time
	Duration         *value.Group
	DurationSecond   *value.Unit
	DurationMinute   *value.Unit
	DurationHour     *value.Unit
	DurationDay      *value.Unit
	DurationYear     *value.Unit
	DurationKiloYear *value.Unit
	DurationMegaYear *value.Unit
	DurationGigaYear *value.Unit
)

func init() {
	DurationSecond = value.NewLowerBoundUnit("durationSecond", "Second", " s", 1, 0)
	DurationMinute = value.NewLowerBoundUnit("durationMinute", "Minute", " m", 2, 0)
	DurationHour = value.NewLowerBoundUnit("durationHour", "Hour", " h", 3, 0)
	DurationDay = value.NewLowerBoundUnit("durationDay", "Day", " d", 6, 0)
	DurationYear = value.NewLowerBoundUnit("durationYear", "Year", " y", 6, 0)
	DurationKiloYear = value.NewLowerBoundUnit("durationKYear", "KiloYear", " ky", 6, 0)
	DurationMegaYear = value.NewLowerBoundUnit("durationMYear", "MegaYear", " my", 6, 0)
	DurationGigaYear = value.NewLowerBoundUnit("durationGYear", "GigaYear", " gy", 6, 0)

	value.NewBasicBiTransform(DurationMinute, DurationSecond, durationMinute)
	value.NewBasicBiTransform(DurationHour, DurationSecond, durationHour)
	value.NewBasicBiTransform(DurationDay, DurationSecond, durationDay)
	value.NewBasicBiTransform(DurationYear, DurationSecond, durationYear)
	value.NewBasicBiTransform(DurationKiloYear, DurationSecond, durationYear*1000)
	value.NewBasicBiTransform(DurationMegaYear, DurationSecond, durationYear*1000000)
	value.NewBasicBiTransform(DurationGigaYear, DurationSecond, durationYear*1000000000)

	value.NewBasicBiTransform(DurationHour, DurationMinute, 60)
	value.NewBasicBiTransform(DurationDay, DurationHour, 24)
	value.NewBasicBiTransform(DurationYear, DurationDay, 365.25)
	value.NewBasicBiTransform(DurationKiloYear, DurationYear, 1000)
	value.NewBasicBiTransform(DurationMegaYear, DurationKiloYear, 1000)
	value.NewBasicBiTransform(DurationGigaYear, DurationMegaYear, 1000)

	Duration = value.NewGroup("Duration", DurationSecond, DurationMinute, DurationHour, DurationDay, DurationYear)
}

const (
	durationMinute = 60
	durationHour   = 60 * durationMinute
	durationDay    = 24 * durationHour
	durationYear   = 365.2425 * durationDay
)

// DurationRoundDown will change the unit of v to a smaller unit as appropriate.
// e.g. If the value passed is DurationDay but the actual value is in minutes then this will
// return the value as DurationMinute.
func DurationRoundDown(v value.Value) value.Value {
	if v.IsValid() {
		if err := Duration.AssertValue(v); err != nil {
			return value.Value{}
		}

		if value.LessThan(v.Float(), 1.0) {
			switch v.Unit() {
			case DurationYear:
				return DurationRoundDown(v.AsOrInvalid(DurationDay))

			case DurationDay:
				return DurationRoundDown(v.AsOrInvalid(DurationHour))

			case DurationHour:
				return DurationRoundDown(v.AsOrInvalid(DurationMinute))

			case DurationMinute:
				return DurationRoundDown(v.AsOrInvalid(DurationSecond))
			}
		}
	}

	return v
}
