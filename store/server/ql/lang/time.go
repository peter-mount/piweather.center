package lang

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util/unit"
	"time"
)

type Time struct {
	Pos      lexer.Position
	Time     time.Time // The parsed time
	Def      string    `parser:"@String"`            // Time definition
	Add      *Duration `parser:"( 'ADD' @@ )?"`      // Add duration to time
	Truncate *Duration `parser:"( 'TRUNCATE' @@ )?"` // truncate time
}

func (a *Time) Accept(v Visitor) error {
	return v.Time(a)
}

func timeInit(v Visitor, t *Time) error {
	if t == nil {
		return nil
	}

	t.Time = unit.ParseTime(t.Def)
	if t.Time.IsZero() {
		return participle.Errorf(t.Pos, "invalid datetime")
	}

	if t.Add != nil {
		if err := v.Duration(t.Add); err != nil {
			return err
		}
		t.Time = t.Time.Add(t.Add.Duration)
	}

	if t.Truncate != nil {
		if err := v.Duration(t.Truncate); err != nil {
			return err
		}
		t.Time = t.Time.Truncate(t.Truncate.Duration)
	}

	return nil
}

type Duration struct {
	Pos      lexer.Position
	Duration time.Duration // Parsed duration
	Def      string        `parser:"@String"` // Duration definition
}

func durationInit(_ Visitor, d *Duration) error {
	if d.Def != "" {
		v, err := time.ParseDuration(d.Def)
		if err != nil {
			return err
		}
		d.Duration = v

		// Ensure we keep within sensible ranges
		switch {
		case d.Duration > 0 && d.Duration < time.Second:
			d.Duration = time.Second

		case d.Duration < 0 && d.Duration > time.Second:
			d.Duration = -time.Second
		}
	}

	return nil
}

func (a *Duration) Set(d time.Duration) {
	a.Duration = d
	a.Def = d.String()
}
