package ql

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
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

func (v *visitor[T]) QueryRange(b *QueryRange) error {
	var err error
	if b != nil {
		if v.queryRange != nil {
			err = v.queryRange(v, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		// AT x
		if err == nil {
			err = v.Time(b.At)
		}

		// FROM x FOR x
		if err == nil {
			err = v.Time(b.From)
		}
		if err == nil {
			err = v.Duration(b.For)
		}

		// BETWEEN x AND x
		if err == nil {
			err = v.Time(b.Start)
		}
		if err == nil {
			err = v.Time(b.End)
		}

		if err == nil {
			err = v.Duration(b.Every)
		}
	}
	return err
}

func initQueryRange(v Visitor[*parserState], q *QueryRange) error {
	// If no Every statement then set it to 1 minute
	if q.Every == nil {
		q.Every = &time2.Duration{Pos: q.Pos, Def: "1m"}
	}

	if err := v.Duration(q.Every); err != nil {
		return err
	}

	// Negative duration for Every is invalid
	if q.Every.Duration(0) < time.Second {
		return fmt.Errorf("invalid step size %v", q.Every.Duration(0))
	}

	if err := v.Time(q.At); err != nil {
		return err
	}

	if err := v.Time(q.From); err != nil {
		return err
	}
	if err := v.Duration(q.For); err != nil {
		return err
	}

	if err := v.Time(q.Start); err != nil {
		return err
	}
	if err := v.Time(q.End); err != nil {
		return err
	}

	return util.VisitorStop
}

func (b *builder[T]) QueryRange(f func(Visitor[T], *QueryRange) error) Builder[T] {
	b.common.queryRange = f
	return b
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

//func (a *QueryRange) SetTime(t time.Time, every time.Duration, v Visitor) error {
//	err := a.At.SetTime(t, every, v)
//	if err == nil {
//		err = a.From.SetTime(t, every, v)
//	}
//	if a.For.IsEvery() {
//		a.For.Set(every)
//	}
//	if err == nil {
//		err = a.Start.SetTime(t, every, v)
//	}
//	if err == nil {
//		err = a.End.SetTime(t, every, v)
//	}
//	return err
//}
