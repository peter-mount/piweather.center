package ql

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	util2 "github.com/peter-mount/piweather.center/config/util"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"strings"
	"time"
)

func scriptInit(q *Query, err error) (*Query, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = NewBuilder().
			Query(queryInit).
			QueryRange(queryRangeInit).
			UsingDefinition(parserState.usingDefinitionInit).
			Select(selectInit).
			Expression(parserState.expressionInit).
			ExpressionModifier(parserState.expressionModifierInit).
			Metric(metricInit).
			Time(timeInit).
			Duration(durationInit).
			Unit(unitInit).
			WindRose(windRoseInit).
			TableSelect(tableSelect).
			Build().
			Query(q)
	}
	return q, err
}

func expressionInit(q *Expression, err error) (*Expression, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = NewBuilder().
			ExpressionModifier(parserState.expressionModifierInit).
			Metric(metricInit).
			Time(timeInit).
			Duration(durationInit).
			Build().
			Expression(q)
	}
	return q, err
}

func assertLimit(p lexer.Position, l int) error {
	if l < 0 {
		return errors.Errorf(p, "invalid LIMIT %d", l)
	}
	return nil
}

func queryInit(_ Visitor, s *Query) error {
	return assertLimit(s.Pos, s.Limit)
}

func queryRangeInit(v Visitor, q *QueryRange) error {
	// If no Every statement then set it to 1 minute
	if q.Every == nil {
		q.Every = &time2.Duration{Pos: q.Pos, Def: "1m"}
	}

	if err := v.Duration(q.Every); err != nil {
		return err
	}

	// Negative duration for Every is invalid
	if q.Every.Duration(0) < time.Second {
		return fmt.Errorf("invalid step size %v", q.Every.Duration(0))
	}

	if err := v.Time(q.At); err != nil {
		return err
	}

	if err := v.Time(q.From); err != nil {
		return err
	}
	if err := v.Duration(q.For); err != nil {
		return err
	}

	if err := v.Time(q.Start); err != nil {
		return err
	}
	if err := v.Time(q.End); err != nil {
		return err
	}

	return util2.VisitorStop
}

func selectInit(_ Visitor, s *Select) error {
	return assertLimit(s.Pos, s.Limit)
}

func tableSelect(v Visitor, t *TableSelect) error {
	var err error
	if t.Unit != nil {
		err = v.Unit(t.Unit)
	}
	return err
}

func unitInit(_ Visitor, s *units.Unit) error {
	return s.Init()
}

func metricInit(_ Visitor, b *Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

func timeInit(v Visitor, t *time2.Time) error {
	if t == nil {
		return nil
	}

	if err := t.SetTime(unit.ParseTime(t.Def), 0, v); err != nil {
		return err
	}

	return util2.VisitorStop
}

func durationInit(_ Visitor, d *time2.Duration) error {
	if d.Def != "" && !d.IsEvery() {
		v, err := time.ParseDuration(d.Def)
		if err != nil {
			return errors.Error(d.Pos, err)
		}
		d.Set(v)
	}

	return nil
}

type parserState struct {
	usingNames util.StringSet
}

func (p *parserState) usingDefinitionInit(v Visitor, u *UsingDefinition) error {
	if !p.usingNames.Add(u.Name) {
		return errors.Errorf(u.Pos, "alias %q already defined", u.Name)
	}
	for _, e := range u.Modifier {
		if err := v.ExpressionModifier(e); err != nil {
			return err
		}
	}
	return nil
}

func (p *parserState) expressionInit(_ Visitor, s *Expression) error {
	if s.Using != "" && !p.usingNames.Contains(s.Using) {
		return errors.Errorf(s.Pos, "%q undefined", s.Using)
	}
	return nil
}

func (p *parserState) expressionModifierInit(v Visitor, s *ExpressionModifier) error {
	err := v.QueryRange(s.Range)
	if err == nil {
		err = v.Duration(s.Offset)
	}
	return err
}

func windRoseInit(_ Visitor, s *WindRose) error {
	// Ensure we have a default option of Rose if none set
	if len(s.Options) == 0 {
		s.Options = append(s.Options, WindRoseOption{Rose: true})
	}
	return nil
}
