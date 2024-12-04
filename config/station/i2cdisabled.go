//go:build aix || plan9 || solaris || windows

package station

import (
	"github.com/alecthomas/participle/v2"
)

func initI2c(_ Visitor[*initState], d *I2C) error {
	return participle.Errorf(d.Pos, "i2c devices are not supported on this platform")
}
