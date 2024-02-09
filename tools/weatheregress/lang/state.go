package lang

import (
	"github.com/alecthomas/participle/v2"
	"strings"
	"sync"
)

type State struct {
	mutex        sync.Mutex
	amqp         map[string]*Amqp     // Map of AMQP definitions
	metricMatch  map[string][]*Metric // Map of exact matches
	metricFilter []MetricFilter       // Slice of filters to filter against
}

type MetricFilter func(string) *Metric

func (s *State) GetAmqp(n string) *Amqp {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.amqp[strings.ToLower(n)]
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

	for k, v := range b.metricMatch {
		s.metricMatch[k] = append(s.metricMatch[k], v...)
	}

	s.metricFilter = append(s.metricFilter, b.metricFilter...)
}

func (s *State) AddAmqp(a *Amqp) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	a.Name = strings.ToLower(strings.TrimSpace(a.Name))

	if e, exists := s.amqp[a.Name]; exists {
		return participle.Errorf(a.Pos, "%q already defined at %s",
			a.Name,
			e.Pos.String())
	}

	// Exchange is optional, default to amq.topic
	a.Exchange = strings.TrimSpace(a.Exchange)
	if a.Exchange == "" {
		a.Exchange = "amq.topic"
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
