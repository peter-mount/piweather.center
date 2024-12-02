package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/store/api"
)

type ComponentList struct {
	Pos     lexer.Position
	Entries []*ComponentListEntry `parser:"@@*"`
}

type ComponentListEntry struct {
	Pos        lexer.Position
	Container  *Container  `parser:"( @@"`
	Gauge      *Gauge      `parser:"| @@"`
	MultiValue *MultiValue `parser:"| @@"`
	Value      *Value      `parser:"| @@ )"`
}

func (c *ComponentListEntry) AcceptMetric(v api.Metric) bool {
	switch {
	case c.Gauge != nil:
		return c.Gauge.AcceptMetric(v)
	case c.MultiValue != nil:
		return c.MultiValue.AcceptMetric(v)
	case c.Value != nil:
		return c.Value.AcceptMetric(v)
	default:
		return false
	}
}

func (c *ComponentListEntry) GetID() string {
	switch {
	case c.Gauge != nil:
		return c.Gauge.Component.GetID()
	case c.MultiValue != nil:
		return c.MultiValue.Component.GetID()
	case c.Value != nil:
		return c.Value.Component.GetID()
	default:
		return ""
	}
}

func (c *ComponentListEntry) GetType() string {
	switch {
	case c.Container != nil:
		return c.Container.GetType()
	case c.Gauge != nil:
		return c.Gauge.GetType()
	case c.MultiValue != nil:
		return c.MultiValue.GetType()
	case c.Value != nil:
		return c.Value.GetType()
	default:
		return ""
	}
}

type Container struct {
	Pos        lexer.Position
	Type       string         `parser:"@('container' | 'col' | 'row') '('"`
	Component  *Component     `parser:"@@"`
	Components *ComponentList `parser:"@@ ')'"`
}

func (c *Container) GetType() string {
	return c.Type
}

type Dashboard struct {
	Pos        lexer.Position
	Name       string              `parser:"'dashboard' '(' @String"`
	Live       bool                `parser:"@'live'?"`
	Update     *time.CronTab       `parser:"('update' @@)?"`
	Component  *Component          `parser:"@@"`
	Components *ComponentListEntry `parser:"@@? ')'"`
}

func (c *Dashboard) GetType() string {
	return "dashboard"
}
