package exec

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	ql2 "github.com/peter-mount/piweather.center/store/ql"
	"github.com/peter-mount/piweather.center/store/ql/functions"
)

// function executes the provided function.
func (ex *Executor) function(v ql.Visitor[*Executor], d *ql.Function) error {
	f, exists := functions.GetFunction(d.Name)
	if !exists {
		// Should never occur as we check this when building the query plan
		return participle.Errorf(d.Pos, "unknown function %q", d.Name)
	}

	//return af.Run(ex, v, f)

	if err := functions.AssertExpressions(d.Pos, d.Name, d.Expressions, f); err != nil {
		return err
	}

	// Process the arguments, applying any aggregator against each one
	var args []ql2.Value
	for _, e := range d.Expressions {
		err := v.Expression(e)
		if err != nil {
			return err
		}

		val, _ := ex.Pop()

		if f.IsAggregator() {
			val, err = f.RunAggregator(val)
			if err != nil {
				return err
			}
		}

		args = append(args, val)
	}

	// We have a function so call it with the arguments
	if f.Function != nil {

		// Aggregate the arguments into a single value
		if f.AggregateArguments {
			val, err := f.RunAggregator(ql2.Value{Time: ex.Time(), Values: args})
			if err != nil {
				return err
			}
			args = []ql2.Value{val}
		}

		if err := f.Function(ex, d, args); err != nil {
			return err
		}
	}

	// No function but we are an Aggregator then ensure we push a result
	if f.Function == nil && f.IsAggregator() {
		switch len(args) {
		// No args then push null
		case 0:
			ex.Push(ql2.Value{Time: ex.Time()})

		// 1 arg so use it
		case 1:
			ex.Push(args[0])

		// Aggregate the args to get the final result
		default:
			val, err := f.RunAggregator(ql2.Value{Time: ex.Time(), Values: args})
			if err != nil {
				return err
			}
			ex.Push(val)
		}
	}

	return util.VisitorStop
}
