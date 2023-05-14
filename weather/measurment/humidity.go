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
	// Humidity generic humidity unit
	Humidity value.Unit
	// RelativeHumidity unit
	RelativeHumidity         value.Unit
	notHumidityError         = errors.New("value not Humidity")
	notRelativeHumidityError = errors.New("value not RelativeHumidity")
)

// IsHumidity returns true if the Value is Humidity
func IsHumidity(v value.Value) bool { return v.Unit() == Humidity }

// AssertHumidity returns an error if the value is not Humidity
func AssertHumidity(v value.Value) error {
	if IsHumidity(v) {
		return nil
	}
	return notHumidityError
}

// IsHumidityErr returns true if the error was returned by AssertHumidity
func IsHumidityErr(e error) bool {
	return e == notHumidityError
}

// IsRelativeHumidity returns true if the Value is RelativeHumidity
func IsRelativeHumidity(v value.Value) bool { return v.Unit() == RelativeHumidity }

// AssertRelativeHumidity returns an error if the value is not RelativeHumidity
func AssertRelativeHumidity(v value.Value) error {
	// If v is RelativeHumidity then call BoundsError which will be nil unless it's invalid
	if IsRelativeHumidity(v) {
		return v.BoundsError()
	}
	return notHumidityError
}

// IsRelativeHumidityErr returns true if the error was returned by AssertRelativeHumidity
func IsRelativeHumidityErr(e error) bool {
	return e == notRelativeHumidityError
}
