//go:build !(aix || dragonfly || illumos || netbsd || plan9 || solaris || (linux && (loong64 || ppc64 || ppc64le)))

package station

import (
	"github.com/alecthomas/participle/v2"
	"strings"
)

func initSerial(_ Visitor[*initState], d *Serial) error {
	d.Driver = strings.TrimSpace(d.Driver)
	if d.Driver == "" {
		return participle.Errorf(d.Pos, "no Driver defined")
	}

	d.Port = strings.TrimSpace(d.Port)
	if d.Port == "" {
		return participle.Errorf(d.Pos, "no serial port defined")
	}

	// TODO define a common list of baud rates here?
	if d.BaudRate < 300 {
		return participle.Errorf(d.Pos, "Invalid baud rate")
	}

	return nil
}
