package ql

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util"
	"strings"
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

func (a *Time) IsRow() bool {
	return a != nil && strings.ToLower(a.Def) == "row"
}

func (a *Time) SetTime(t time.Time, every time.Duration, v Visitor) error {
	if a == nil {
		return nil
	}

	a.Time = t

	if a.Time.IsZero() && !a.IsRow() {
		return participle.Errorf(a.Pos, "invalid datetime")
	}

	for _, e := range a.Expression {
		switch {
		case e.Add != nil:
			if err := v.Duration(e.Add); err != nil {
				return err
			}
			a.Time = a.Time.Add(e.Add.Duration(every))

		case e.Truncate != nil:
			if err := v.Duration(e.Truncate); err != nil {
				return err
			}
			a.Time = a.Time.Truncate(e.Truncate.Duration(every))
		}
	}

	return nil
}

type TimeExpression struct {
	Pos      lexer.Position
	Add      *Duration `parser:"( 'ADD' @@"`        // Add duration to time
	Truncate *Duration `parser:"| 'TRUNCATE' @@ )"` // truncate time
}

type Duration struct {
	Pos      lexer.Position
	duration time.Duration // Parsed duration
	Def      string        `parser:"@String"` // Duration definition
}

func (a *Duration) IsEvery() bool {
	return a != nil && util.In(a.Def, "every", "step")
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
