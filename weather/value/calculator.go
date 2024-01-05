package value

import (
	"fmt"
	"sort"
	"strings"
)

var (
	calculators = make(map[string]Calculator)
)

// Calculator performs a calculation to generate a Value.
// It takes a time.Time to represent when this calculation represents
type Calculator func(Time, ...Value) (Value, error)

// As returns a Calculator which will attempt to transform the returned
// value from the wrapped Calculator to the required Unit
func (c Calculator) As(to *Unit) Calculator {
	return func(t Time, args ...Value) (Value, error) {
		v, err := c(t, args...)
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

// Calculator2arg utility to convert a function that takes two Value's into a Calculator
func Calculator2arg(f func(_, _ Value) (Value, error)) Calculator {
	return func(_ Time, args ...Value) (Value, error) {
		return f(args[0], args[1])
	}
}

// Calculator3arg utility to convert a function that takes three Value's into a Calculator
func Calculator3arg(f func(_, _, _ Value) (Value, error)) Calculator {
	return func(_ Time, args ...Value) (Value, error) {
		return f(args[0], args[1], args[2])
	}
}

func Basic2ArgCalculator(f func(float64, float64) float64) Calculator {
	return func(_ Time, value ...Value) (Value, error) {
		if len(value) != 2 {
			return Value{}, invalidValue
		}
		u1 := value[0].Unit()
		v2, err := value[1].As(u1)
		if err != nil {
			return Value{}, err
		}
		return u1.Value(f(value[0].Float(), v2.Float())), nil
	}
}
