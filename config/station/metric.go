package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

// Metric represents a metric the containing type requires.
// This is used either as inbound, e.g. we need to retrieve the named metric,
// or outbound as in the target metric in a calculation
type Metric struct {
	Pos          lexer.Position
	Name         string      `parser:"@String"` // metric name
	Unit         *units.Unit `parser:"@@?"`     // optional Unit we require the metric to be in
	OriginalName string      // Set by init, original Name before any changes
}

func (c *visitor[T]) Metric(d *Metric) error {
	var err error
	if d != nil {
		if c.metric != nil {
			err = c.metric(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Unit(d.Unit)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initMetric(v Visitor[*initState], d *Metric) error {
	s := v.Get()

	var err error

	// enforce metrics to be lower case
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		err = errors.Errorf(d.Pos, "metric name is required")
	}

	if err == nil && strings.ContainsAny(d.Name, " /") {
		err = errors.Errorf(d.Pos, "metric name must not include '/' or spaces")
	}

	// Prefix with the stationId & sensorId to become a full metric id
	if err == nil && d.OriginalName == "" {
		d.OriginalName = d.Name
		d.Name = s.prefixMetric(d.Name)
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (m *Metric) AcceptMetric(v api.Metric) bool {
	return m != nil && v.Metric == m.Name
}

// Convert converts the passed Value to that of the Metric based on the requested Unit
func (m *Metric) Convert(v value.Value) (value.Value, error) {
	if m.Unit != nil {
		return v.As(m.Unit.Unit())
	}
	return v, nil
}
