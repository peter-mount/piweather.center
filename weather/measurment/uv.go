package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	UV = value.NewLowerBoundUnit("UV", "", value.Dp0, 0)
}

var (
	// UV Unit representing UV (Ultra Violet) index
	UV value.Unit
)

// IsUV returns true if the value is a UV index
func IsUV(v value.Value) bool {
	return v.Unit() == UV
}
