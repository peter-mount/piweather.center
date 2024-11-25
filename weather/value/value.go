package value

import (
	"errors"
	"math"
)

// Value is a float64 with an associated Unit.
// Value's can be transformed between different Unit's if those unit's support
// a specific conversion.
type Value struct {
	v float64
	u *Unit
}

// Float returns the float64 in the Value's Unit
func (v Value) Float() float64 {
	return v.v
}

// Unit returns the Unit for this Value
func (v Value) Unit() *Unit {
	return v.u
}

// String returns this Value as a string with the appropriate Unit attached
func (v Value) String() string {
	if v.u == nil {
		return invalidValue.Error()
	}
	return v.u.String(v.v)
}

func (v Value) PlainString() string {
	if v.u == nil {
		return invalidValue.Error()
	}
	return v.u.PlainString(v.v)
}

// IsValid returns true if the Value is valid, specifically if it's Within
// the bounds of the Unit.
//
// If the value is NaN or either Infinity then this returns false.
func (v Value) IsValid() bool {
	return v.u != nil && v.u.Valid(v.v)
}

var (
	invalidValue = errors.New("invalid Value")
)

// BoundsError returns an error if IsValid() returns false, nil otherwise.
func (v Value) BoundsError() error {
	if v.u == nil {
		return invalidValue
	}
	return v.u.BoundsError(v.v)
}

// As converts this Value to another Unit.
// This will return an error if this value is invalid,
// if there is no available transform from this Value's Unit to the requested one,
// or if the result is invalid.
func (v Value) As(to *Unit) (Value, error) {
	// Source value is invalid
	if !v.IsValid() {
		return Value{}, v.BoundsError()
	}

	// If to is the same unit as Value return the value unchanged
	if v.u.Equals(to) {
		return v, nil
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
// a panic is issued. This function is normally used Within tests
func (v Value) AsGuard(to *Unit) Value {
	n, err := v.As(to)
	if err != nil {
		panic(err)
	}
	return n
}

// Value returns a new Value with the same unit as this one.
// This is the same as v.Unit().Value(f).
//
// If the value has no Unit then this returns an invalid Unit.
//
// If the value of the new unit is outside its bounds,
// it will still be returned however it will indicate it's not valid.
func (v Value) Value(f float64) Value {
	if v.Unit() == nil {
		return Value{}
	}
	return v.Unit().Value(f)
}

// Equals returns true if both values are Equal. This accounts for differing units.
// Returns an error if either value is invalid or if it's not possible to transform
// b to the same unit as v.
//
// Equality here is if the two values are Within 1e-9 of each other to account for
// rounding errors Within float64.
func (v Value) Equals(b Value) (bool, error) { return v.Compare(b, Equal) }

// Equal returns true if both values are Equal.
//
// Equality here is if the two values are Within 1e-9 of each other to account for
// rounding errors Within float64.
func Equal(a, b float64) bool { return math.Abs(a-b) <= 1e-9 }

// NotEqual returns true if both values are Equal. It's the same as !Equal() and
// follows the same rules.
func (v Value) NotEqual(b Value) (bool, error) { return v.Compare(b, NotEqual) }

// NotEqual returns true if both values are Equal. It's the same as !Equal() and
// follows the same rules.
func NotEqual(a, b float64) bool { return !Equal(a, b) }

// LessThan returns true if v < b, accounting for different units.
// It will return false if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) LessThan(b Value) (bool, error) { return v.Compare(b, LessThan) }

// LessThan returns true if a < b, accounting for different units.
// It will return false if |a-b|<=1e-9 to account for rounding errors in float64.
func LessThan(a, b float64) bool { return a < b && NotEqual(a, b) }

// LessThanEqual returns true if v <= b, accounting for different units.
// It will return true if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) LessThanEqual(b Value) (bool, error) { return v.Compare(b, LessThanEqual) }

// LessThanEqual returns true if a <= b, accounting for different units.
// It will return true if |a-b|<=1e-9 to account for rounding errors in float64.
func LessThanEqual(a, b float64) bool { return a < b || Equal(a, b) }

// GreaterThan returns true if v > b, accounting for different units.
// It will return false if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) GreaterThan(b Value) (bool, error) { return v.Compare(b, GreaterThan) }

// GreaterThan returns true if a > b, accounting for different units.
// It will return false if |a-b|<=1e-9 to account for rounding errors in float64.
func GreaterThan(a, b float64) bool { return a > b && NotEqual(a, b) }

// GreaterThanEqual returns true if v >= b, accounting for different units.
// It will return true if |v-b|<=1e-9 to account for rounding errors in float64.
func (v Value) GreaterThanEqual(b Value) (bool, error) { return v.Compare(b, GreaterThanEqual) }

// GreaterThanEqual returns true if a >= b, accounting for different units.
// It will return true if |a-b|<=1e-9 to account for rounding errors in float64.
func GreaterThanEqual(a, b float64) bool { return a > b || Equal(a, b) }

// IsZero returns true if the value is zero.
// Specifically if |v|<1e-9 to account for rounding errors in float64.
func (v Value) IsZero() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return IsZero(v.Float()), nil
}

// IsZero returns true if the value is zero.
// Specifically if |v|<1e-9 to account for rounding errors in float64.
func IsZero(f float64) bool {
	return Equal(f, 0)
}

// IsOne returns true if the value is 1.
// Specifically if |v-1|<1e-9 to account for rounding errors in float64.
func (v Value) IsOne() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return Equal(v.Float(), 1), nil
}

