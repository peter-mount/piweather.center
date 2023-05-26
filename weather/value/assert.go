package value

import "fmt"

// Assertion tests a Value and returns an error if something is wrong
type Assertion func(Value) error

// AssertValid is a simple Assertion that tests a Value is valid.
// This is the same as Value.BoundsError()
func AssertValid(v Value) error {
	return v.BoundsError()
}

// AssertEntry always returns nil.
// It's used as a placeholder to AssertCalculator to indicate a Value is required but can be any Value
func AssertEntry(_ Value) error {
	return nil
}

// AssertCalculator wraps a Calculator to enforce the number and type of Value's
// passed to the Calculator.
// If the type of value is not required to be enforced, just that it exists
// then use AssertEntry or AssertValid.
func AssertCalculator(c Calculator, a ...Assertion) Calculator {
	return func(t Time, v ...Value) (Value, error) {
		if len(v) != len(a) {
			return Value{}, fmt.Errorf("Calculator requires %d arguments, got %d", len(a), len(v))
		}
		for i, ae := range a {
			if err := ae(v[i]); err != nil {
				return Value{}, err
			}
		}
		return c(t, v...)
	}
}
