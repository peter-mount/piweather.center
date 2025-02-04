package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"sort"
)

// AggregatorList is a list of required aggregators.
type AggregatorList struct {
	Pos         lexer.Position
	Aggregators []string `parser:"'(' @('avg'|'min'|'max'|'sum')+ ')'"`
	normalized  bool
}

// GetAggregators returns a sorted list of the declared aggregators.
// This ensures that the list is unique
func (a *AggregatorList) GetAggregators() []string {
	// normalise the list only once
	if !a.normalized {
		m := make(map[string]bool)
		for _, p := range a.Aggregators {
			m[p] = true
		}
		var r []string
		for _, p := range a.Aggregators {
			r = append(r, p)
		}
		sort.SliceStable(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		a.Aggregators = r
		a.normalized = true
	}

	return a.Aggregators
}
