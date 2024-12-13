package time

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util/strings"
	"time"
)

type Duration struct {
	Pos      lexer.Position
	duration time.Duration // Parsed duration
	Def      string        `parser:"@String"` // Duration definition
}

func (a *Duration) IsEvery() bool {
	return a != nil && strings.In(a.Def, "every", "step")
}

func (a *Duration) Duration(every time.Duration) time.Duration {
	if a.IsEvery() {
		return every
	}
	return a.duration
}

func (a *Duration) Set(d time.Duration) {
	d = d.Truncate(time.Second)

	switch {
	case d > 0 && d < time.Second:
		d = time.Second

	case d < 0 && d > -time.Second:
		d = -time.Second
	}

	a.duration = d
}
