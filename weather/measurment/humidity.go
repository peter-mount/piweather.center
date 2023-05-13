package measurment

import (
	"errors"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	Humidity = value.NewBoundedUnit("Humidity", "%", value.Dp0, 0, 100)
	RelativeHumidity = value.NewBoundedUnit("Relative Humidity", "%", value.Dp0, 0, 100)
}

var (
	Humidity                 value.Unit
	RelativeHumidity         value.Unit
	notHumidityError         = errors.New("value not Humidity")
	notRelativeHumidityError = errors.New("value not RelativeHumidity")
)

func IsHumidity(v value.Value) bool { return v.Unit() == Humidity }

func AssertHumidity(v value.Value) error {
	if IsHumidity(v) {
		return nil
	}
	return notHumidityError
}

func IsHumidityErr(e error) bool {
	return e == notHumidityError
}

func IsRelativeHumidity(v value.Value) bool { return v.Unit() == RelativeHumidity }

func AssertRelativeHumidity(v value.Value) error {
	// If v is RelativeHumidity then call BoundsError which will be nil unless it's invalid
	if IsRelativeHumidity(v) {
		return v.BoundsError()
	}
	return notHumidityError
}

func IsRelativeHumidityErr(e error) bool {
	return e == notRelativeHumidityError
}
