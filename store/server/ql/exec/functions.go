package exec

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
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

// funcTimeOf implements TIMEOF which marks the value as requiring the TIME not the Value of a metric
func (ex *executor) funcTimeOf(v lang.Visitor, f *lang.Function) error {
	log.Printf("TIMEOF")
	if err := v.Expression(f.Expression); err != nil {
		return err
	}

	r, ok := ex.pop()
	log.Printf("TIMEOF %v %v", ok, r)
	if ok {
		r.IsTime = true
		ex.push(r)
	}
	return lang.VisitorStop
}
