package value

import (
	"fmt"
	"sort"
	"strings"
)

var (
	calculators = make(map[string]Calculator)
)

// Calculator performs a calculation to generate a Value
type Calculator func(...Value) (Value, error)

// As returns a Calculator which will attempt to transform the returned
// value from the wrapped Calculator to the required Unit
func (c Calculator) As(to *Unit) Calculator {
	return func(args ...Value) (Value, error) {
		v, err := c(args...)
		if err == nil {
			v, err = v.As(to)
		}
		return v, err
	}
}

// NewCalculator registers a named Calculator
func NewCalculator(id string, calc Calculator) {
	id = strings.ToLower(id)

	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := calculators[id]; exists {
		panic(fmt.Errorf("calculator %q already defined", id))
	}
	calculators[id] = calc
}

// GetCalculator returns a named Calculator
func GetCalculator(id string) (Calculator, error) {
	id = strings.ToLower(id)

	mutex.Lock()
	defer mutex.Unlock()

	calc, exists := calculators[id]
	if exists {
		return calc, nil
	}

	return nil, fmt.Errorf("calculator %q not defined", id)
}

// GetCalculatorIDs return's the ID's of all registered Calculator's
func GetCalculatorIDs() []string {
	var r []string

	mutex.Lock()
	defer mutex.Unlock()

	for k, _ := range calculators {
		r = append(r, k)
	}

	sort.SliceStable(r, func(i, j int) bool { return r[i] < r[j] })

	return r
}
