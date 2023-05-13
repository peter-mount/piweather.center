package value

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

type Unit struct {
	name   string  // Name of Unit, e.g. "Celsius"
	unit   string  // Short name of Unit e.g. "°C"
	format string  // Format for Sprintf
	min    float64 // min valid value
	max    float64 // max valid value
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) Unit() string {
	return u.unit
}

func (u Unit) String(f float64) string {
	return fmt.Sprintf(u.format, f, u.unit)
}

func (u Unit) Valid(f float64) bool {
	return f >= u.min && f <= u.max
}

func (u Unit) BoundsError(f float64) error {
	if u.Valid(f) {
		return nil
	}

	lb, ub := u.min > -math.MaxFloat64, u.max < math.MaxFloat64
	switch {
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

func (u Unit) Value(v float64) Value {
	return Value{v: v, u: u}
}

var mutex sync.Mutex
var units = make(map[string]Unit)

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

func NewBoundedUnit(name, unit, format string, min, max float64) Unit {
	u := NewUnit(name, unit, format)
	if min > max {
		min, max = max, min
	}
	u.min = min
	u.max = max
	return u
}

func NewLowerBoundUnit(name, unit, format string, min float64) Unit {
	return NewBoundedUnit(name, unit, format, min, math.MaxFloat64)
}

func NewUpperBoundUnit(name, unit, format string, max float64) Unit {
	return NewBoundedUnit(name, unit, format, -math.MaxFloat64, max)
}

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
	None    = NewUnit("None", "", Dp0)
	Percent = NewUnit("Percent", "%", Dp0)
)
