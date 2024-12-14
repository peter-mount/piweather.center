package time

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"gopkg.in/robfig/cron.v2"
	"strings"
	"time"
)

type CronTab interface {
	Position() lexer.Position
	Definition() string
	Schedule() cron.Schedule
	SetLocation(*time.Location) error
}

type cronTab struct {
	pos        lexer.Position
	definition string
	schedule   cron.Schedule
}

func (c *cronTab) Definition() string {
	return c.definition
}

func (c *cronTab) Position() lexer.Position {
	return c.pos
}

func (c *cronTab) String() string {
	return c.definition
}

func (c *cronTab) Schedule() cron.Schedule {
	return c.schedule
}

func CronTabParser(l lexer.Definition) func(pl *lexer.PeekingLexer) (CronTab, error) {
	// We need to use just "@String" for the definition
	stringToken := l.Symbols()["String"]

	return func(pl *lexer.PeekingLexer) (CronTab, error) {
		token := pl.Peek()
		if token.EOF() {
			return nil, participle.Errorf(token.Pos, "EOF expected crontab definition")
		}

		if token.Type != stringToken {
			return nil, participle.Errorf(token.Pos, "Expected crontab definition string")
		}

		// Get the definition, moving the cursor forward
		definition := pl.Next().String()
		definition = strings.TrimSpace(definition)

		// Disallow setting timezone here - that should be done higher up
		if strings.HasPrefix(definition, "TZ=") {
			return nil, participle.Errorf(token.Pos, "crontab definition cannot define timezone")
		}

		// Convert aliases to actual definitions
		definition = strings.ToLower(definition)
		switch definition {
		case "":
			return nil, participle.Errorf(token.Pos, "crontab definition cannot be empty")

		case "day", "daily", "midnight":
			definition = "0 0 0 * * *"

		case "hour", "hourly":
			definition = "0 0 * * * *"

		case "minute":
			definition = "0 * * * * *"

		case "second":
			definition = "* * * * * *"
		}

		schedule, err := cron.Parse(definition)
		if err != nil {
			return nil, participle.Errorf(token.Pos, "invalid crontab definition %q: %v", definition, err)
		}

		return &cronTab{pos: token.Pos, definition: definition, schedule: schedule}, nil
	}
}

func (c *cronTab) SetLocation(loc *time.Location) error {
	schedule, err := cron.Parse("TZ=" + loc.String() + " " + c.definition)
	if err != nil {
		return participle.Errorf(c.pos, "invalid crontab definition %q: %v", c.definition, err)
	}
	c.schedule = schedule
	return nil
}
