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

		err = NewBuilder[*State]().
			Amqp(defineAmqp).
			Metric(defineMetric).
			Publish(definePublish).
			Build().
			SetData(state).
			Script(q)

		if err == nil {
			q.state = state
		}
	}
	return q, err
}

func defineAmqp(v Visitor[*State], a *Amqp) error {
	err := v.GetData().AddAmqp(a)
	if err != nil {
		return err
	}
	return VisitorStop
}

func defineMetric(v Visitor[*State], a *Metric) error {
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
		v.GetData().AddMetric(m, a)
	}

	return nil
}

func definePublish(v Visitor[*State], a *Publish) error {
	state := v.GetData()
	a.Amqp = strings.TrimSpace(a.Amqp)
	switch {
	case a.Console,
		a.Amqp != "" && state.GetAmqp(a.Amqp) != nil:
		return nil
	default:
		return participle.Errorf(a.Pos, "invalid publisher")
	}
}
