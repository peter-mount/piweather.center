package calc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/util"
	location2 "github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"sync"
)

func NewParser() util.Parser[Script] {
	return util.NewParser[Script](nil, nil, calcInit)
}

func calcInit(q *Script, err error) (*Script, error) {
	if err == nil {
		state := &State{
			calculations: make(map[string]*Calculation),
		}

		b := NewBuilder[*State]().
			Calculation(state.initCalculation).
			Load(state.load).
			CronTab(state.initCronTab).
			Function(state.initFunction).
			Metric(state.initMetric).
			Unit(state.initUnit)

		b.Location(state.initLocation)

		v := b.Build()
		v.SetData(state)

		err = v.Script(q)

		if err == nil {
			q.State = state
		}
	}
	return q, err
}

type State struct {
	location2.MapContainer
	mutex        sync.Mutex
	calculations map[string]*Calculation
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

func (s *State) initCalculation(_ CalcVisitor[*State], l *Calculation) error {
	l.Target = strings.ToLower(l.Target)
	l.At = strings.ToLower(l.At)

	if e, exists := s.calculations[l.Target]; exists {
		return participle.Errorf(l.Pos, "calculation for %q already defined at %s", l.Target, e.Pos.String())
	}

	s.calculations[l.Target] = l
	return nil
}

func (s *State) load(_ CalcVisitor[*State], l *Load) error {
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

func (s *State) initCronTab(_ CalcVisitor[*State], l *time.CronTab) error {
	return l.Init()
}

func (s *State) initFunction(_ CalcVisitor[*State], l *Function) error {
	l.Name = strings.ToLower(l.Name)

	if !value.CalculatorExists(l.Name) {
		return participle.Errorf(l.Pos, "function %q is undefined", l.Name)
	}

	return nil
}

func (s *State) initLocation(_ location2.LocationVisitor[*State], l *location2.Location) error {
	if err := l.Init(); err != nil {
		return err
	}

	if _, added := s.SetLocation(l); !added {
		return participle.Errorf(l.Pos, "location %q already defined", l.Name)
	}

	return nil
}

func (s *State) initMetric(_ CalcVisitor[*State], l *Metric) error {
	l.Name = strings.ToLower(strings.Join(l.Metric, "."))
	return nil
}

func (s *State) initUnit(_ CalcVisitor[*State], l *units.Unit) error {
	return l.Init()
}
