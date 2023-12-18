package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
)

// QueryRange defines the time range to query
type QueryRange struct {
	Pos lexer.Position

	At    *Time     `parser:"( 'AT' @@"`       // AT time for a specific time
	From  *Time     `parser:"| 'FROM' @@"`     // FROM time
	For   *Duration `parser:"  'FOR' @@ "`     // Duration from FROM
	Start *Time     `parser:"| 'BETWEEN' @@"`  // Between a start time
	End   *Time     `parser:"  'AND' @@ )"`    // and an end time
	Every *Duration `parser:"( 'EVERY' @@ )?"` // Every duration time
}

func (a *QueryRange) Accept(v Visitor) error {
	return v.QueryRange(a)
}

func (a *QueryRange) Range() Range {
	var r Range
	if a != nil {
		switch {
		case a.At != nil:
			r = RangeAt(a.At.Time)

		case a.From != nil && a.For != nil:
			r = RangeFrom(a.From.Time, a.For.Duration)

		case a.Start != nil && a.End != nil:
			r = RangeBetween(a.Start.Time, a.End.Time)
		}

		if a.Every != nil {
			r.Every = a.Every.Duration
		}
	}
	return r
}
