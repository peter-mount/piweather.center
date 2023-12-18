package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
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

type Duration struct {
	Pos      lexer.Position
	Duration time.Duration // Parsed duration
	Def      string        `parser:"@String"` // Duration definition
}

func (a *Duration) Set(d time.Duration) {
	d = d.Truncate(time.Second)

	switch {
	case d > 0 && d < time.Second:
		d = time.Second

	case d < 0 && d > -time.Second:
		d = -time.Second
	}

	a.Duration = d
	a.Def = d.String()
}
