package value

import (
	"fmt"
	"math"
)

// Value is a float64 with an associated Unit.
// Value's can be transformed between different Unit's if those unit's support
// a specific conversion.
type Value struct {
	v float64
	u Unit
}

// Float returns the float64 in the Value's Unit
func (v Value) Float() float64 {
	return v.v
}

// Unit returns the Unit for this Value
func (v Value) Unit() Unit {
	return v.u
}

// String returns this Value as a string with the appropriate Unit attached
func (v Value) String() string {
	return v.u.String(v.v)
}

// IsValid returns true if the Value is valid, specifically if it's within
// the bounds of the Unit.
//
// If the value is NaN or either Infinity then this returns false.
func (v Value) IsValid() bool {
	return v.u.Valid(v.v)
}

// BoundsError returns an error if IsValid() returns false, nil otherwise.
func (v Value) BoundsError() error {
	return v.u.BoundsError(v.v)
}

// As converts this Value to another Unit.
// This will return an error if this value is invalid, or if there is no
// available transform from this Value's Unit to the requested one.
func (v Value) As(to Unit) (Value, error) {
	// Source value is invalid
	if !v.IsValid() {
		return Value{}, v.BoundsError()
	}

	t, err := Transform(v.v, v.u, to)
	if err != nil {
		return Value{}, err
	}

	// Create new value but return an error if it's now invalid
	nv := to.Value(t)
	return nv, nv.BoundsError()
}

// AsGuard is the same as the As function except that if an error would be returned then
// a panic is issued. This function is normally used within tests
func (v Value) AsGuard(to Unit) Value {
	n, err := v.As(to)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return n
}

// Equals returns true if both values are equal. This accounts for differing units.
// Returns an error if either value is invalid or if it's not possible to transform
// b to the same unit as v.
//
// Equality here is if the two values are within 1e-9 of each other to account for
// rounding errors within float64.
func (v Value) Equals(b Value) (bool, error) { return v.Compare(b, equal) }

func equal(a, b float64) bool { return math.Abs(a-b) <= 1e-9 }

// NotEqual returns true if both values are equal. It's the same as !Equal() and
// follows the same rules.
func (v Value) NotEqual(b Value) (bool, error) { return v.Compare(b, notEqual) }

func notEqual(a, b float64) bool { return !equal(a, b) }

// LessThan returns true if v < b, accounting for different units.
// It will return false if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) LessThan(b Value) (bool, error) { return v.Compare(b, lessThan) }

// a<b unless they are within 1e-9
func lessThan(a, b float64) bool { return a < b && notEqual(a, b) }

// LessThanEqual returns true if v <= b, accounting for different units.
// It will return true if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) LessThanEqual(b Value) (bool, error) { return v.Compare(b, lessThanEqual) }

func lessThanEqual(a, b float64) bool { return a < b || equal(a, b) }

// GreaterThan returns true if v > b, accounting for different units.
// It will return false if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) GreaterThan(b Value) (bool, error) { return v.Compare(b, greaterThan) }

// a>b unless they are within 1e-9
func greaterThan(a, b float64) bool { return a > b && notEqual(a, b) }

// GreaterThanEqual returns true if v >= b, accounting for different units.
// It will return true if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) GreaterThanEqual(b Value) (bool, error) { return v.Compare(b, greaterThanEqual) }

func greaterThanEqual(a, b float64) bool { return a > b || equal(a, b) }

// IsZero returns true if the value is zero.
// Specifically if |v|<1e-9 to account for rounding errors in float64.
func (v Value) IsZero() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return equal(v.Float(), 0), nil
}

// IsOne returns true if the value is 1.
// Specifically if |v-1|<1e-9 to account for rounding errors in float64.
func (v Value) IsOne() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return equal(v.Float(), 1), nil
}

// IsPositive returns true if the value is positive.
// 0 is neither positive nor negative/
// Specifically if v > 1e-9 to account for rounding errors in float64.
func (v Value) IsPositive() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return greaterThan(v.Float(), 0), nil
}

// IsNegative returns true if the value is negative.
// 0 is neither positive nor negative/
// Specifically if v < -1e-9 to account for rounding errors in float64.
func (v Value) IsNegative() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return lessThan(v.Float(), 1), nil
}

type Comparator func(a, b float64) bool

// Compare will return the result of a Comparator when it's passed values from v and b.
// It will transform b to the same unit as v before passing it to the Comparator.
// An error will be returned if either value is invalid or it's not possible to transform
// b to v.
func (v Value) Compare(b Value, f Comparator) (bool, error) {
	// One of the values are invalid
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	if !b.IsValid() {
		return false, b.BoundsError()
	}

	// Same unit just do the comparison
	if v.Unit() == b.Unit() {
		return f(v.v, b.v), nil
	}

	c, err := b.As(v.u)
	if err != nil {
		return false, err
	}
	return f(v.v, c.v), nil
}

// CompareGuard is the same as Compare except that if the comparison fails it will panic.
// This is normally used within tests.
func (v Value) CompareGuard(b Value, f Comparator) bool {
	r, err := v.Compare(b, f)
	if err != nil {
		panic(err)
	}
	return r
}

// Within returns true if b >= v <= c.
// It returns an error if any of the three values are invalid or if it's not possible to
// transform b and c to the same unit as v.
func (v Value) Within(b, c Value) (bool, error) {
	b1, err := b.As(v.Unit())
	if err != nil {
		return false, err
	}

	c1, err := c.As(v.Unit())
	if err != nil {
		return false, err
	}

	return within(v.Float(), b1.Float(), c1.Float()), nil
}

func within(a, b, c float64) bool {
	// Ensure b < c
	if greaterThan(b, c) {
		b, c = c, b
	}

	return greaterThanEqual(a, b) && lessThanEqual(a, c)
}

// Add returns the sum of two values. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Add(b Value) (Value, error) { return v.calculate(b, add) }
func add(a, b float64) float64             { return a + b }

// Subtract subtracts b from this value. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Subtract(b Value) (Value, error) { return v.calculate(b, subtract) }
func subtract(a, b float64) float64             { return a - b }

// Multiply returns the product of two values. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Multiply(b Value) (Value, error) { return v.calculate(b, multiply) }
func multiply(a, b float64) float64             { return a * b }

// Divide divides this value with b. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
// If the transformed value of b is zero then a normal divide
func (v Value) Divide(b Value) (Value, error) { return v.calculate(b, divide) }
func divide(a, b float64) float64             { return a / b }

func (v Value) calculate(b Value, f func(float64, float64) float64) (Value, error) {
	c, err := b.As(v.Unit())
	if err != nil {
		return Value{}, err
	}

	r := v.Unit().Value(f(v.Float(), c.Float()))
	return r, r.BoundsError()
}
