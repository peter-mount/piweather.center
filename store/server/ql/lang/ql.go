package lang

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/util/unit"
	"strings"
	"time"
)

type Query struct {
	Pos lexer.Position

	QueryRange *QueryRange `parser:"@@"`
	Select     []*Select   `parser:"( @@ )+"`
}

func (a *Query) Accept(v Visitor) error {
	return v.Query(a)
}

type Select struct {
	Pos lexer.Position

	Expression *SelectExpression `parser:"'SELECT' @@"`
	//Limit      *Expression       `( "LIMIT" @@ )?`
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

	At       *Time         `parser:"( 'AT' @@"`       // AT time for a specific time
	From     *Time         `parser:"| 'FROM' @@"`     // FROM time
	For      *Duration     `parser:"  'FOR' @@ "`     // Duration from FROM
	Start    *Time         `parser:"| 'BETWEEN' @@"`  // Between a start time
	End      *Time         `parser:"  'AND' @@ )"`    // and an end time
	Every    *Duration     `parser:"( 'EVERY' @@ )?"` // Every duration time
	StepSize time.Duration // The required step size
}

func queryRangeInit(_ Visitor, q *QueryRange) error {
	if q.Every != nil {
		q.StepSize = q.Every.Duration
	}

	// Ensure we have a valid positive step size.
	// If not defined then default to 1 minute,
	// otherwise set minimum of 1 second as that is our db resolution
	switch {
	case q.StepSize < 0:
		return fmt.Errorf("invalid step size %v", q.StepSize)

	case q.StepSize == 0:
		q.StepSize = time.Minute

	case q.StepSize < time.Second:
		q.StepSize = time.Second
	}

	return nil
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
	Pos  lexer.Position
	Def  string    `parser:"@String"` // Time definition
	Time time.Time // The parsed time
}

func (a *Time) Accept(v Visitor) error {
	return v.Time(a)
}

func timeInit(_ Visitor, t *Time) error {
	if t == nil {
		return nil
	}

	t.Time = unit.ParseTime(t.Def)
	if t.Time.IsZero() {
		return participle.Errorf(t.Pos, "invalid datetime")
	}
	return nil
}

type Duration struct {
	Pos      lexer.Position
	Def      string        `parser:"@String"` // Duration definition
	Duration time.Duration // Parsed duration
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
