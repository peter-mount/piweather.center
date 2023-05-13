package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	UV = value.NewLowerBoundUnit("UV", "", value.Dp0, 0)
}

var (
	UV value.Unit
)
