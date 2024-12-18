package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"time"
)

// MetricExpression represents a Metric within a Calculation
type MetricExpression struct {
	Pos    lexer.Position
	Metric *Metric `parser:"@@"`                    // Metric reference
	Offset string  `parser:"( 'offset' @String )?"` // optional offset in time, usually negative in the past
	// the parsed value of Offset
	offset time.Duration
}

func (c *visitor[T]) MetricExpression(d *MetricExpression) error {
	var err error
	if d != nil {
		if c.metricExpression != nil {
			err = c.metricExpression(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Metric(d.Metric)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initMetricExpression(_ Visitor[*initState], d *MetricExpression) error {
	var err error

	if d.Offset != "" {
		d.offset, err = time.ParseDuration(d.Offset)
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) MetricExpression(f func(Visitor[T], *MetricExpression) error) Builder[T] {
	b.metricExpression = f
	return b
}

// HasOffset returns true of Offset is defined and the parsed value is not 0
func (m *MetricExpression) HasOffset() bool {
	return m.offset != 0
}

// GetOffset returns the parsed offset, 0 if undefined
func (m *MetricExpression) GetOffset() time.Duration {
	return m.offset
}
