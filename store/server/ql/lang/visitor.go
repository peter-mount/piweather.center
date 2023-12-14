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
	Function(*Function) error
	Metric(*Metric) error
	QueryRange(*QueryRange) error
	Time(*Time) error
	Duration(*Duration) error
}

var VisitorStop = errors.New("visitor stop")

func IsVisitorStop(err error) bool {
	return err != nil && errors.Is(err, VisitorStop)
}

type Visitable interface {
	Accept(v Visitor) error
}

type visitor struct {
	common
}

type common struct {
	query             func(Visitor, *Query) error
	_select           func(Visitor, *Select) error
	selectExpression  func(Visitor, *SelectExpression) error
	aliasedExpression func(Visitor, *AliasedExpression) error
	expression        func(Visitor, *Expression) error
	function          func(Visitor, *Function) error
	metric            func(Visitor, *Metric) error
	queryRange        func(Visitor, *QueryRange) error
	time              func(Visitor, *Time) error
	duration          func(Visitor, *Duration) error
}

func (v *visitor) Query(b *Query) error {
	var err error
	if b != nil {
		// Process QueryRange first
		err = v.QueryRange(b.QueryRange)

		if err == nil {
			err = v.Select(b.Select)
		}
		if err == nil && v.query != nil {
			err = v.query(v, b)
		}
	}
	return err
}

func (v *visitor) Select(b *Select) error {
	var err error
	if b != nil {
		if v._select != nil {
			err = v._select(v, b)
			if IsVisitorStop(err) {
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
		for _, e := range b.Expressions {
			err = v.AliasedExpression(e)
			if err != nil {
				break
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
		if err == nil {
			err = v.Function(b.Function)
		}
		if err == nil {
			err = v.Metric(b.Metric)
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
		if err == nil {
			err = v.Time(b.At)
		}
		if err == nil {
			err = v.Time(b.From)
		}
		if err == nil {
			err = v.Time(b.To)
		}
		if err == nil {
			err = v.Duration(b.Every)
		}
	}
	return err
}

func (v *visitor) Time(b *Time) error {
	if b != nil && v.time != nil {
		return v.time(v, b)
	}
	return nil
}

func (v *visitor) Duration(b *Duration) error {
	if b != nil && v.duration != nil {
		return v.duration(v, b)
	}
	return nil
}
