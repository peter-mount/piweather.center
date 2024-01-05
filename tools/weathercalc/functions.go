package weathercalc

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	value.NewCalculator("max", value.Basic2ArgCalculator(math.Max))
	value.NewCalculator("min", value.Basic2ArgCalculator(math.Min))
}
