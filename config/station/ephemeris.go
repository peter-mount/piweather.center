package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"strings"
)

type Ephemeris struct {
	Pos       lexer.Position
	Target    string               `parser:"'ephemeris' '(' @String"` // Target sensorId
	Schedules []*EphemerisSchedule `parser:"@@+ ')'"`                 // schedules for targets
}

func (c *visitor[T]) Ephemeris(d *Ephemeris) error {
	var err error
	if d != nil {
		if c.ephemeris != nil {
			err = c.ephemeris(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Schedules {
				err = c.EphemerisSchedule(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initEphemeris(v Visitor[*initState], d *Ephemeris) error {
	s := v.Get()

	target := strings.ToLower(d.Target)

	if e, exists := s.calculations[target]; exists {
		return participle.Errorf(d.Pos, "calculation for %q already defined at %s", d.Target, e.String())
	}

	d.Target = s.prefixMetric(target)
	s.sensorPrefix = d.Target + "."
	s.ephemeris = d

	s.calculations[target] = d.Pos
	return nil
}

func (b *builder[T]) Ephemeris(f func(Visitor[T], *Ephemeris) error) Builder[T] {
	b.ephemeris = f
	return b
}
