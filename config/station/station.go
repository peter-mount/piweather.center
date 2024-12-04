package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"strings"
)

type Station struct {
	util.CheckSum
	Pos          lexer.Position
	Name         string             `parser:"'station' '(' @String"`
	Location     *location.Location `parser:"@@?"`
	Calculations *CalculationList   `parser:"@@"`
	Sensors      *SensorList        `parser:"@@"`
	Dashboards   *DashboardList     `parser:"@@ ')'"`
}

func (c *visitor[T]) Station(d *Station) error {
	var err error
	if d != nil {
		if c.station != nil {
			err = c.station(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Location(d.Location)
		}

		if err == nil {
			err = c.CalculationList(d.Calculations)
		}

		if err == nil {
			err = c.DashboardList(d.Dashboards)
		}

		if err == nil {
			err = c.SensorList(d.Sensors)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func initStation(v Visitor[*initState], d *Station) error {
	s := v.Get()
	// Reset the state
	s.calculations = nil
	s.dashboards = nil
	s.sensors = nil
	s.sensorPrefix = ""
	s.stationPrefix = ""

	var err error

	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		err = errors.Errorf(d.Pos, "station id is required")
	}

	if err == nil && strings.ContainsAny(d.Name, ". _") {
		err = errors.Errorf(d.Pos, "station id must not contain '.', '_' or spaces")
	}

	if err == nil {
		s.stationId = d.Name
		s.stationPrefix = s.stationId + "."

		// Ensure we have a Location
		if d.Location == nil {
			// This will place the station at Null Island
			d.Location = &location.Location{Pos: d.Pos}
		}

		// Ensure stationId is unique
		err = assertStationUnique(&s.stationIds, d)
	}

	return errors.Error(d.Pos, err)
}

func (b *builder[T]) Station(f func(Visitor[T], *Station) error) Builder[T] {
	b.station = f
	return b
}

func assertStationUnique(m *map[string]*Station, s *Station) error {
	n := strings.ToLower(s.Name)
	if e, exists := (*m)[n]; exists {
		return errors.Errorf(s.Pos, "station %q already defined at %s", s.Name, e.Pos)
	}
	(*m)[n] = s
	return nil
}

// HomeDashboard returns the home dashboard which is by definition the first dashboard in the definition
func (s *Station) HomeDashboard() *Dashboard {
	if s.Dashboards == nil || len(s.Dashboards.Dashboards) == 0 {
		return nil
	}
	return s.Dashboards.Dashboards[0]
}
