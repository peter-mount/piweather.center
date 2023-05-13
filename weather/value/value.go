package value

import "math"

type Value struct {
	v float64
	u Unit
}

func (v Value) Float() float64 {
	return v.v
}

func (v Value) Unit() Unit {
	return v.u
}

func (v Value) String() string {
	return v.u.String(v.v)
}

func (v Value) Valid() bool {
	return v.u.Valid(v.v)
}

func (v Value) BoundsError() error {
	return v.u.BoundsError(v.v)
}

func (v Value) Transform(to Unit) (Value, error) {
	// Source value is invalid
	if !v.Valid() {
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

func (v Value) Equals(b Value) (bool, error) {
	// One of the values are invalid
	if !v.Valid() {
		return false, v.BoundsError()
	}
	if !b.Valid() {
		return false, b.BoundsError()
	}

	// Same unit just do the comparison
	if v.Unit() == b.Unit() {
		return equals(v.v, b.v), nil
	}

	c, err := b.Transform(v.u)
	if err != nil {
		return false, err
	}
	return equals(v.v, c.v), nil
}

func equals(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}
