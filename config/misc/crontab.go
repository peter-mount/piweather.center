package misc

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

type CronTab struct {
	Pos        lexer.Position
	Definition string `parser:"@String"` // CronTab definition
}

func (l *CronTab) Init() error {
	// Convert aliases to actual definitions
	switch strings.ToLower(l.Definition) {
	case "day", "daily", "midnight":
		l.Definition = "0 0 0 * * *"
	case "hour", "hourly":
		l.Definition = "0 0 * * * *"
	case "minute":
		l.Definition = "0 * * * * *"
	case "second":
		l.Definition = "* * * * * *"
	}

	return nil
}
