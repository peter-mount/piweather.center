package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type SourceParameter struct {
	Pos    lexer.Position
	Metric *Metric     `parser:"@@ '='"` // Metric to create with optional destination unit
	Source *SourcePath `parser:"@@"`     // Parameter path to get the value
	Unit   *units.Unit `parser:"@@"`     // Unit of Source - required
}

func (c *visitor[T]) SourceParameter(d *SourceParameter) error {
	var err error

	if d != nil {
		if c.sourceParameter != nil {
			err = c.sourceParameter(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Metric(d.Metric)
		}

		if err == nil {
			err = c.SourcePath(d.Source)
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		err = errors.Error(d.Pos, err)
	}

	return err
}

func initSourceParameter(v Visitor[*initState], d *SourceParameter) error {
	s := v.Get()

	err := v.Metric(d.Metric)

	if err == nil {
		err = v.SourcePath(d.Source)
	}

	if err == nil {
		err = v.Unit(d.Unit)
	}

	if err == nil {
		if e, exists := s.sensorParameters[d.Metric.Name]; exists {
			err = participle.Errorf(d.Metric.Pos, "metric %q already defined at %s", d.Metric.Name, e.Metric.Pos)
		} else {
			s.sensorParameters[d.Metric.Name] = d
		}
	}

	if err == nil {
		err = util.VisitorStop
	}

	return err
}

func (b *builder[T]) SourceParameter(f func(Visitor[T], *SourceParameter) error) Builder[T] {
	b.sourceParameter = f
	return b
}