// IsOne returns true if the value is 1.
// Specifically if |v-1|<1e-9 to account for rounding errors in float64.
func IsOne(f float64) bool {
	return Equal(f, 1)
}

// IsPositive returns true if the value is positive.
// 0 is neither positive nor negative/
// Specifically if v > 1e-9 to account for rounding errors in float64.
func (v Value) IsPositive() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return IsPositive(v.Float()), nil
}

// IsPositive returns true if the value is positive.
// 0 is neither positive nor negative/
// Specifically if v > 1e-9 to account for rounding errors in float64.
func IsPositive(f float64) bool {
	return GreaterThan(f, 0)
}

// IsNegative returns true if the value is negative.
// 0 is neither positive nor negative/
// Specifically if v < -1e-9 to account for rounding errors in float64.
func (v Value) IsNegative() (bool, error) {
	if !v.IsValid() {
		return false, v.BoundsError()
	}
	return IsNegative(v.Float()), nil
}

// IsNegative returns true if the value is negative.
// 0 is neither positive nor negative/
// Specifically if v < -1e-9 to account for rounding errors in float64.
func IsNegative(f float64) bool {
	return LessThan(f, 0)
}

// Comparator is a function that can be passed to Compare two Values.
// It is guaranteed that the values supplied to it will be of the same Unit.
type Comparator func(a, b float64) bool

// TrueComparator is a Comparator that always returns true
func TrueComparator(_, _ float64) bool { return true }

// FalseComparator is a Comparator that always returns false
func FalseComparator(_, _ float64) bool { return false }

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
// This is normally used Within tests.
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

	return Within(v.Float(), b1.Float(), c1.Float()), nil
}

// Within returns true if b >= a <= c
func Within(a, b, c float64) bool {
	// Ensure b < c
	if GreaterThan(b, c) {
		b, c = c, b
	}

	return GreaterThanEqual(a, b) && LessThanEqual(a, c)
}

// Without returns true if a <= b || a >= c
func Without(a, b, c float64) bool {
	// Ensure b < c
	if GreaterThan(b, c) {
		b, c = c, b
	}

	return LessThanEqual(a, b) || GreaterThanEqual(a, c)
}

// Add returns the sum of two values. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Add(b Value) (Value, error) { return v.Calculate(b, Add) }
func Add(a, b float64) float64             { return a + b }

// Subtract subtracts b from this value. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Subtract(b Value) (Value, error) { return v.Calculate(b, Subtract) }
func Subtract(a, b float64) float64             { return a - b }

// Multiply returns the product of two values. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
func (v Value) Multiply(b Value) (Value, error) { return v.Calculate(b, Multiply) }
func Multiply(a, b float64) float64             { return a * b }

// Divide divides this value with b. The result is the same unit as v.
// An error is returned if either value is invalid,
// b could not be transformed into the same unit as v
// or if the result is itself invalid.
// If the transformed value of b is zero then a normal divide
func (v Value) Divide(b Value) (Value, error) { return v.Calculate(b, Divide) }
func Divide(a, b float64) float64             { return a / b }

// Calculation is a function that applies some mathematical calculation against
// two values. It is guaranteed that both values passed are in the same Unit.
type Calculation func(float64, float64) float64

// Calculate applies a Calculation against v and b returning a new Value
// or an error if either v or b are invalid,
// if b cannot be transformed into v's unit,
// or if the result is itself invalid.
func (v Value) Calculate(b Value, f Calculation) (Value, error) {
	c, err := b.As(v.Unit())
	if err != nil {
		return Value{}, err
	}

	r := v.Value(f(v.Float(), c.Float()))
	return r, r.BoundsError()
}

// Enforce returns a new Value so that it's within the bounds b...c,
// or an error if either v, b or c are invalid,
// if either b or c cannot be transformed to v's unit,
// or if the result is itself invalid.
func (v Value) Enforce(b, c Value) (Value, error) {
	b1, err := b.As(v.Unit())
	if err != nil {
		return Value{}, err
	}

	c1, err := c.As(v.Unit())
	if err != nil {
		return Value{}, err
	}

	r := v.Value(Enforce(v.Float(), b1.Float(), c1.Float()))
	return r, r.BoundsError()
}

// Enforce a so that it's within the bounds b...c
func Enforce(a, b, c float64) float64 {
	if GreaterThan(b, c) {
		b, c = c, b
	}
	if LessThanEqual(a, b) {
		return b
	}
	if GreaterThanEqual(a, c) {
		return c
	}
	return a
}
