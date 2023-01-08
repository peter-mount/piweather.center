package ephemeris

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sun"
	"github.com/soniakeys/meeus/v3/rise"
)

type SunProvider struct {
	meta Meta
}

func (s *SunProvider) Init(e *Ephemeris) error {
	s.meta = e.Meta
	return nil

}
func (s *SunProvider) Generate(day julian.Day) (*Entry, error) {

	eq := sun.ApparentEquatorial(day)

	// Add angular diameter of the sun
	eq.Diameter = rise.Stdh0Solar

	return &Entry{
		Name:       "Sun",
		Ord:        1,
		Date:       day,
		Equatorial: &eq,
	}, nil
}
