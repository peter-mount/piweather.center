package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	util2 "github.com/peter-mount/piweather.center/config/util"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"time"
)

var (
	scriptInitVisitor = NewBuilder[*parserState]().
				Query(initQuery).
				QueryRange(initQueryRange).
				UsingDefinition(initUsingDefinition).
				Select(initSelect).
				Expression(initExpression).
				ExpressionModifier(initExpressionModifier).
				Metric(initMetric).
				Time(timeInit).
				TimeZone(timeZoneInit).
				Duration(durationInit).
				Unit(unitInit).
				WindRose(initWindRose).
				TableSelect(initTableSelect).
				Build()

	expressionInitVisitor = NewBuilder[*parserState]().
				ExpressionModifier(initExpressionModifier).
				Metric(initMetric).
				Time(timeInit).
				Duration(durationInit).
				Build()
)

func scriptInit(q *Query, err error) (*Query, error) {
	if err == nil {
		err = scriptInitVisitor.Clone().
			Set(newParserState()).
			Query(q)
	}
	return q, err
}

func expressionInit(q *Expression, err error) (*Expression, error) {
	if err == nil {
		err = expressionInitVisitor.Clone().
			Set(newParserState()).
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

func unitInit(_ Visitor[*parserState], s *units.Unit) error {
	return s.Init()
}

func timeInit(v Visitor[*parserState], t *time2.Time) error {
	if t == nil {
		return nil
	}

	tz := v.Get().timeZone

	if err := t.SetTime(unit.ParseTimeIn(t.Def, tz), 0, v); err != nil {
		return err
	}

	return util2.VisitorStop
}

func timeZoneInit(v Visitor[*parserState], d *time2.TimeZone) error {
	if err := d.Init(); err != nil {
		return err
	}
	v.Get().timeZone = d.Location()
	return nil
}

func durationInit(_ Visitor[*parserState], d *time2.Duration) error {
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
	timeZone   *time.Location
}

func newParserState() *parserState {
	return &parserState{usingNames: util.NewStringSet(), timeZone: time.UTC}
}
