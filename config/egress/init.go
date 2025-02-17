package egress

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/amqp"
	"strings"
)

func NewParser() util.Parser[Script] {
	return util.NewParser[Script](nil, nil, egressInit)
}

func egressInit(_ util.Parser[Script], q *Script, err error) (*Script, error) {
	if err == nil {
		state := NewState()

		err = NewBuilder[*State]().
			Amqp(defineAmqp).
			Metric(defineMetric).
			Publish(definePublish).
			Build().
			SetData(state).
			Script(q)

		if err == nil {
			q.state = state.Cleanup()
		}
	}
	return q, err
}

func defineAmqp(v EgressVisitor[*State], a *amqp.Amqp) error {
	err := v.GetData().AddAmqp(a)
	if err != nil {
		return err
	}
	return errors.VisitorStop
}

func defineMetric(v EgressVisitor[*State], a *Metric) error {
	// Should never occur as this would be a parser error
	if len(a.Metrics) == 0 {
		return participle.Errorf(a.Pos, "metric undefined")
	}

	state := v.GetData()

	for i, m := range a.Metrics {
		m = strings.TrimSpace(m)
		if m == "" {
			return participle.Errorf(a.Pos, "metric undefined")
		}
		a.Metrics[i] = m
		state.AddMetric(m, a)
	}

	if a.Statement != nil {
		return state.scriptInit.Statements(a.Statement)
	}

	return nil
}

func definePublish(v EgressVisitor[*State], a *Publish) error {
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
