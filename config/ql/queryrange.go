package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/api"
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

func (a *QueryRange) Accept(v Visitor) error {
	return v.QueryRange(a)
}

func (a *QueryRange) Range() api.Range {
	var r api.Range
	if a != nil {
		switch {
		case a.At != nil:
			r = api.RangeAt(a.At.Time)

		case a.From != nil && a.For != nil:
			r = api.RangeFrom(a.From.Time, a.For.Duration(0))

		case a.Start != nil && a.End != nil:
			r = api.RangeBetween(a.Start.Time, a.End.Time)
		}

		if a.Every != nil {
			r.Every = a.Every.Duration(0)
		}
	}
	return r
}

func (a *QueryRange) IsRow() bool {
	return a.At.IsRow() || a.From.IsRow() || a.Start.IsRow() || a.End.IsRow()
}

func (a *QueryRange) SetTime(t time.Time, every time.Duration, v Visitor) error {
	err := a.At.SetTime(t, every, v)
	if err == nil {
		err = a.From.SetTime(t, every, v)
	}
	if a.For.IsEvery() {
		a.For.Set(every)
	}
	if err == nil {
		err = a.Start.SetTime(t, every, v)
	}
	if err == nil {
		err = a.End.SetTime(t, every, v)
	}
	return err
}
