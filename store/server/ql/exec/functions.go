package exec

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
)

func (ex *executor) function(v lang.Visitor, f *lang.Function) error {
	switch {
	case f.TimeOf:
		return ex.funcTimeOf(v, f)
	default:
		return fmt.Errorf("unknown function %q", f.Name)
	}
}

func assertExpressions(p lexer.Position, e []*lang.Expression, min, max int) error {
	if min > max {
		min, max = max, min
	}
	l := len(e)
	if l < min || l > max {
		if min == max {
			return participle.Errorf(p, "require %d expressions", min)
		}
		return participle.Errorf(p, "require %d..%d expressions", min, max)
	}
	return nil
}

// funcTimeOf implements TIMEOF which marks the value as requiring the TIME not the Value of a metric
func (ex *executor) funcTimeOf(v lang.Visitor, f *lang.Function) error {
	if err := assertExpressions(f.Pos, f.Expressions, 1, 1); err != nil {
		return err
	}

	if err := v.Expression(f.Expressions[0]); err != nil {
		return err
	}

	r, ok := ex.pop()
	if ok {
		r.IsTime = true
		ex.push(r)
	}
	return lang.VisitorStop
}
