package renderer

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
)

func Text(v station.Visitor[*State], d *station.Text) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			s.Builder().TextNbsp(d.Text).End()
			return nil
		})

	if err != nil {
		return err
	}

	return errors.VisitorStop
}
