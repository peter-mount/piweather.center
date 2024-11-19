//go:build aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le))

package weathersensor

import (
	"github.com/alecthomas/participle/v2"
	sensors2 "github.com/peter-mount/piweather.center/config/util/sensors"
)

func (s *Service) serialSensor(_ sensors2.SensorVisitor[any], sensor *sensors2.Sensor) error {
	return participle.Errorf(sensor.Pos, "serial devices are not supported on this platform")
}
