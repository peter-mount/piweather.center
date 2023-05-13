package value

import "math"

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
// the bounds of the Unit
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
	nv := Value{v: t, u: to}
	return nv, nv.BoundsError()
}

// AsGuard is the same as the As function except that if an error would be returned then
// a panic is issued. This function is normally used within tests
func (v Value) AsGuard(to Unit) Value {
	n, err := v.As(to)
	if err != nil {
		panic(err)
	}
	return n
}

func (v Value) Equals(b Value) (bool, error) { return v.Compare(v, equal) }

func (v Value) NotEqual(b Value) (bool, error) { return v.Compare(v, notEqual) }

func (v Value) LessThan(b Value) (bool, error) { return v.Compare(v, lessThan) }

func (v Value) LessThanEqual(b Value) (bool, error) { return v.Compare(v, lessThanEqual) }

func (v Value) GreaterThan(b Value) (bool, error) { return v.Compare(v, greaterThan) }

func (v Value) GreaterThanEqual(b Value) (bool, error) { return v.Compare(v, greaterThanEqual) }

func (v Value) Compare(b Value, f func(a, b float64) bool) (bool, error) {
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

func equal(a, b float64) bool { return math.Abs(a-b) <= 1e-9 }

func notEqual(a, b float64) bool { return !equal(a, b) }

// a<b unless they are within 1e-9
func lessThan(a, b float64) bool { return notEqual(a, b) && a < b }

func lessThanEqual(a, b float64) bool { return equal(a, b) || a < b }

func greaterThanEqual(a, b float64) bool { return equal(a, b) || a > b }

// a>b unless they are within 1e-9
func greaterThan(a, b float64) bool { return notEqual(a, b) && a > b }

func (v Value) Add(b Value) (Value, error) { return v.calculate(b, add) }
func add(a float64, b float64) float64     { return a + b }

func (v Value) Subtract(b Value) (Value, error) { return v.calculate(b, subtract) }
func subtract(a float64, b float64) float64     { return a - b }

func (v Value) Multiply(b Value) (Value, error) { return v.calculate(b, multiply) }
func multiply(a float64, b float64) float64     { return a * b }

func (v Value) Divide(b Value) (Value, error) { return v.calculate(b, divide) }
func divide(a float64, b float64) float64     { return a / b }

func (v Value) calculate(b Value, f func(float64, float64) float64) (Value, error) {
	c, err := b.As(v.Unit())
	if err != nil {
		return Value{}, err
	}

	return v.Unit().Value(f(v.Float(), c.Float())), nil
}
