package lang

import (
	"errors"
)

type Visitor interface {
	Query(*Query) error
	Select(*Select) error
	SelectExpression(*SelectExpression) error
	AliasedExpression(*AliasedExpression) error
	Expression(*Expression) error
	ExpressionModifier(*ExpressionModifier) error
	Function(*Function) error
	Metric(*Metric) error
	QueryRange(*QueryRange) error
	Time(*Time) error
	Duration(*Duration) error
	UsingDefinitions(definitions *UsingDefinitions) error
	UsingDefinition(definitions *UsingDefinition) error
}

// VisitorStop is an error which causes the current step in a Visitor to stop processing.
// It's used to enable a Visitor to handle all processing of a node within itself rather
// than the Visitor proceeding to any child nodes of that node.
var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

// VisitorExit is an error which will terminate the Visitor.
// This is the same as any error occurring within a Visitor except that the final error
// returned from specific handlers will become nil.
var VisitorExit = errors.New("visitor exit")

func IsVisitorExit(err error) bool {
	return err != nil && errors.Is(err, VisitorExit)
}

type Visitable interface {
	Accept(v Visitor) error
}

type visitor struct {
	common
}

type common struct {
	query              func(Visitor, *Query) error
	_select            func(Visitor, *Select) error
	selectExpression   func(Visitor, *SelectExpression) error
	aliasedExpression  func(Visitor, *AliasedExpression) error
	expression         func(Visitor, *Expression) error
	expressionModifier func(Visitor, *ExpressionModifier) error
	function           func(Visitor, *Function) error
	metric             func(Visitor, *Metric) error
	queryRange         func(Visitor, *QueryRange) error
	time               func(Visitor, *Time) error
	duration           func(Visitor, *Duration) error
	usingDefinitions   func(Visitor, *UsingDefinitions) error
	usingDefinition    func(Visitor, *UsingDefinition) error
}

func (v *visitor) Query(b *Query) error {
	var err error
	if b != nil {
		// Process QueryRange first
		err = v.QueryRange(b.QueryRange)

		if err == nil && v.query != nil {
			err = v.query(v, b)
		}
		if IsVisitorStop(err) || IsVisitorExit(err) {
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
	}
	return err
}

func (v *visitor) Select(b *Select) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v._select(v, b)
			if IsVisitorStop(err) || IsVisitorExit(err) {
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
			if IsVisitorStop(err) {
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
		if IsVisitorStop(err) {
			return nil
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

		if IsVisitorStop(err) {
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
		if IsVisitorStop(err) {
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
		if IsVisitorStop(err) {
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
			if IsVisitorStop(err) {
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

func (v *visitor) Time(b *Time) error {
	var err error
	if b != nil {
		if v.time != nil {
			err = v.time(v, b)
		}
		if IsVisitorStop(err) {
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

func (v *visitor) Duration(b *Duration) error {
	if b != nil && v.duration != nil {
		return v.duration(v, b)
	}
	return nil
}

func (v *visitor) UsingDefinitions(b *UsingDefinitions) error {
	var err error

	if b != nil {
		if v.usingDefinitions != nil {
			err = v.usingDefinitions(v, b)
		}
		if IsVisitorStop(err) {
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
