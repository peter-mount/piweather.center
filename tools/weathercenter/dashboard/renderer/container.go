package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
)

func Container(v station.Visitor[*State], d *station.Container) error {
	err := v.Get().
		Component(v, d, d.Component, func(_ *State) error {
			return v.ComponentList(d.Components)
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}
