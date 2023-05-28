package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	RelativeHumidity = value.NewBoundedUnit("RelativeHumidity", "Relative Humidity", "%", 0, 0, 100)
	Humidity = value.NewGroup("Humidity", RelativeHumidity)
}

var (
	Humidity         *value.Group
	RelativeHumidity *value.Unit
)
