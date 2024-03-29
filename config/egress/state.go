package egress

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/parser"
	amqp2 "github.com/peter-mount/piweather.center/config/util/amqp"
	"strings"
	"sync"
)

type State struct {
	mutex        sync.Mutex
	amqp         map[string]*amqp2.Amqp // Map of AMQP definitions
	metricMatch  map[string][]*Metric   // Map of exact matches
	metricFilter []MetricFilter         // Slice of filters to filter against
	scriptInit   parser.Initialiser     // from go-script, initialiser for embedded scripts
}

func NewState() *State {
	return &State{
		amqp:        make(map[string]*amqp2.Amqp),
		metricMatch: make(map[string][]*Metric),
		scriptInit:  parser.NewInitialiser(),
	}
}

func (s *State) Cleanup() *State {
	s.scriptInit = nil
	return s
}

type MetricFilter func(string) *Metric

func (s *State) GetAmqp(n string) *amqp2.Amqp {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.amqp[strings.ToLower(n)]
}

func (s *State) GetAmqpNames() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var a []string
	for k, _ := range s.amqp {
		a = append(a, k)
	}
	return a
}

func (s *State) GetMetrics(n string) []*Metric {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Exact match
	if a, exists := s.metricMatch[n]; exists {
		return a
	}

	// Slower by filters
	var a []*Metric
	for _, f := range s.metricFilter {
		if m := f(n); m != nil {
			a = append(a, m)
		}
	}
	return a
}

func (s *State) merge(b *State) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for k, v := range b.amqp {
		if _, exists := s.amqp[k]; !exists {
			s.amqp[k] = v
		}
	}

	for k, v := range b.metricMatch {
		s.metricMatch[k] = append(s.metricMatch[k], v...)
	}

	s.metricFilter = append(s.metricFilter, b.metricFilter...)
}

func (s *State) AddAmqp(a *amqp2.Amqp) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := a.Init(); err != nil {
		return err
	}

	if e, exists := s.amqp[a.Name]; exists {
		return participle.Errorf(a.Pos, "%q already defined at %s",
			a.Name,
			e.Pos.String())
	}
	s.amqp[a.Name] = a

	return nil
}

func (s *State) AddMetric(n string, m *Metric) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var f func(string, string) bool

	hasPrefix := strings.HasPrefix(n, "*")
	hasSuffix := strings.HasSuffix(n, "*")
	n = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(n, "*"), "*"))

	switch {
	case hasPrefix && hasSuffix:
		f = strings.Contains
	case hasPrefix:
		f = strings.HasSuffix
	case hasSuffix:
		f = strings.HasPrefix
	default:
		s.metricMatch[n] = append(s.metricMatch[n], m)
		return
	}

	s.metricFilter = append(s.metricFilter, func(s string) *Metric {
		if f(s, n) {
			return m
		}
		return nil
	})
}
