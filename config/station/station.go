package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
)

type Stations struct {
	Pos      lexer.Position
	Stations []*Station `parser:"@@+"`
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

type Station struct {
	util.CheckSum
	Pos      lexer.Position
	Name     string             `parser:"'station' '(' @String"`
	Location *location.Location `parser:"@@?"`

	Dashboards *DashboardList `parser:"@@ ')'"`
}

// HomeDashboard returns the home dashboard which is by definition the first dashboard in the definition
func (s *Station) HomeDashboard() *Dashboard {
	if s.Dashboards == nil || len(s.Dashboards.Dashboards) == 0 {
		return nil
	}
	return s.Dashboards.Dashboards[0]
}

type DashboardList struct {
	Pos        lexer.Position
	Dashboards []*Dashboard `parser:"@@*"`
}
