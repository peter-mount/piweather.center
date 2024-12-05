package utils

import (
	"github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/ql/functions"
)

// GetAggregators returns a slice of top level aggregator functions used in a query.
// If no error is returned, the slice will be the same length of expressions in the select.
// If an entry is not null then it will be a pointer to a Function definition which
// IsAggregator() returns true.
func GetAggregators(d *ql.Select) ([]*functions.Function, error) {
	st := &getAggregatorsState{}
	err := getAggregatorsVisitor.Clone().Set(st).Select(d)
	if err != nil {
		return nil, err
	}
	return st.functions, nil
}

type getAggregatorsState struct {
	functions []*functions.Function // Aggregator functions by column
	col       int                   // Current column being visited
}

var (
	getAggregatorsVisitor = ql.NewBuilder[*getAggregatorsState]().
		Select(func(v ql.Visitor[*getAggregatorsState], d *ql.Select) error {
			st := v.Get()
			st.functions = make([]*functions.Function, len(d.Expression.Expressions))

			for i, e := range d.Expression.Expressions {
				st.col = i
				err := v.AliasedExpression(e)
				if err != nil {
					return err
				}
			}

			return util.VisitorStop
		}).
		Function(func(v ql.Visitor[*getAggregatorsState], d *ql.Function) error {
			st := v.Get()

			// Stop looking if we have already got a function
			if st.functions[st.col] != nil {
				return util.VisitorStop
			}

			// Lookup and accept if it's an aggregator
			if f, exists := functions.GetFunction(d.Name); exists && f.IsAggregator() {
				st.functions[st.col] = &f
				return util.VisitorStop
			}

			// Carry on until we get the first aggregator function.
			// e.g. this is to handle last(max("metric","metric")) - hopefully this will not break things
			return nil
		}).
		Build()
)
