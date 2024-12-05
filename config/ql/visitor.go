package ql

import (
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor interface {
	AliasedExpression(*AliasedExpression) error
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
	TableSelect(*TableSelect) error
	Time(*time.Time) error
	Unit(*units.Unit) error
	UsingDefinitions(*UsingDefinitions) error
	UsingDefinition(*UsingDefinition) error
	WindRose(*WindRose) error
}

type visitor struct {
	common
}

type common struct {
	aliasedExpression  func(Visitor, *AliasedExpression) error
	expression         func(Visitor, *Expression) error
	expressionModifier func(Visitor, *ExpressionModifier) error
	duration           func(Visitor, *time.Duration) error
	function           func(Visitor, *Function) error
	histogram          func(Visitor, *Histogram) error
	metric             func(Visitor, *Metric) error
	query              func(Visitor, *Query) error
	queryRange         func(Visitor, *QueryRange) error
	_select            func(Visitor, *Select) error
	selectExpression   func(Visitor, *SelectExpression) error
	tableSelect        func(Visitor, *TableSelect) error
	time               func(Visitor, *time.Time) error
	unit               func(Visitor, *units.Unit) error
	usingDefinition    func(Visitor, *UsingDefinition) error
	usingDefinitions   func(Visitor, *UsingDefinitions) error
	windRose           func(Visitor, *WindRose) error
}

func (v *visitor) Duration(b *time.Duration) error {
	if b != nil && v.duration != nil {
		return v.duration(v, b)
	}
	return nil
}

func (v *visitor) Time(b *time.Time) error {
	var err error
	if b != nil {
		if v.time != nil {
			err = v.time(v, b)
		}
		if util.IsVisitorStop(err) {
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

func (v *visitor) Unit(b *units.Unit) error {
	if b != nil && v.unit != nil {
		return v.unit(v, b)
	}
	return nil
}
