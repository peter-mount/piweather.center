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
	name   string  // Name of Unit, e.g. "Celsius"
	unit   string  // Short name of Unit e.g. "°C"
	format string  // Format for Sprintf
	min    float64 // min valid value
	max    float64 // max valid value
}

// Name of the Unit. e.g. "Celsius"
func (u Unit) Name() string {
	return u.name
}

// Unit for strings, e.g. "°C"
func (u Unit) Unit() string {
	return u.unit
}

// String returns a float64 in it's supported format for this unit.
// This will be the value with the string from Unit() appended to it.
func (u Unit) String(f float64) string {
	return fmt.Sprintf(u.format, f, u.unit)
}

// Equals returns true if the unit's names are identical.
// This is case-insensitive.
func (u Unit) Equals(b Unit) bool {
	return strings.ToLower(u.name) == strings.ToLower(b.name)
}

// Valid returns true if f is within the bounds of this unit.
// e.g. For Kelvin f must be >=0 as negative values would be invalid as you cannot have
// temperatures below Absolute Zero.
//
// If the value is NaN or either Infinity then this returns false.
func (u Unit) Valid(f float64) bool {
	return !math.IsNaN(f) && !math.IsInf(f, 0) && within(f, u.min, u.max)
}

// BoundsError returns an error if the unit is outside its bounds, NaN or Infinity
func (u Unit) BoundsError(f float64) error {
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
func (u Unit) Value(v float64) Value {
	return Value{v: v, u: u}
}

// NewUnit creates a new Unit, registering it with the system.
func NewUnit(name, unit, format string) Unit {
	mutex.Lock()
	defer mutex.Unlock()
	n := strings.ToLower(name)
	if _, exists := units[n]; exists {
		panic(fmt.Errorf("unit %q already exists", name))
	}
	u := Unit{name: name, unit: unit, format: format, min: -math.MaxFloat64, max: math.MaxFloat64}
	units[n] = u
	return u
}

// NewBoundedUnit creates a new Unit which has both min and max values.
func NewBoundedUnit(name, unit, format string, min, max float64) Unit {
	u := NewUnit(name, unit, format)
	if min > max {
		min, max = max, min
	}
	u.min = min
	u.max = max
	return u
}

// NewLowerBoundUnit creates a new Unit which has a lower limit on it's permitted values.
func NewLowerBoundUnit(name, unit, format string, min float64) Unit {
	return NewBoundedUnit(name, unit, format, min, math.MaxFloat64)
}

// NewUpperBoundUnit creates a new Unit which has an upper limit on it's permitted values.
func NewUpperBoundUnit(name, unit, format string, max float64) Unit {
	return NewBoundedUnit(name, unit, format, -math.MaxFloat64, max)
}

// GetUnit returns a registered Unit based on its name.
// If the unit is not registered then this returns (nil,false).
// Names are case insensitive.
func GetUnit(name string) (Unit, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	u, e := units[strings.ToLower(name)]
	return u, e
}

const (
	Dp0 = "%.0f%s"
	Dp1 = "%.1f%s"
	Dp2 = "%.2f%s"
	Dp3 = "%.3f%s"
)

var (
	// Integer is effectively an integer value with no unit.
	Integer = NewUnit("Integer", "", Dp0)
	// Percent is a Unit bounded by 0..100 and has no decimal places
	Percent = NewBoundedUnit("Percent", "%", Dp0, 0, 100)
)
