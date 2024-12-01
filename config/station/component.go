package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
)

type ComponentType interface {
	GetType() string
}

type ComponentId interface {
	GetID() string
}

type ComponentDefinition interface {
	Definition() any
}

type ComponentProcessor interface {
	Process(ComponentStore, api.Metric)
}

type ComponentStore interface {
	// Add to a store.
	//
	// t component type
	//
	// id component id
	//
	// i & s metric index & suffix within the Component
	//
	// m api.Metric to add
	Add(t, id string, i int, s string, m api.Metric)
}

type Component struct {
	Pos lexer.Position
	//Type      string     `yaml:"type"`            // type of component - required
	Title     string     `parser:"('title' @String)?"` // title, optional based on component
	Class     string     `parser:"('class' @String)?"` // optional CSS class
	Style     string     `parser:"('style' @String)?"` // optional inline CSS
	ID        string     // Unique ID, generated on load
	dashboard *Dashboard // link to dashboard
}

func (c *Component) GetID() string {
	return c.ID
}

func (c *Component) GetType() string {
	return "component"
}

type Value struct {
	Pos       lexer.Position
	Type      string      `parser:"@('value') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

func (c *Value) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Metrics.AcceptMetric(v)
}

func (c *Value) GetID() string {
	return c.Component.GetID()
}

func (c *Value) GetType() string {
	return c.Type
}

type MultiValue struct {
	Pos       lexer.Position
	Component *Component     `parser:"'multivalue' '(' @@"`
	Pattern   *MetricPattern `parser:"@@"`
	Time      bool           `parser:"@'time'? ')'"`
}

func (c *MultiValue) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Pattern.AcceptMetric(v)
}

func (c *MultiValue) GetID() string {
	return c.Component.GetID()
}

func (c *MultiValue) GetType() string {
	return "multivalue"
}

type Gauge struct {
	Pos       lexer.Position
	Type      string      `parser:"@('gauge'|'barometer'|'compass'|'inclinometer'|'raingauge') '('"`
	Component *Component  `parser:"@@"`
	Label     string      `parser:"@String"`
	Unit      *units.Unit `parser:"(@@)?"`
	Axis      *Axis       `parser:"(@@)?"`
	Metrics   *MetricList `parser:"@@ ')'"`
}

func (c *Gauge) AcceptMetric(v api.Metric) bool {
	return c != nil && c.Metrics.AcceptMetric(v)
}

func (c *Gauge) GetID() string {
	return c.Component.GetID()
}

func (c *Gauge) GetType() string {
	return c.Type
}

func (c *Gauge) Convert(v value.Value) (value.Value, error) {
	var err error

	// Convert v to either the specified unit or that of the first metric
	if c.Unit == nil {
		// This is safe as the parser ensures that metrics contain at least 1 metric for gauges
		v, err = c.Metrics.Metrics[0].Convert(v)
	} else {
		v, err = c.Unit.Convert(v)
	}

	return v, errors.Error(c.Pos, err)
}

func (c *Gauge) ConvertAll(vals []value.Value) ([]value.Value, error) {
	var err error

	if c.Unit != nil {
		for i, v := range vals {
			vals[i], err = c.Unit.Convert(v)
			if err != nil {
				return nil, errors.Error(c.Metrics.Metrics[i].Pos, err)
			}
		}
	}

	return vals, nil
}
