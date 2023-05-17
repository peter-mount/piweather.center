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
	id       string  // Unique ID, case insensitive
	category string  // Category of Unit
	name     string  // Name of Unit, e.g. "Celsius"
	unit     string  // Short name of Unit e.g. "°C"
	format   string  // Format for Sprintf
	min      float64 // min valid value
	max      float64 // max valid value
}

func (u *Unit) ID() string { return u.id }

func (u *Unit) Category() string { return u.category }

// Name of the Unit. e.g. "Celsius"
func (u *Unit) Name() string { return u.name }

// Unit for strings, e.g. "°C"
func (u *Unit) Unit() string { return u.unit }

func (u *Unit) HasMax() bool { return u.max < math.MaxFloat64 }

func (u *Unit) Max() float64 { return u.max }

func (u *Unit) HasMin() bool { return u.min > -math.MaxFloat64 }

func (u *Unit) Min() float64 { return u.min }

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
	if u == nil || b == nil {
		return false
	}
	return strings.ToLower(u.name) == strings.ToLower(b.name)
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
		return fmt.Errorf("%s "+u.format+" out of bounds", u.name, f, u.unit)
	}
}

// Value returns a Value with this Unit.
// This is the only method to create a Value.
func (u *Unit) Value(v float64) Value {
	return Value{v: v, u: u}
}

// NewUnit creates a new Unit, registering it with the system.
func NewUnit(id, category, name, unit, format string) *Unit {
	mutex.Lock()
	defer mutex.Unlock()
	n := strings.ToLower(id)
	if _, exists := units[n]; exists {
		panic(fmt.Errorf("unit %q already exists", id))
	}
	u := &Unit{id: n, category: category, name: name, unit: unit, format: format, min: -math.MaxFloat64, max: math.MaxFloat64}
	units[n] = u
	return u
}

// NewBoundedUnit creates a new Unit which has both min and max values.
func NewBoundedUnit(id, category, name, unit, format string, min, max float64) *Unit {
	u := NewUnit(id, category, name, unit, format)
	if min > max {
		min, max = max, min
	}
	u.min = min
	u.max = max
	return u
}

// NewLowerBoundUnit creates a new Unit which has a lower limit on it's permitted values.
func NewLowerBoundUnit(id, category, name, unit, format string, min float64) *Unit {
	return NewBoundedUnit(id, category, name, unit, format, min, math.MaxFloat64)
}

// NewUpperBoundUnit creates a new Unit which has an upper limit on it's permitted values.
func NewUpperBoundUnit(id, category, name, unit, format string, max float64) *Unit {
	return NewBoundedUnit(id, category, name, unit, format, -math.MaxFloat64, max)
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

const (
	// Dp0 format for Unit's that have 0 decimal places
	Dp0 = "%.0f%s"
	// Dp1 format for Unit's that have 1 decimal places
	Dp1 = "%.1f%s"
	// Dp2 format for Unit's that have 2 decimal places
	Dp2 = "%.2f%s"
	// Dp3 format for Unit's that have 3 decimal places
	Dp3 = "%.3f%s"
)

var (
	// Integer is effectively an integer value with no unit.
	Integer = NewUnit("Integer", "Misc", "Integer", "", Dp0)
	// Float is a value with no unit
	Float = NewUnit("Float", "Misc", "Float", "", Dp3)
	// Percent is a Unit bounded by 0..100 and has no decimal places
	Percent = NewBoundedUnit("Percent", "Misc", "Percent", "%", Dp0, 0, 100)
)
