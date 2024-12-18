package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/time"
	"strings"
)

type Dashboard struct {
	Pos        lexer.Position
	Name       string              `parser:"'dashboard' '(' @String"`
	Live       bool                `parser:"@'live'?"`
	Update     time.CronTab        `parser:"('update' @@)?"`
	Component  *Component          `parser:"@@"`
	Components *ComponentListEntry `parser:"@@? ')'"`
}

func (c *visitor[T]) Dashboard(d *Dashboard) error {
	var err error
	if d != nil {
		if c.dashboard != nil {
			err = c.dashboard(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Component(d.Component)
		}

		if err == nil {
			err = c.ComponentListEntry(d.Components)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initDashboard(v Visitor[*initState], d *Dashboard) error {
	s := v.Get()

	var err error

	// sensorPrefix is not used for dashboards
	s.sensorPrefix = ""
	if s.stationPrefix == "" {
		// should never occur
		err = errors.Errorf(d.Pos, "stationPrefix not defined")
	}

	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))
	if err == nil && d.Name == "" {
		err = errors.Errorf(d.Pos, "dashboard name is required")
	}
	if err == nil && strings.ContainsAny(d.Name, ". _") {
		err = errors.Errorf(d.Pos, "dashboard name must not contain '.', '_' or spaces")
	}

	// Check name is unique
	if err == nil {
		if e, exists := s.dashboards[d.Name]; exists {
			err = errors.Errorf(d.Pos, "dashboard %q already exists at %s", d.Name, e.String())
		}
	}

	if err == nil {
		// Ensure we have an entry present so we don't need to check this in templates
		if d.Components == nil {
			d.Components = &ComponentListEntry{}
		}

		s.dashboards[d.Name] = d.Pos

		// Ensure Component exists, require by templates
		if d.Component == nil {
			d.Component = &Component{}
		}
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) Dashboard(f func(Visitor[T], *Dashboard) error) Builder[T] {
	b.dashboard = f
	return b
}

func (c *Dashboard) GetType() string {
	return "dashboard"
}
