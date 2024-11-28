package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Component struct {
	Pos lexer.Position
	//Type      string     `yaml:"type"`            // type of component - required
	Title     string     `parser:"('title' @String)?"` // title, optional based on component
	Class     string     `parser:"('class' @String)?"` // optional CSS class
	Style     string     `parser:"('style' @String)?"` // optional inline CSS
	ID        string     // Unique ID, generated on load
	dashboard *Dashboard // link to dashboard
}

type Value struct {
	Pos       lexer.Position
	Type      string      `parser:"@('value') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

type MultiValue struct {
	Pos       lexer.Position
	Component *Component     `parser:"'multiValue' '(' @@"`
	Pattern   *MetricPattern `parser:"@@"`
	Time      bool           `parser:"@'time'? ')'"`
}

type Gauge struct {
	Pos       lexer.Position
	Type      string      `parser:"@('gauge' | 'barometer') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Unit      *units.Unit `parser:"@@?"`
	Metrics   *MetricList `parser:"@@"`
	Min       float64     `parser:"('min' @Number)?"`
	Max       float64     `parser:"('max' @Number)?"`
	Ticks     int         `parser:"('ticks' @Number)? ')'"`
}

func (g *Gauge) Convert(v value.Value) (value.Value, error) {
	var err error

	// Convert v to either the specified unit or that of the first metric
	if g.Unit == nil {
		// This is safe as the parser ensures that metrics contain at least 1 metric for gauges
		v, err = g.Metrics.Metrics[0].Convert(v)
	} else {
		v, err = g.Unit.Convert(v)
	}

	return v, errors.Error(g.Pos, err)
}
