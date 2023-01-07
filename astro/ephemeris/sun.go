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
func (s *SunProvider) Generate(day julian.Day) (Entry, error) {

	th0 := day.Apparent0UT()

	eq := sun.ApparentEquatorial(day)
	rs := eq.RiseSet(s.meta.LatLong.Coord, th0, rise.Stdh0Solar)

	return Entry{
		Date:       day,
		Equatorial: &eq,
		RiseSet:    &rs,
	}, nil
}
