//go:build aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le))

package weathersensor

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/station"
)

func (s *Service) serialSensor(v station.Visitor[*state], sensor *station.Serial) error {
	return participle.Errorf(sensor.Pos, "serial devices are not supported on this platform")
}
