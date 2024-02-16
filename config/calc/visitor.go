package calc

import (
	"errors"
	"github.com/peter-mount/piweather.center/config/data"
	"github.com/peter-mount/piweather.center/config/location"
)

type Visitor[T any] interface {
	data.DataVisitor[T]
	location.LocationVisitor[T]
	Clone() Visitor[T]
	Calculation(*Calculation) error
	CronTab(*CronTab) error
	Current(*Current) error
	Expression(*Expression) error
	Function(*Function) error
	Load(b *Load) error
	Metric(*Metric) error
	Script(*Script) error
	Unit(b *Unit) error
	UseFirst(b *UseFirst) error
}

type visitorCommon[T any] struct {
	calculation func(Visitor[T], *Calculation) error
	cronTab     func(Visitor[T], *CronTab) error
	current     func(Visitor[T], *Current) error
	expression  func(Visitor[T], *Expression) error
	function    func(Visitor[T], *Function) error
	load        func(Visitor[T], *Load) error
	metric      func(Visitor[T], *Metric) error
	script      func(Visitor[T], *Script) error
	unit        func(Visitor[T], *Unit) error
	useFirst    func(Visitor[T], *UseFirst) error
}

type visitor[T any] struct {
	data.DataVisitorCommon[T]
	location.LocationVisitorBase[T]
	visitorCommon[T]
}

func (v *visitor[T]) Clone() Visitor[T] {
	return &visitor[T]{
		DataVisitorCommon:   v.DataVisitorCommon,
		LocationVisitorBase: v.LocationVisitorBase,
		visitorCommon:       v.visitorCommon,
	}
}

// VisitorStop is an error which causes the current step in a Visitor to stop processing.
// It's used to enable a Visitor to handle all processing of a node within itself rather
// than the Visitor proceeding to any child nodes of that node.
var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

func (v *visitor[T]) Script(b *Script) error {
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
					err = v.Location(l)
				}
			}
		}

		if err == nil {
			for _, c := range b.Calculations {
				if err == nil {
					err = v.Calculation(c)
				}
			}
		}
	}
	return err
}

func (v *visitor[T]) Calculation(b *Calculation) error {
	var err error
	if b != nil {
		if v.calculation != nil {
			err = v.calculation(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = v.CronTab(b.Every)
		}

		if err == nil {
			err = v.CronTab(b.ResetEvery)
		}

		if err == nil {
			err = v.Load(b.Load)
		}

		if err == nil {
			err = v.UseFirst(b.UseFirst)
		}

		if err == nil {
			err = v.Expression(b.Expression)
		}
	}
	return err
}

func (v *visitor[T]) Load(b *Load) error {
	var err error
	if b != nil && v.load != nil {
		err = v.load(v, b)
	}
	if IsVisitorStop(err) {
		return nil
	}
	return err
}

func (v *visitor[T]) CronTab(b *CronTab) error {
	var err error
	if b != nil && v.cronTab != nil {
		err = v.cronTab(v, b)
	}
	if IsVisitorStop(err) {
		return nil
	}
	return err
}

func (v *visitor[T]) Expression(b *Expression) error {
	var err error
	if b != nil {
		if v.expression != nil {
			err = v.expression(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil && b.Current != nil {
			err = v.Current(b.Current)
		}

		if err == nil && b.Function != nil {
			err = v.Function(b.Function)
		}

		if err == nil && b.Metric != nil {
			err = v.Metric(b.Metric)
		}

		if err == nil && b.Using != nil {
			err = v.Unit(b.Using)
		}
	}
	return err
}

func (v *visitor[T]) Current(b *Current) error {
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

func (v *visitor[T]) Unit(b *Unit) error {
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

func (v *visitor[T]) Function(b *Function) error {
	var err error
	if b != nil {
		if v.function != nil {
			err = v.function(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			for _, exp := range b.Expressions {
				if err == nil {
					err = v.Expression(exp)
				}
			}
		}
	}
	return err
}

func (v *visitor[T]) Metric(b *Metric) error {
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

func (v *visitor[T]) UseFirst(b *UseFirst) error {
	var err error
	if b != nil {
		if v.useFirst != nil {
			err = v.useFirst(v, b)
		}
		if IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = v.Metric(b.Metric)
		}
	}
	return err
}
