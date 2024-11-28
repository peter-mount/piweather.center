package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/location"
	"strings"
)

type Stations struct {
	Pos      lexer.Position
	Stations []*Station `parser:"@@+"`
}

// Merge two Stations into a single instance
func (a *Stations) Merge(b *Stations) (*Stations, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	m := make(map[string]*Station)
	for _, s := range a.Stations {
		m[strings.ToLower(s.Name)] = s
	}
	for _, s := range b.Stations {
		if err := assertStationUnique(&m, s); err != nil {
			return nil, err
		}
		a.Stations = append(a.Stations, s)
	}

	return a, nil
}

type Station struct {
	Pos        lexer.Position
	Name       string             `parser:"'station' '(' @String"`
	Location   *location.Location `parser:"@@?"`
	Dashboards *DashboardList     `parser:"@@ ')'"`
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
