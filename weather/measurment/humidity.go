package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	Humidity = value.NewBoundedUnit("Humidity", "%", value.Dp0, 0, 100)
	RelativeHumidity = value.NewBoundedUnit("Relative Humidity", "%", value.Dp0, 0, 100)
}

var (
	Humidity         value.Unit
	RelativeHumidity value.Unit
)
