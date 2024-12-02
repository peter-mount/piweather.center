package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/api"
)

type Text struct {
	Pos       lexer.Position
	Type      string     `parser:"@('text') '('"`
	Component *Component `parser:"@@"`
	Text      string     `parser:"@String ')'"`
}

func (c *Text) GetID() string {
	return c.Component.GetID()
}

func (c *Text) GetType() string {
	return c.Type
}

type Forecast struct {
	Pos           lexer.Position
	Type          string     `parser:"@('forecast') '('"`
	Component     *Component `parser:"@@"`
	Temperature   *Metric    `parser:"@@"`
	Pressure      *Metric    `parser:"@@"`
	WindDirection *Metric    `parser:"@@ ')'"`
}

func (c *Forecast) GetID() string {
	return c.Component.GetID()
}

func (c *Forecast) GetType() string {
	return c.Type
}

func (c *Forecast) AcceptMetric(v api.Metric) bool {
	return c != nil && (c.Pressure.AcceptMetric(v) || c.WindDirection.AcceptMetric(v))
}
