package lang

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util/unit"
	"strings"
	"time"
)

type Select struct {
	Pos lexer.Position

	Expression *SelectExpression `parser:"'SELECT' @@"`
	Limit      int               `parser:"( 'LIMIT' @Int )?"`
}

func selectInit(_ Visitor, s *Select) error {
	return assertLimit(s.Pos, s.Limit)
}

func (a *Select) Accept(v Visitor) error {
	return v.Select(a)
}

type SelectExpression struct {
	Pos lexer.Position

	All         bool                 `parser:"( @'*'"`
	Expressions []*AliasedExpression `parser:"| @@ ( ',' @@ )* )"`
}

func (a *SelectExpression) Accept(v Visitor) error {
	return v.SelectExpression(a)
}

// AliasedExpression handles expression AS name to create aliases
type AliasedExpression struct {
	Pos lexer.Position

	Expression *Expression `parser:"@@"`
	As         string      `parser:"( 'AS' @Ident )?"`
}

func (a *AliasedExpression) Accept(v Visitor) error {
	return v.AliasedExpression(a)
}

// Expression handles function calls or direct metric values
type Expression struct {
	Pos lexer.Position

	Function *Function `parser:"( @@"`
	Metric   *Metric   `parser:"| @@ )"`
	Offset   *Duration `parser:"( 'OFFSET' @@ )?"`
}

func (a *Expression) Accept(v Visitor) error {
	return v.Expression(a)
}

// Function handles function calls
type Function struct {
	Pos lexer.Position

	TimeOf      bool          `parser:"( @'TIMEOF'"`
	Name        string        `parser:"| @Ident"`
	Expressions []*Expression `parser:") '(' (@@ (',' @@)*)? ')'"`
}

func (a *Function) Accept(v Visitor) error {
	return v.Function(a)
}

// Metric handles a metric reference
type Metric struct {
	Pos lexer.Position

	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

func (a *Metric) Accept(v Visitor) error {
	return v.Metric(a)
}

func metricInit(_ Visitor, b *Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

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

	// Negative duration for Every is invalid
	if q.Every.Duration < 0 {
		return fmt.Errorf("invalid step size %v", q.Every.Duration)
	}

	// The DB has a second resolution so any Every less than that defaults to 1 s
	q.Every.Set(q.Every.Duration.Truncate(time.Second))

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
