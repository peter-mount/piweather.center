//go:build aix || plan9 || solaris || windows

package weathersensor

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/util/sensors"
)

func (s *Service) i2cSensor(_ sensors.SensorVisitor[any], sensor *sensors.Sensor) error {
	return participle.Errorf(sensor.Pos, "i2c devices are not supported on this platform")
}
