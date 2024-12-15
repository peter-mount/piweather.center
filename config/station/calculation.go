package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
	"strings"
)

// Calculation defines a metric to calculate
type Calculation struct {
	Pos        lexer.Position
	Target     string       `parser:"'calculate' '(' @String"` // Name of metric to calculate
	Every      time.CronTab `parser:"('every' @@)?"`           // Calculate at specified intervals
	ResetEvery time.CronTab `parser:"('reset' 'every' @@)?"`   // Crontab to reset the value
	Load       *Load        `parser:"(@@)?"`                   // Load from the DB on startup
	UseFirst   *UseFirst    `parser:"(@@)?"`                   // If set and no value use this expression
	Expression *Expression  `parser:"('as' @@) ')'"`           // Expression to perform calculation
}

func (c *visitor[T]) Calculation(d *Calculation) error {
	var err error
	if d != nil {
		if c.calculation != nil {
			err = c.calculation(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.CronTab(d.Every)
		}

		if err == nil {
			err = c.CronTab(d.ResetEvery)
		}

		if err == nil {
			err = c.Load(d.Load)
		}

		if err == nil {
			err = c.UseFirst(d.UseFirst)
		}

		if err == nil {
			err = c.Expression(d.Expression)
		}

		err = errors.Error(d.Pos, err)
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
