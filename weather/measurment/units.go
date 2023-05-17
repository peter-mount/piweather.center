package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	// Miscellaneous units which are not part of a family
	UV = value.NewLowerBoundUnit("UV", "Indices", "UV", " UV index", value.Dp0, 0)
	Strike = value.NewLowerBoundUnit("Strikes", "Lightning", "Lightning Strikes", " strikes", value.Dp0, 0)
}

var (
	// UV Unit representing UV (Ultra Violet) index
	UV     *value.Unit
	Strike *value.Unit
)
