package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
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

	return util.VisitorStop
}
