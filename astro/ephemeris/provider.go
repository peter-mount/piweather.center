package ephemeris

import (
	"github.com/peter-mount/piweather.center/astro/julian"
)

type Provider interface {
	Init(e *Ephemeris) error
	Generate(day julian.Day) (*Entry, error)
}

func Generate(start, end julian.Day, step float64, p Provider) (*Ephemeris, error) {
	ep := &Ephemeris{}
	ep.IncludePeriod(start, end)

	if err := ep.Generate(step, p); err != nil {
		return nil, err
	}
	return ep, nil
}

func (e *Ephemeris) Generate(step float64, p Provider) error {
	if err := p.Init(e); err != nil {
		return err
	}

	return e.Range.ForEach(step, func(day julian.Day) error {
		entry, err := p.Generate(day)
		if err != nil {
			return err
		}
		if entry.Date == day {
			e.Entries = append(e.Entries, entry)
		}
		return nil
	})
}
