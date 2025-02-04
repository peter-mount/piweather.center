package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"strings"
)

// CalculateFrom is a helper to reduce boilerplate calculations.
type CalculateFrom struct {
	Pos         lexer.Position
	From        string          `parser:"@String"`               // Src metric
	Unit        *units.Unit     `parser:"@@?"`                   // optional Unit we require the results to use
	Aggregators *AggregatorList `parser:"@@"`                    // List of aggregators to calculate
	ResetEvery  time.CronTab    `parser:"('reset' 'every' @@)?"` // Crontab to reset the value
}

func (c *visitor[T]) CalculateFrom(d *CalculateFrom) error {
	var err error
	if d != nil {
		if c.calculateFrom != nil {
			err = c.calculateFrom(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		if err == nil {
			err = c.CronTab(d.ResetEvery)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initCalculateFrom(v Visitor[*initState], d *CalculateFrom) error {
	s := v.Get()

	from := strings.ToLower(d.From)

	for _, ag := range d.Aggregators.GetAggregators() {
		target := from + "." + ag

		if e, exists := s.calculations[target]; exists {
			return participle.Errorf(d.Pos, "calculation for %q already defined at %s", target, e.String())
		}

		//d.Target = s.prefixMetric(target)
		s.calculations[target] = d.Pos
	}

	return nil
}

func (b *builder[T]) CalculateFrom(f func(Visitor[T], *CalculateFrom) error) Builder[T] {
	b.calculateFrom = f
	return b
}
