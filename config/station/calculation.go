package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/time"
	"strings"
)

// Calculation defines a metric to calculate.
type Calculation struct {
	Pos        lexer.Position
	Target     string       `parser:"'(' @String"`           // Name of metric to calculate
	Every      time.CronTab `parser:"('every' @@)?"`         // Calculate at specified intervals
	ResetEvery time.CronTab `parser:"('reset' 'every' @@)?"` // Crontab to reset the value
	Load       *Load        `parser:"(@@)?"`                 // Load from the DB on startup
	UseFirst   *UseFirst    `parser:"(@@)?"`                 // If set and no value use this expression
	Expression *Expression  `parser:"('as' @@) ')'"`         // Expression to perform calculation
}

func (c *visitor[T]) Calculation(d *Calculation) error {
	var err error
	if d != nil {
		if c.calculation != nil {
			err = c.calculation(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = visitCalculation[T](c, d)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func visitCalculation[T any](v Visitor[T], d *Calculation) error {
	var err error
	if d != nil {
		err = v.CronTab(d.Every)

		if err == nil {
			err = v.CronTab(d.ResetEvery)
		}

		if err == nil {
			err = v.Load(d.Load)
		}

		if err == nil {
			err = v.UseFirst(d.UseFirst)
		}

		if err == nil {
			err = v.Expression(d.Expression)
		}
	}
	return err
}

func initCalculation(v Visitor[*initState], d *Calculation) error {
	s := v.Get()

	target := strings.ToLower(d.Target)

	if e, exists := s.calculations[target]; exists {
		return participle.Errorf(d.Pos, "calculation for %q already defined at %s", d.Target, e.String())
	}

	d.Target = s.prefixMetric(target)
	s.calculations[target] = d.Pos
	return nil
}

func (b *builder[T]) Calculation(f func(Visitor[T], *Calculation) error) Builder[T] {
	b.calculation = f
	return b
}

func printCalculation(v Visitor[*printState], d *Calculation) error {
	return v.Get().Run(d.Pos, func(st *printState) error {
		st.AppendHead("").
			AppendHead("calculate( %q", d.Target).
			AppendFooter(")")

		var err error

		if d.Every != nil {
			st.AppendBody("every %q", d.Every.Definition())
		}

		if d.ResetEvery != nil {
			st.AppendBody("reset every %q", d.ResetEvery.Definition())
		}

		if d.Load != nil {
			_ = v.Load(d.Load)
		}

		if d.UseFirst != nil && err == nil {
			_ = v.UseFirst(d.UseFirst)
		}

		return st.Run(d.Expression.Pos, func(st *printState) error {
			st.AppendHead("as")
			return v.Expression(d.Expression)
		})
	})
}
