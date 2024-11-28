package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/util"
	"strings"
)

type MetricList struct {
	Pos     lexer.Position
	Metrics []*Metric `parser:"@@*"`
}

type Metric struct {
	Pos  lexer.Position
	Name string      `parser:"@String"`
	Unit *units.Unit `parser:"@@?"`
	As   string      `parser:"('as' @String)?"`
}

type MetricPattern struct {
	Pos     lexer.Position
	Pattern string           `parser:"@String"` // the pattern, either "", "*abc", "abd*" or "*abc*"
	Type    util.PatternType // the type of match
	Prefix  string           // prefix for server/sensor
}

func (m *MetricPattern) Match(s string) bool {
	if strings.HasPrefix(s, m.Prefix) {
		return m.Type.Match(strings.TrimPrefix(s, m.Prefix), m.Pattern)
	}
	return false
}
