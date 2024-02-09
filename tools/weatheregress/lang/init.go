package lang

import (
	"github.com/alecthomas/participle/v2"
	"strings"
)

func (p *defaultParser) init(q *Script, err error) (*Script, error) {
	if err == nil {
		state := &State{
			amqp:        make(map[string]*Amqp),
			metricMatch: make(map[string][]*Metric),
		}

		err = q.Accept(NewBuilder().
			Amqp(state.defineAmqp).
			Metric(state.defineMetric).
			Publish(state.definePublish).
			Build())

		if err == nil {
			q.state = state
		}
	}
	return q, err
}

func (s *State) defineAmqp(v Visitor, a *Amqp) error {
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
	return VisitorStop
}

func (s *State) defineMetric(_ Visitor, a *Metric) error {
	// Should never occur as this would be a parser error
	if len(a.Metrics) == 0 {
		return participle.Errorf(a.Pos, "metric undefined")
	}

	for i, m := range a.Metrics {
		m = strings.TrimSpace(m)
		if m == "" {
			return participle.Errorf(a.Pos, "metric undefined")
		}
		a.Metrics[i] = m
		s.AddMetric(m, a)
	}

	return nil
}

func (s *State) definePublish(_ Visitor, a *Publish) error {
	a.Amqp = strings.TrimSpace(a.Amqp)
	switch {
	case a.Console,
		a.Amqp != "" && s.GetAmqp(a.Amqp) != nil:
		return nil
	default:
		return participle.Errorf(a.Pos, "invalid publisher")
	}
}
