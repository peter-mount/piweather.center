package calc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"sync"
)

func (p *defaultParser) init(q *Script, err error) (*Script, error) {
	if err == nil {
		state := &State{
			locations:    make(map[string]*Location),
			calculations: make(map[string]*Calculation),
		}

		err = NewBuilder[*State]().
			Calculation(state.initCalculation).
			Load(state.load).
			CronTab(state.initCronTab).
			Function(state.initFunction).
			Location(state.initLocation).
			Metric(state.initMetric).
			Unit(state.initUnit).
			Build().
			SetData(state).
			Script(q)

		if err == nil {
			q.State = state
		}
	}
	return q, err
}

type State struct {
	mutex        sync.Mutex
	locations    map[string]*Location
	calculations map[string]*Calculation
}

func (s *State) GetLocation(n string) *Location {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.locations[n]
}

func (s *State) GetLocations() []*Location {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var r []*Location
	for _, l := range s.locations {
		r = append(r, l)
	}
	return r
}

func (s *State) GetCalculation(n string) *Calculation {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.calculations[n]
}

func (s *State) GetCalculations() []*Calculation {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var r []*Calculation
	for _, v := range s.calculations {
		r = append(r, v)
	}
	return r
}

func (s *State) initCalculation(_ Visitor[*State], l *Calculation) error {
	l.Target = strings.ToLower(l.Target)
	l.At = strings.ToLower(l.At)

	if e, exists := s.calculations[l.Target]; exists {
		return participle.Errorf(l.Pos, "calculation for %q already defined at %s", l.Target, e.Pos.String())
	}

	s.calculations[l.Target] = l
	return nil
}

func (s *State) load(_ Visitor[*State], l *Load) error {
	l.When = strings.ToLower(l.When)
	l.With = strings.TrimSpace(l.With)

	switch l.When {
	case "today", "hour", "minute":
	default:
		return participle.Errorf(l.Pos, "Unsupported When %q", l.When)
	}

	if l.With == "" {
		return participle.Errorf(l.Pos, "Undefined With")
	}

	return nil
}

func (s *State) initCronTab(_ Visitor[*State], l *CronTab) error {
	// Convert aliases to actual definitions
	switch strings.ToLower(l.Definition) {
	case "day", "daily", "midnight":
		l.Definition = "0 0 0 * * *"
	case "hour", "hourly":
		l.Definition = "0 0 * * * *"
	case "minute":
		l.Definition = "0 * * * * *"
	case "second":
		l.Definition = "* * * * * *"
	}

	return nil
}

func (s *State) initFunction(_ Visitor[*State], l *Function) error {
	l.Name = strings.ToLower(l.Name)

	if !value.CalculatorExists(l.Name) {
		return participle.Errorf(l.Pos, "function %q is undefined", l.Name)
	}

	return nil
}

func (s *State) initLocation(_ Visitor[*State], l *Location) error {
	l.Name = strings.ToLower(l.Name)

	lat, err := util.ParseAngle(l.Latitude)
	if err != nil {
		return err
	}

	lon, err := util.ParseAngle(l.Longitude)
	if err != nil {
		return err
	}

	l.latLong = &coord.LatLong{
		Longitude: lon,
		Latitude:  lat,
		Altitude:  l.Altitude,
		Name:      l.Name,
	}

	if e, exists := s.locations[l.Name]; exists {
		return participle.Errorf(l.Pos, "location %q already defined at %s", l.Name, e.Pos.String())
	}

	s.locations[l.Name] = l
	return nil
}

func (s *State) initMetric(_ Visitor[*State], l *Metric) error {
	l.Name = strings.ToLower(strings.Join(l.Metric, "."))
	return nil
}

func (s *State) initUnit(_ Visitor[*State], l *Unit) error {
	u, exists := value.GetUnit(l.Using)
	if exists {
		l.unit = u
		return nil
	}
	return participle.Errorf(l.Pos, "unsupported unit %q", l.Using)
}
