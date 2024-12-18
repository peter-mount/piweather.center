package renderer

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
)

func Dashboard(v station.Visitor[*State], d *station.Dashboard) error {
	s := v.Get()

	e := s.Builder().
		Div().Class("dashboard-outer")

	err := s.With(v, e, func(s *State) error {
		return s.Component(v, d, d.Component, func(_ *State) error {
			return v.ComponentListEntry(d.Components)
		})
	})

	if err != nil {
		return err
	}

	if s.IsLive() {
		stn := s.Dashboard().Station()
		script := s.GenerateJavaScript(stn.Station().Name, d.Name, stn.GetUid())
		if script != "" {
			s.Builder().Script().Text(script).End()
		}
	}

	return errors.VisitorStop
}
