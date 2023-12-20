package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/store/ql/lang"
)

// function executes the provided function.
func (ex *Executor) function(v lang.Visitor, f *lang.Function) error {
	if af, exists := functions.GetFunction(f.Name); exists {
		return af.Run(ex, v, f)
	}
	// Should never occur as we check this when building the query plan
	return participle.Errorf(f.Pos, "unknown function %q", f.Name)
}
