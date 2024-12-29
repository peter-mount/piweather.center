package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"strings"
)

// MetricPattern defines we require any metrics that match a simple pattern, either "*", "*abc", "abd*" or "*abc*".
// Note: "" or "**" is equivalent to "*".
type MetricPattern struct {
	Pos     lexer.Position
	Pattern string `parser:"@String"` // the pattern
	// the type of match
	patternType util.PatternType
	// prefix for server/sensor
	prefix string
}

func (c *visitor[T]) MetricPattern(d *MetricPattern) error {
	var err error
	if d != nil {
		if c.metricPattern != nil {
			err = c.metricPattern(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initMetricPattern(v Visitor[*initState], d *MetricPattern) error {
	s := v.Get()

	var err error

	t, p := util.ParsePatternType(d.Pattern)

	if strings.ContainsAny(p, " *") {
		err = errors.Errorf(d.Pos, "pattern must not include '*' or spaces")
	}

	// Disallow equality as that makes no sense for this component
	if err == nil && t == util.PatternEquals {
		err = errors.Errorf(d.Pos, "No wildcard provided")
	}

	// For MetricPattern we want "" as an alias for "*"
	if err == nil && t == util.PatternNone {
		t = util.PatternAll
	}

	if err == nil {
		d.Pattern = strings.ToLower(p)
		d.patternType = t
		d.prefix = s.prefixMetric("")
	}

	return err
}

func (b *builder[T]) MetricPattern(f func(Visitor[T], *MetricPattern) error) Builder[T] {
	b.metricPattern = f
	return b
}

func (m *MetricPattern) AcceptMetric(v api.Metric) bool {
	return m != nil && m.Match(v.Metric)
}

// Match returns true if the metric name passed matches the pattern
func (m *MetricPattern) Match(s string) bool {
	if strings.HasPrefix(s, m.prefix) {
		return m.patternType.Match(strings.TrimPrefix(s, m.prefix), m.Pattern)
	}
	return false
}
