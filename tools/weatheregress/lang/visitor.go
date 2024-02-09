package lang

import "errors"

type Visitor interface {
	Action(a *Action) error
	Amqp(amqp *Amqp) error
	Format(f *Format) error
	Metric(m *Metric) error
	Script(*Script) error
}

type Visitable interface {
	Accept(v Visitor) error
}

type visitorCommon struct {
	action func(Visitor, *Action) error
	amqp   func(Visitor, *Amqp) error
	format func(Visitor, *Format) error
	metric func(Visitor, *Metric) error
	script func(Visitor, *Script) error
}

type visitor struct {
	visitorCommon
}

// VisitorStop is an error which causes the current step in a Visitor to stop processing.
// It's used to enable a Visitor to handle all processing of a node within itself rather
// than the Visitor proceeding to any child nodes of that node.
var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

func (v *visitor) Script(b *Script) error {
	var err error
	if b != nil {
		if v.script != nil {
			err = v.script(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			for _, l := range b.Amqp {
				if err == nil {
					err = l.Accept(v)
				}
			}
		}

		if err == nil {
			for _, c := range b.Actions {
				if err == nil {
					err = c.Accept(v)
				}
			}
		}
	}
	return err
}

func (v *visitor) Action(b *Action) error {
	var err error
	if b != nil {
		if v.action != nil {
			err = v.action(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = b.Metric.Accept(v)
		}
	}
	return err
}

func (v *visitor) Amqp(b *Amqp) error {
	var err error
	if b != nil {
		if v.amqp != nil {
			err = v.amqp(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor) Format(b *Format) error {
	var err error
	if b != nil {
		if v.format != nil {
			err = v.format(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor) Metric(b *Metric) error {
	var err error
	if b != nil {
		if v.metric != nil {
			err = v.metric(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}
