package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	state2 "github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/html"
)

type State struct {
	root      *State            // Root state
	parent    *State            // Parent to this state
	dashboard *state2.Dashboard // Dashboard being built
	b         *html.Element     // Current element
	// These exist in the root State only
	live bool                   // true if live
	js   map[string]interface{} // JavaScript types
}

func NewState(dashboard *state2.Dashboard) *State {
	s := &State{
		dashboard: dashboard,
		live:      dashboard.Dashboard().Live,
		b:         html.Builder(),
	}
	// Root is itself here
	s.root = s
	if s.live {
		s.js = make(map[string]interface{})
	}
	return s
}

func (s *State) Builder() *html.Element {
	return s.b
}

func (s *State) Dashboard() *state2.Dashboard {
	return s.dashboard
}

func (s *State) IsLive() bool {
	return s.root.live
}

func (s *State) String() string {
	return s.b.String()
}

type StateHandler func(*State) error

func (s *State) With(v station.Visitor[*State], e *html.Element, f StateHandler) error {
	c := &State{
		parent:    s,
		root:      s.root,
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
	compType := d.GetType()

	// If live indicate we need to add any applicable javascript for this component
	if s.IsLive() && HasJavaScript(compType) {
		s.root.js[compType] = true
	}

	e := s.Builder()

	if c.Title != "" {
		e = e.Div().Class("dash-title-box").
			Div().Class("dash-title-title").
			Text(c.Title).
			End()
	}

	e = e.Div().
		Class("dash-%s", compType).
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

func (s *State) GenerateJavaScript(stId, dashId, dashUid string) string {
	if s.live {
		return GenerateJavaScript(stId, dashId, dashUid, s.js)
	}
	return ""
}
