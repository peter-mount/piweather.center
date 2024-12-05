package time

import (
	"github.com/alecthomas/participle/v2/lexer"
	"time"
)

type TimeZone struct {
	Pos      lexer.Position
	TimeZone string `parser:"'timezone' @String"`
	location *time.Location
}

func (t *TimeZone) Init() error {
	if t.TimeZone != "" {
		t.location = time.UTC
	}
	loc, err := time.LoadLocation(t.TimeZone)
	if err != nil {
		return err
	}
	t.location = loc
	return nil
}

func (t *TimeZone) Location() *time.Location {
	if t.location == nil {
		t.location = time.Local
	}
	return t.location
}
