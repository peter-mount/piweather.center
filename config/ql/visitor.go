package ql

import (
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type QueryVisitor interface {
	time.TimeVisitor
	units.UnitsVisitor
	Query(*Query) error
	Select(*Select) error
	SelectExpression(*SelectExpression) error
	AliasedExpression(*AliasedExpression) error
	Expression(*Expression) error
	ExpressionModifier(*ExpressionModifier) error
	Function(*Function) error
	Metric(*Metric) error
	QueryRange(*QueryRange) error
	UsingDefinitions(*UsingDefinitions) error
	UsingDefinition(*UsingDefinition) error
	Histogram(*Histogram) error
	WindRose(*WindRose) error
	TableSelect(*TableSelect) error
}

type visitor struct {
	common
}

type common struct {
	query              func(QueryVisitor, *Query) error
	_select            func(QueryVisitor, *Select) error
	selectExpression   func(QueryVisitor, *SelectExpression) error
	aliasedExpression  func(QueryVisitor, *AliasedExpression) error
	expression         func(QueryVisitor, *Expression) error
	expressionModifier func(QueryVisitor, *ExpressionModifier) error
	function           func(QueryVisitor, *Function) error
	metric             func(QueryVisitor, *Metric) error
	queryRange         func(QueryVisitor, *QueryRange) error
	time               func(QueryVisitor, *time.Time) error
	duration           func(QueryVisitor, *time.Duration) error
	usingDefinitions   func(QueryVisitor, *UsingDefinitions) error
	usingDefinition    func(QueryVisitor, *UsingDefinition) error
	histogram          func(QueryVisitor, *Histogram) error
	unit               func(QueryVisitor, *units.Unit) error
	windRose           func(QueryVisitor, *WindRose) error
	tableSelect        func(QueryVisitor, *TableSelect) error
}

func (v *visitor) Query(b *Query) error {
	var err error
	if b != nil {
		// Process QueryRange first
		err = v.QueryRange(b.QueryRange)

		if err == nil && v.query != nil {
			err = v.query(v, b)
		}
		if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
			return nil
		}

		if err == nil {
			err = v.UsingDefinitions(b.Using)
		}

		if err == nil {
			for _, sel := range b.Select {
				if err == nil {
					err = v.Select(sel)
				}
			}
		}

		if err == nil {
			for _, sel := range b.WindRose {
				if err == nil {
					err = v.WindRose(sel)
				}
			}
		}

		if err == nil {
			err = v.TableSelect(b.TableSelect)
		}
	}
	return err
}

func (v *visitor) Select(b *Select) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v._select(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}

		if err == nil {
			err = v.SelectExpression(b.Expression)
		}
	}
	return err
}

func (v *visitor) SelectExpression(b *SelectExpression) error {
	var err error
	if b != nil {
		if v.selectExpression != nil {
			err = v.selectExpression(v, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}
		if err == nil {
			for _, e := range b.Expressions {
				err = v.AliasedExpression(e)
				if err != nil {
					break
				}
			}
		}
	}
	return err
}

func (v *visitor) AliasedExpression(b *AliasedExpression) error {
	var err error
	if b != nil {
		if v.aliasedExpression != nil {
			err = v.aliasedExpression(v, b)
		}
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = v.Unit(b.Unit)
		}
		if err == nil {
			err = v.Expression(b.Expression)
		}
	}
	return err
}

func (v *visitor) Expression(b *Expression) error {
	var err error
	if b != nil {
		if v.expression != nil {
			err = v.expression(v, b)
		}

		if util.IsVisitorStop(err) {
			return nil
		}

		for _, e := range b.Modifier {
			if err == nil {
				err = v.ExpressionModifier(e)
			}
		}

		if err == nil {
			err = v.Function(b.Function)
		}

		if err == nil {
			err = v.Metric(b.Metric)
		}
	}

	return err
}

func (v *visitor) ExpressionModifier(b *ExpressionModifier) error {
	var err error
	if b != nil {
		if v.expressionModifier != nil {
			err = v.expressionModifier(v, b)
		}
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			err = v.Duration(b.Offset)
		}
		if err == nil {
			err = v.QueryRange(b.Range)
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
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			for _, ex := range b.Expressions {
				err = v.Expression(ex)
				if err != nil {
					break
				}
			}
		}
	}
	return err
}

func (v *visitor) Metric(b *Metric) error {
	if b != nil && v.metric != nil {
		return v.metric(v, b)
	}
	return nil
}

func (v *visitor) QueryRange(b *QueryRange) error {
	var err error
	if b != nil {
		if v.queryRange != nil {
			err = v.queryRange(v, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		// AT x
		if err == nil {
			err = v.Time(b.At)
		}

		// FROM x FOR x
		if err == nil {
			err = v.Time(b.From)
		}
		if err == nil {
			err = v.Duration(b.For)
		}

		// BETWEEN x AND x
		if err == nil {
			err = v.Time(b.Start)
		}
		if err == nil {
			err = v.Time(b.End)
		}

		if err == nil {
			err = v.Duration(b.Every)
		}
	}
	return err
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

func (v *visitor) Duration(b *time.Duration) error {
	if b != nil && v.duration != nil {
		return v.duration(v, b)
	}
	return nil
}

func (v *visitor) Unit(b *units.Unit) error {
	if b != nil && v.unit != nil {
		return v.unit(v, b)
	}
	return nil
}

func (v *visitor) UsingDefinitions(b *UsingDefinitions) error {
	var err error

	if b != nil {
		if v.usingDefinitions != nil {
			err = v.usingDefinitions(v, b)
		}
		if util.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			for _, e := range b.Defs {
				err = v.UsingDefinition(e)
			}
		}
	}

	return err
}

func (v *visitor) UsingDefinition(b *UsingDefinition) error {
	if b != nil && v.usingDefinition != nil {
		return v.usingDefinition(v, b)
	}
	return nil
}

func (v *visitor) Histogram(b *Histogram) error {
	var err error
	if b != nil {
		if v.histogram != nil {
			err = v.histogram(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}
		if err == nil {
			err = v.AliasedExpression(b.Expression)
		}
	}
	return err
}

func (v *visitor) WindRose(b *WindRose) error {
	var err error
	if b != nil {
		if v.windRose != nil {
			err = v.windRose(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}
		if err == nil {
			err = v.Expression(b.Degrees)
		}
		if err == nil {
			err = v.Expression(b.Speed)
		}
	}
	return err
}

func (v *visitor) TableSelect(b *TableSelect) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v.tableSelect(v, b)
			if util.IsVisitorStop(err) || util.IsVisitorExit(err) {
				return nil
			}
		}

		if err == nil {
			err = v.Expression(b.Time)
		}
		if err == nil {
			err = v.Expression(b.Metric)
		}
		if err == nil {
			err = v.Unit(b.Unit)
		}
	}
	return err
}
