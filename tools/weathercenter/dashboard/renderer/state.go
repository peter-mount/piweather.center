package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	state2 "github.com/peter-mount/piweather.center/tools/weathercenter/dashboard/state"
	"github.com/peter-mount/piweather.center/util/html"
)

type State struct {
	parent    *State
	dashboard *state2.Dashboard
	b         *html.Element
}

func NewState(dashboard *state2.Dashboard) *State {
	return &State{
		dashboard: dashboard,
		b:         html.Builder(),
	}
}

func (s *State) Builder() *html.Element {
	return s.b
}

func (s *State) Dashboard() *state2.Dashboard {
	return s.dashboard
}

func (s *State) IsLive() bool {
	return s.dashboard.Dashboard().Live
}

func (s *State) String() string {
	return s.b.String()
}

type StateHandler func(*State) error

func (s *State) With(v station.Visitor[*State], e *html.Element, f StateHandler) error {
	c := &State{
		parent:    s,
		dashboard: s.dashboard,
		b:         e,
	}
	defer func() {
		v.Set(s)
	}()
	v.Set(c)
	return f(c)
}

func (s *State) Component(v station.Visitor[*State], d station.ComponentType, c *station.Component, f StateHandler) error {
	e := s.Builder()

	if c.Title != "" {
		e = e.Div().Class("dash-title-box").
			Div().Class("dash-title-title").
			Text(c.Title).
			End()
	}

	e = e.Div().
		Class("dash-%s", d.GetType()).
		Class(c.Class).
		Attr("style", c.Style)

	err := s.With(v, e, f)

	if err == nil {
		// Close the enclosing dash-{type} div
		e = e.End()

		// Close the enclosing dash-title-box div
		if c.Title != "" {
			e = e.End()
		}
	}

	return err
}
