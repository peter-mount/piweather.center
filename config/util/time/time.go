package time

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
	"time"
)

type Time struct {
	Pos        lexer.Position
	Time       time.Time         // The parsed time
	Def        string            `parser:"@String"` // Time definition
	Expression []*TimeExpression `parser:"(@@)*"`
}

func (a *Time) IsRow() bool {
	return a != nil && strings.ToLower(a.Def) == "row"
}

func (a *Time) SetTime(t time.Time, every time.Duration, v TimeVisitor) error {
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
