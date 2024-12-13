//go:build aix || plan9 || solaris || windows

package weathersensor

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/station"
)

func (s *Service) i2cSensor(v station.Visitor[*state], sensor *station.I2C) error {
	return participle.Errorf(sensor.Pos, "i2c devices are not supported on this platform")
}
