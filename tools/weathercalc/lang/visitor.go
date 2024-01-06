package lang

import "errors"

type Visitor interface {
	Calculation(*Calculation) error
	CronTab(*CronTab) error
	Current(*Current) error
	Expression(*Expression) error
	Function(*Function) error
	Location(*Location) error
	Metric(*Metric) error
	Script(*Script) error
	Unit(b *Unit) error
	UseFirst(b *Metric) error
}

type Visitable interface {
	Accept(v Visitor) error
}

type visitorCommon struct {
	calculation func(Visitor, *Calculation) error
	cronTab     func(Visitor, *CronTab) error
	current     func(Visitor, *Current) error
	expression  func(Visitor, *Expression) error
	function    func(Visitor, *Function) error
	location    func(Visitor, *Location) error
	metric      func(Visitor, *Metric) error
	script      func(Visitor, *Script) error
	unit        func(Visitor, *Unit) error
	useFirst    func(Visitor, *Metric) error
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
			for _, l := range b.Locations {
				if err == nil {
					err = l.Accept(v)
				}
			}
		}

		if err == nil {
			for _, c := range b.Calculations {
				if err == nil {
					err = c.Accept(v)
				}
			}
		}
	}
	return err
}

func (v *visitor) Location(b *Location) error {
	var err error
	if b != nil && v.location != nil {
		err = v.location(v, b)
	}
	if IsVisitorStop(err) {
		return nil
	}
	return err
}

func (v *visitor) Calculation(b *Calculation) error {
	var err error
	if b != nil {
		if v.calculation != nil {
			err = v.calculation(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = b.Every.Accept(v)
		}

		if err == nil {
			err = b.ResetEvery.Accept(v)
		}

		if err == nil {
			err = b.UseFirst.Accept(v)
		}

		if err == nil {
			err = b.Expression.Accept(v)
		}
	}
	return err
}

func (v *visitor) CronTab(b *CronTab) error {
	var err error
	if b != nil && v.cronTab != nil {
		err = v.cronTab(v, b)
	}
	if IsVisitorStop(err) {
		return nil
	}
	return err
}

func (v *visitor) Expression(b *Expression) error {
	var err error
	if b != nil {
		if v.expression != nil {
			err = v.expression(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil && b.Current != nil {
			err = b.Current.Accept(v)
		}

		if err == nil && b.Function != nil {
			err = b.Function.Accept(v)
		}

		if err == nil && b.Metric != nil {
			err = b.Metric.Accept(v)
		}

		if err == nil && b.Using != nil {
			err = b.Using.Accept(v)
		}
	}
	return err
}

func (v *visitor) Current(b *Current) error {
	var err error
	if b != nil {
		if v.current != nil {
			err = v.current(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor) Unit(b *Unit) error {
	var err error
	if b != nil {
		if v.unit != nil {
			err = v.unit(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor) Function(b *Function) error {
	var err error
	if b != nil {
		if v.function != nil {
			err = v.function(v, b)
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

func (v *visitor) UseFirst(b *Metric) error {
	var err error
	if b != nil {
		if v.useFirst != nil {
			err = v.useFirst(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}
	}
	return err
}
