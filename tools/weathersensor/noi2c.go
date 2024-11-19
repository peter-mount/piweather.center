//go:build aix || plan9 || solaris || windows

package weathersensor

import sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"

func (s *Service) i2cSensor(_ sensors2.SensorVisitor[any], _ *sensors2.Sensor) error {
	return participle.Errorf(sensor.Pos, "i2c devices are not supported on this platform")
}
