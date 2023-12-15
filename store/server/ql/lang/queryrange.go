package lang

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"time"
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

func queryRangeInit(v Visitor, q *QueryRange) error {
	// If no Every statement then set it to 1 minute
	if q.Every == nil {
		q.Every = &Duration{
			Pos:      q.Pos,
			Duration: time.Minute,
			Def:      "1m",
		}
	}

	if err := v.Duration(q.Every); err != nil {
		return err
	}

	// Negative duration for Every is invalid
	if q.Every.Duration < 0 {
		return fmt.Errorf("invalid step size %v", q.Every.Duration)
	}

	// The DB has a second resolution so any Every less than that defaults to 1 s
	q.Every.Set(q.Every.Duration.Truncate(time.Second))

	if q.At != nil {
		if err := v.Time(q.At); err != nil {
			return err
		}
	}

	if q.From != nil {
		if err := v.Time(q.From); err != nil {
			return err
		}
	}

	if q.For != nil {
		if err := v.Duration(q.For); err != nil {
			return err
		}
	}

	if q.Start != nil {
		if err := v.Time(q.Start); err != nil {
			return err
		}
	}

	if q.End != nil {
		if err := v.Time(q.End); err != nil {
			return err
		}
	}

	return VisitorStop
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
