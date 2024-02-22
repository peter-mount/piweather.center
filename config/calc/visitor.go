package calc

import (
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/data"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type CalcVisitor[T any] interface {
	data.DataVisitor[T]
	location.LocationVisitor[T]
	Clone() CalcVisitor[T]
	Calculation(*Calculation) error
	CronTab(tab *time.CronTab) error
	Current(*Current) error
	Expression(*Expression) error
	Function(*Function) error
	Load(b *Load) error
	Metric(*Metric) error
	Script(*Script) error
	Unit(b *units.Unit) error
	UseFirst(b *UseFirst) error
}

type visitorCommon[T any] struct {
	calculation func(CalcVisitor[T], *Calculation) error
	cronTab     func(CalcVisitor[T], *time.CronTab) error
	current     func(CalcVisitor[T], *Current) error
	expression  func(CalcVisitor[T], *Expression) error
	function    func(CalcVisitor[T], *Function) error
	load        func(CalcVisitor[T], *Load) error
	metric      func(CalcVisitor[T], *Metric) error
	script      func(CalcVisitor[T], *Script) error
	unit        func(CalcVisitor[T], *units.Unit) error
	useFirst    func(CalcVisitor[T], *UseFirst) error
}

type visitor[T any] struct {
	data.DataVisitorCommon[T]
	location.LocationVisitorBase[T]
	visitorCommon[T]
}

func (v *visitor[T]) Clone() CalcVisitor[T] {
	return &visitor[T]{
		DataVisitorCommon:   v.DataVisitorCommon,
		LocationVisitorBase: v.LocationVisitorBase,
		visitorCommon:       v.visitorCommon,
	}
}

func (v *visitor[T]) Script(b *Script) error {
	var err error
	if b != nil {
		if v.script != nil {
			err = v.script(v, b)
		}
		if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
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
	if util.IsVisitorStop(err) {
		return nil
	}
	return err
}

func (v *visitor[T]) CronTab(b *time.CronTab) error {
	var err error
	if b != nil && v.cronTab != nil {
		err = v.cronTab(v, b)
	}
	if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
			return nil
		}
	}
	return err
}

func (v *visitor[T]) Unit(b *units.Unit) error {
	var err error
	if b != nil {
		if v.unit != nil {
			err = v.unit(v, b)
		}
		if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
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
		if util.IsVisitorStop(err) {
			return nil
		}

		if err == nil {
			err = v.Metric(b.Metric)
		}
	}
	return err
}
