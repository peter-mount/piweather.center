package lang

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util/unit"
	"time"
)

type Time struct {
	Pos        lexer.Position
	Time       time.Time         // The parsed time
	Def        string            `parser:"@String"` // Time definition
	Expression []*TimeExpression `parser:"(@@)*"`
}

func (a *Time) Accept(v Visitor) error {
	return v.Time(a)
}

type TimeExpression struct {
	Pos      lexer.Position
	Add      *Duration `parser:"( 'ADD' @@"`        // Add duration to time
	Truncate *Duration `parser:"| 'TRUNCATE' @@ )"` // truncate time
}

func timeInit(v Visitor, t *Time) error {
	if t == nil {
		return nil
	}

	t.Time = unit.ParseTime(t.Def)
	if t.Time.IsZero() {
		return participle.Errorf(t.Pos, "invalid datetime")
	}

	for _, e := range t.Expression {
		switch {
		case e.Add != nil:
			if err := v.Duration(e.Add); err != nil {
				return err
			}
			t.Time = t.Time.Add(e.Add.Duration)

		case e.Truncate != nil:
			if err := v.Duration(e.Truncate); err != nil {
				return err
			}
			t.Time = t.Time.Truncate(e.Truncate.Duration)
		}
	}

	return VisitorStop
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
