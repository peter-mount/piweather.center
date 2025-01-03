package ql

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor[T any] interface {
	AliasedExpression(*AliasedExpression) error
	AliasedGroup(*AliasedGroup) error
	Duration(*time.Duration) error
	Expression(*Expression) error
	ExpressionModifier(*ExpressionModifier) error
	Function(*Function) error
	Histogram(*Histogram) error
	Metric(*Metric) error
	Query(*Query) error
	QueryRange(*QueryRange) error
	Select(*Select) error
	SelectExpression(*SelectExpression) error
	Summarize(*Summarize) error
	TableSelect(*TableSelect) error
	Time(*time.Time) error
	TimeZone(*time.TimeZone) error
	Unit(*units.Unit) error
	UsingDefinitions(*UsingDefinitions) error
	UsingDefinition(*UsingDefinition) error
	WindRose(*WindRose) error

	Clone() Visitor[T]
	Set(T) Visitor[T]
	Get() T
}

type visitor[T any] struct {
	common[T]
	data T
}

func (v *visitor[T]) Clone() Visitor[T] {
	return &visitor[T]{common: v.common}
}

func (v *visitor[T]) Get() T {
	return v.data
}

func (v *visitor[T]) Set(data T) Visitor[T] {
	v.data = data
	return v
}

type common[T any] struct {
	aliasedExpression  func(Visitor[T], *AliasedExpression) error
	aliasedGroup       func(Visitor[T], *AliasedGroup) error
	expression         func(Visitor[T], *Expression) error
	expressionModifier func(Visitor[T], *ExpressionModifier) error
	duration           func(Visitor[T], *time.Duration) error
	function           func(Visitor[T], *Function) error
	histogram          func(Visitor[T], *Histogram) error
	metric             func(Visitor[T], *Metric) error
	query              func(Visitor[T], *Query) error
	queryRange         func(Visitor[T], *QueryRange) error
	_select            func(Visitor[T], *Select) error
	selectExpression   func(Visitor[T], *SelectExpression) error
	summarize          func(Visitor[T], *Summarize) error
	tableSelect        func(Visitor[T], *TableSelect) error
	time               func(Visitor[T], *time.Time) error
	timeZone           func(Visitor[T], *time.TimeZone) error
	unit               func(Visitor[T], *units.Unit) error
	usingDefinition    func(Visitor[T], *UsingDefinition) error
	usingDefinitions   func(Visitor[T], *UsingDefinitions) error
	windRose           func(Visitor[T], *WindRose) error
}

func (v *visitor[T]) Duration(b *time.Duration) error {
	if b != nil && v.duration != nil {
		return v.duration(v, b)
	}
	return nil
}

func (v *visitor[T]) TimeZone(b *time.TimeZone) error {
	var err error
	if b != nil {
		if v.timeZone != nil {
			err = v.timeZone(v, b)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}
	}
	return err
}

func (v *visitor[T]) Time(b *time.Time) error {
	var err error
	if b != nil {
		if v.time != nil {
			err = v.time(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
		for _, e := range b.Expression {
			if err == nil {
				err = v.Duration(e.Add)
			}
			if err == nil {
				err = v.Duration(e.Truncate)
			}
		}
	}
	return err
}

func (v *visitor[T]) Unit(b *units.Unit) error {
	if b != nil && v.unit != nil {
		return v.unit(v, b)
	}
	return nil
}
