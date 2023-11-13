package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	CountPerMinute = value.NewLowerBoundUnit("CountPerMinute", "Count Per Minute", " cpm", 0, 0)
	Sievert = value.NewLowerBoundUnit("Sievert", "Sievert", " Sv", 3, 0)
	MicroSievert = value.NewLowerBoundUnit("MicroSievert", "MicroSievert", " µSv", 3, 0)
	RoentgenEquivalentMan = value.NewLowerBoundUnit("RoentgenEquivalentMan", "Roentgen equivalent man", " rem", 0, 0)

	value.NewBasicBiTransform(CountPerMinute, MicroSievert, cpm2uSv)
	value.NewBasicBiTransform(Sievert, MicroSievert, Micro)
	value.NewBasicBiTransform(CountPerMinute, Sievert, cpm2uSv*Micro)

	value.NewBasicBiTransform(Sievert, RoentgenEquivalentMan, 0.01)
	value.NewBasicBiTransform(CountPerMinute, RoentgenEquivalentMan, cpm2uSv*Micro*0.01)

	// This is so the group is valid.
	// We record CPS as a raw value in geiger realtime mode, so we need a Unit present, but it's
	// not valid to actually transform it, hence a transform factor of 0 will issue a fail
	CountPerSecond = value.NewLowerBoundUnit("CountPerSecond", "Count Per Second", " cps", 0, 0)
	value.NewBasicBiTransform(CountPerSecond, CountPerMinute, 0)

	Radiation = value.NewGroup("Radiation", CountPerMinute, Sievert, MicroSievert, RoentgenEquivalentMan, CountPerSecond)
}

const (
	// NOTE: This is based on my Geiger counter, other devices will vary!
	// For me 15cpm equates to 0.10µSv
	cpm2uSv = 0.10 / 15.0
)

var (
	Radiation             *value.Group
	CountPerMinute        *value.Unit
	CountPerSecond        *value.Unit
	Sievert               *value.Unit
	MicroSievert          *value.Unit
	RoentgenEquivalentMan *value.Unit
)
