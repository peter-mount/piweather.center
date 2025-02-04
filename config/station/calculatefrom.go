package station

import (
	"fmt"
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
	initialised bool            // true once this instance has been initialised
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
			err = visitCalculateFrom[T](c, d)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func visitCalculateFrom[T any](v Visitor[T], d *CalculateFrom) error {
	var err error
	if d != nil {
		err = v.Unit(d.Unit)

		if err == nil {
			err = v.CronTab(d.ResetEvery)
		}
	}
	return err
}

func initCalculateFrom(v Visitor[*initState], d *CalculateFrom) error {
	// Only initialise once
	if d.initialised {
		return nil
	}

	s := v.Get()

	from := strings.ToLower(d.From)

	var script []string
	for _, ag := range d.Aggregators.GetAggregators() {
		target := strings.ToLower(from + "." + ag)
		if err := s.assertCalculation(d.Pos, target); err != nil {
			return err
		}

		script = append(script, fmt.Sprintf("calculate( %q", target))

		if d.ResetEvery != nil {
			script = append(script, fmt.Sprintf("reset every %q", d.ResetEvery.Definition()))
		}

		script = append(script,
			fmt.Sprintf(`load %q with "%s(%s)"`, "today", ag, from),
			fmt.Sprintf("usefirst %q", from),
			fmt.Sprintf("as %s(current,%q)", ag, from),
			")")
		s.calculations[target] = d.Pos
	}

	// Now parse this new script
	err := s.ParseAndInclude(d.Pos, script)
	if err == nil {
		// Mark so we don't initialise again
		d.initialised = true
	}

	if err == nil {
		err = errors.VisitorStop
	}
	return errors.Error(d.Pos, err)
}

func (b *builder[T]) CalculateFrom(f func(Visitor[T], *CalculateFrom) error) Builder[T] {
	b.calculateFrom = f
	return b
}

func printCalculateFrom(v Visitor[*printState], d *CalculateFrom) error {
	return v.Get().Run(d.Pos, func(st *printState) error {
		// calculate from is expanded so we show it as a comment
		st.Comment().
			AppendPos(d.Pos).
			AppendHead("// calculate from %q", d.From)

		if d.Unit != nil {
			st.Start().
				AppendHead("unit %q", d.Unit.Using).
				End()
		}

		if d.Aggregators != nil {
			st.Start().
				AppendHead("( %s )", strings.Join(d.Aggregators.GetAggregators(), " ")).
				End()
		}

		if d.ResetEvery != nil {
			st.Start().
				AppendHead("reset every %q", d.ResetEvery.Definition()).
				End()
		}

		return nil
	})
}
