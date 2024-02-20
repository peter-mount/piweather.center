package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	util2 "github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/ql"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

func scriptInit(q *ql.Query, err error) (*ql.Query, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = lang2.NewBuilder().
			Query(queryInit).
			QueryRange(queryRangeInit).
			UsingDefinition(parserState.usingDefinitionInit).
			Select(selectInit).
			AliasedExpression(aliasedExpressionInit).
			Expression(parserState.expressionInit).
			ExpressionModifier(parserState.expressionModifierInit).
			Function(functionInit).
			Metric(metricInit).
			Time(timeInit).
			Duration(durationInit).
			WindRose(windRoseInit).
			Build().
			Query(q)
	}
	return q, err
}

func expressionInit(q *ql.Expression, err error) (*ql.Expression, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = lang2.NewBuilder().
			ExpressionModifier(parserState.expressionModifierInit).
			Function(functionInit).
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

func queryInit(_ ql.QueryVisitor, s *ql.Query) error {
	return assertLimit(s.Pos, s.Limit)
}

func queryRangeInit(v ql.QueryVisitor, q *ql.QueryRange) error {
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

func selectInit(_ ql.QueryVisitor, s *ql.Select) error {
	return assertLimit(s.Pos, s.Limit)
}

func aliasedExpressionInit(_ ql.QueryVisitor, s *ql.AliasedExpression) error {
	if s.Unit != "" {
		u, exists := value.GetUnit(s.Unit)
		if !exists {
			return errors.Errorf(s.Pos, "unsupported unit %q", s.Unit)
		}
		s.SetUnit(u)
	}
	return nil
}

func functionInit(_ ql.QueryVisitor, b *ql.Function) error {
	if functions.HasFunction(b.Name) {
		return nil
	}
	return errors.Errorf(b.Pos, "unknown function %q", b.Name)
}

func metricInit(_ ql.QueryVisitor, b *ql.Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

func timeInit(v ql.QueryVisitor, t *time2.Time) error {
	if t == nil {
		return nil
	}

	if err := t.SetTime(unit.ParseTime(t.Def), 0, v); err != nil {
		return err
	}

	return util2.VisitorStop
}

func durationInit(_ ql.QueryVisitor, d *time2.Duration) error {
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

func (p *parserState) usingDefinitionInit(v ql.QueryVisitor, u *ql.UsingDefinition) error {
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

func (p *parserState) expressionInit(_ ql.QueryVisitor, s *ql.Expression) error {
	if s.Using != "" && !p.usingNames.Contains(s.Using) {
		return errors.Errorf(s.Pos, "%q undefined", s.Using)
	}
	return nil
}

func (p *parserState) expressionModifierInit(v ql.QueryVisitor, s *ql.ExpressionModifier) error {
	err := v.QueryRange(s.Range)
	if err == nil {
		err = v.Duration(s.Offset)
	}
	return err
}

func windRoseInit(_ ql.QueryVisitor, s *ql.WindRose) error {
	// Ensure we have a default option of Rose if none set
	if len(s.Options) == 0 {
		s.Options = append(s.Options, ql.WindRoseOption{Rose: true})
	}
	return nil
}
