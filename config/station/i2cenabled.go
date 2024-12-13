//go:build !(aix || plan9 || solaris || windows)

package station

import (
	"github.com/alecthomas/participle/v2"
	"strings"
)

func initI2c(_ Visitor[*initState], d *I2C) error {
	d.Driver = strings.TrimSpace(d.Driver)
	if d.Driver == "" {
		return participle.Errorf(d.Pos, "no Driver defined")
	}

	if d.Bus < 1 || d.Device < 1 {
		return participle.Errorf(d.Pos, "invalid i2c address, got (%d:%d)", d.Bus, d.Device)
	}
	return nil
}
