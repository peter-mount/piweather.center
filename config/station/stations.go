package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Stations struct {
	Pos      lexer.Position
	Stations []*Station `parser:"@@*"`
}

func (c *visitor[T]) Stations(d *Stations) error {
	var err error
	if d != nil {
		if c.stations != nil {
			err = c.stations(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, s := range d.Stations {
				err = c.Station(s)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initStations(v Visitor[*initState], d *Stations) error {
	v.Get().stationIds = make(map[string]*Station)
	return nil
}

func (b *builder[T]) Stations(f func(Visitor[T], *Stations) error) Builder[T] {
	b.stations = f
	return b
}

func (a *Stations) GetChecksum() [20]byte {
	panic("not implemented")
}

func (a *Stations) SetChecksum(checksum [20]byte) {
	// Set the checksum against the individual station
	// Done as we merge Stations so the checksum would be for the wrong file
	for _, s := range a.Stations {
		s.SetChecksum(checksum)
	}
}

// Merge two Stations into a single instance
func (a *Stations) Merge(b *Stations) (*Stations, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	a.Stations = append(a.Stations, b.Stations...)
	return a, nil
}

func (a *Stations) Replace(b *Stations) (*Stations, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	// Work against a new instance so if we fail we don't alter the existing working copy
	c := &Stations{Pos: lexer.Position{}}

	m := make(map[string]*Station)
	for _, ae := range a.Stations {
		m[ae.Name] = ae
		c.Stations = append(c.Stations, ae)
	}

	for _, be := range b.Stations {
		if err := a.replace(&m, c, be); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (a *Stations) Remove(p string) *Stations {
	if a == nil {
		return nil
	}

	c := &Stations{Pos: lexer.Position{}}
	for _, e := range a.Stations {
		if e.Pos.Filename != p {
			c.Stations = append(c.Stations, e)
		}
	}

	return c
}

func (a *Stations) replace(m *map[string]*Station, c *Stations, e *Station) error {
	for ai, ae := range c.Stations {
		if ae.Name == e.Name {
			c.Stations[ai] = e
			return nil
		}
	}

	if err := assertStationUnique(m, e); err != nil {
		return err
	}

	c.Stations = append(c.Stations, e)

	return nil
}
