package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/store/api"
	"time"
)

// QueryRange defines the time range to query
type QueryRange struct {
	Pos lexer.Position

	At    *time2.Time     `parser:"( 'at' @@"`       // AT time for a specific time
	From  *time2.Time     `parser:"| 'from' @@"`     // FROM time
	For   *time2.Duration `parser:"  'for' @@ "`     // Duration from FROM
	Start *time2.Time     `parser:"| 'between' @@"`  // Between a start time
	End   *time2.Time     `parser:"  'and' @@ )"`    // and an end time
	Every *time2.Duration `parser:"( 'every' @@ )?"` // Every duration time
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

func (a *QueryRange) SetTime(t time.Time, every time.Duration, v QueryVisitor) error {
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
