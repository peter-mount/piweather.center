package value

import (
	"fmt"
	"math"
	"strings"
)

// Unit represents a unit of some kind.
// For example for temperature we have Kelvin, Celsius and Fahrenheit.
//
// Unit's support transformations with other units, which are registered with
// NewTransform. If no transform is registered for two units then they cannot
// be transformed.
type Unit struct {
	id        string  // Unique ID, case insensitive
	group     *Group  // Group the unit belongs to, nil for no membership
	name      string  // Name of Unit, e.g. "Celsius"
	unit      string  // Short name of Unit e.g. "°C"
	format    string  // Format for Sprintf
	precision int     // Precision of the unit
	min       float64 // min valid value
	max       float64 // max valid value
	err       error   // Error from assertion
}

func (u *Unit) ID() string { return u.id }

func (u *Unit) Category() string {
	if u.group == nil {
		return "Uncategorized"
	}
	return u.group.Name()
}

// Group returns the Group this Unit is a member of, or nil if
// it's not a member of one.
func (u *Unit) Group() *Group { return u.group }

// Name of the Unit. e.g. "Celsius"
func (u *Unit) Name() string { return u.name }

// Unit for strings, e.g. "°C"
func (u *Unit) Unit() string { return u.unit }

func (u *Unit) HasMax() bool { return u.max < math.MaxFloat64 }

func (u *Unit) Max() float64 { return u.max }

func (u *Unit) HasMin() bool { return u.min > -math.MaxFloat64 }

func (u *Unit) Min() float64 { return u.min }

func (u *Unit) Precision() int { return u.precision }

// String returns a float64 in it's supported format for this unit.
// This will be the value with the string from Unit() appended to it.
func (u *Unit) String(f float64) string {
	return fmt.Sprintf(u.format, f, u.unit)
}

func (u *Unit) PlainString(f float64) string {
	return fmt.Sprintf(u.format, f, "")
}

// Equals returns true if the unit's names are identical.
// This is case-insensitive.
func (u *Unit) Equals(b *Unit) bool {
	// nil Unit's do not equal anything
	if u == nil || b == nil {
		return false
	}
	// Either they are the same instance or their IDs are the same
	return u == b || strings.ToLower(u.id) == strings.ToLower(b.id)
}

// Valid returns true if f is Within the bounds of this unit.
// e.g. For Kelvin f must be >=0 as negative values would be invalid as you cannot have
// temperatures below Absolute Zero.
//
// If the value is NaN or either Infinity then this returns false.
func (u *Unit) Valid(f float64) bool {
	return u != nil && !math.IsNaN(f) && !math.IsInf(f, 0) && Within(f, u.min, u.max)
}

// BoundsError returns an error if the unit is outside its bounds, NaN or Infinity
func (u *Unit) BoundsError(f float64) error {
	if u.Valid(f) {
		return nil
	}

	// Check if we have bounds set. If not then min,max are set to the appropriate limits
	// of float64. As such we don't have to care about accuracy here
	lb, ub := u.min > -math.MaxFloat64, u.max < math.MaxFloat64
	switch {
	case u == nil:
		return nilErr
	case math.IsNaN(f):
		return nan
	case math.IsInf(f, 1):
		return pInf
	case math.IsInf(f, -1):
		return nInf
	case lb && ub:
		return fmt.Errorf("%s "+u.format+" out of bounds "+u.format+"..."+u.format, u.name, f, u.unit, u.min, u.unit, u.max, u.unit)
	case lb:
		return fmt.Errorf("%s "+u.format+" out of bounds "+u.format+"...", u.name, f, u.unit, u.min, u.unit)
	case ub:
		return fmt.Errorf("%s "+u.format+" out of bounds ..."+u.format, u.name, f, u.unit, u.max, u.unit)
	default:
		return fmt.Errorf("%s "+u.format+" out of bounds?", u.name, f, u.unit)
	}
}

// Value returns a Value with this Unit.
// This is the only method to create a Value.
func (u *Unit) Value(v float64) Value {
	return Value{v: v, u: u}
}

// AssertUnit will return an error if the two Unit's do not match.
// If either Unit is nil then this returns nil.
func (u *Unit) AssertUnit(b *Unit) error {
	if u == nil || b == nil || u.Equals(b) {
		return nil
	}
	return u.err
}

// AssertValue returns an error if the Value's Unit does not match this Unit.
func (u *Unit) AssertValue(v Value) error {
	return u.AssertUnit(v.Unit())
}

// IsErr returns true if the error was returned by AssertUnit or AssertValue.
func (u *Unit) IsErr(err error) bool {
	return u != nil && u.err == err
}

// NewUnit creates a new Unit, registering it with the system.
func NewUnit(id, name, unit string, precision int) *Unit {
	mutex.Lock()
	defer mutex.Unlock()
	n := strings.ToLower(id)
	if _, exists := units[n]; exists {
		panic(fmt.Errorf("unit %q already registered", n))
	}
	u := &Unit{
		id:        n,
		name:      name,
		unit:      unit,
		precision: precision,
		format:    fmt.Sprintf("%%.%df%%s", precision),
		min:       -math.MaxFloat64,
		max:       math.MaxFloat64,
		err:       fmt.Errorf("not a %s %q", id, name),
	}
	units[n] = u
	return u
}

// NewBoundedUnit creates a new Unit which has both min and max values.
func NewBoundedUnit(id, name, unit string, precision int, min, max float64) *Unit {
	u := NewUnit(id, name, unit, precision)
	if min > max {
		min, max = max, min
	}
	u.min = min
	u.max = max
	return u
}

// NewLowerBoundUnit creates a new Unit which has a lower limit on it's permitted values.
func NewLowerBoundUnit(id, name, unit string, precision int, min float64) *Unit {
	return NewBoundedUnit(id, name, unit, precision, min, math.MaxFloat64)
}

// NewUpperBoundUnit creates a new Unit which has an upper limit on it's permitted values.
func NewUpperBoundUnit(id, name, unit string, precision int, max float64) *Unit {
	return NewBoundedUnit(id, name, unit, precision, -math.MaxFloat64, max)
}

// GetUnit returns a registered Unit based on its name.
// If the unit is not registered then this returns (nil,false).
// Names are case insensitive.
func GetUnit(id string) (*Unit, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	u, e := units[strings.ToLower(id)]
	return u, e
}

func GetUnits() []*Unit {
	var r []*Unit
	mutex.Lock()
	defer mutex.Unlock()
	for _, e := range units {
		r = append(r, e)
	}
	return r
}

var (
	// Integer is effectively an integer value with no unit.
	Integer = NewUnit("Integer", "Integer", "", 0)
	// Float is a value with no unit
	Float = NewUnit("Float", "Float", "", 3)
	// Percent is a Unit bounded by 0..100 and has no decimal places
	Percent = NewBoundedUnit("Percent", "Percent", "%", 0, 0, 100)
)
