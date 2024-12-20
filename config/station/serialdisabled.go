//go:build aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le))

package station

import (
	"github.com/alecthomas/participle/v2"
)

func initSerial(_ Visitor[*initState], d *Serial) error {
	return participle.Errorf(d.Pos, "serial devices are not supported on this platform")
}
