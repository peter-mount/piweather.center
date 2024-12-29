package weather

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	packages.RegisterPackage(&Weather{})
}

type Weather struct{}

func (_ Weather) DewPoint(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return measurement.GetDewPoint(temp, relHumidity)
}

func (_ Weather) FeelsLike(temp, relHumidity, windSpeed value.Value) (value.Value, error) {
	return measurement.FeelsLike(temp, relHumidity, windSpeed)
}

func (_ Weather) HeatIndex(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return measurement.HeatIndex(temp, relHumidity)
}

func (_ Weather) WindChill(temp, windSpeed value.Value) (value.Value, error) {
	return measurement.WindChill(temp, windSpeed)
}
