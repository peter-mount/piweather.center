package value

import (
	"errors"
	"fmt"
	"strings"
)

// Transformer is an operation that will transform a float64 between Unit's.
// This is the core of how a Unit can be transformed to another Unit.
type Transformer func(f float64) (float64, error)

// Then allows for one Transformer to pass its result to another one.
// This is used for transforming between two Units with an intermediate one.
// e.g. Fahrenheit to Kelvin is actually Fahrenheit -> Celsius -> Kelvin
func (a Transformer) Then(b Transformer) Transformer {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(f float64) (float64, error) {
		f1, err := a(f)
		if err != nil {
			return 0, err
		}
		return b(f1)
	}
}

// Of allows for a sequence of Transformer's to be chained together to form a new Transform.
//
// An example is value.Of(fahrenheitCelsius, celsiusKelvin) which creates a Transform
// that converts Fahrenheit to Kelvin by converting it to Celsius first.
func Of(transforms ...Transformer) Transformer {
	var t Transformer
	for _, e := range transforms {
		t = t.Then(e)
	}
	return t
}

func transformName(from, to *Unit) string {
	return strings.ToLower(from.ID() + "->" + to.ID())
}

// NewTransform registers a Transformer that will handle transforming from one Unit to Another.
// If a specific transform has already been registered or if both from and to represent the same Unit
// then this will panic.
func NewTransform(from, to *Unit, t Transformer) {
	// Guard to ensure from and to are different units
	if from.Equals(to) {
		panic(fmt.Errorf("transform units cannot be the same %q", from.Name()))
	}

	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	if _, exists := transformers[n]; exists {
		panic(fmt.Errorf("transform %q already defined", n))
	}
	transformers[n] = t
}

// newTransformations generates transformations to convert between two or more Unit's
// Using a base Unit as the intermediary. It will create every possible transform based on the list of units.
//
// e.g. for speed you can convert KilometersPerHour to FeetPerSecond by converting to MetersPerSecond first.
//
// If an error occurs this will panic. This is either units has less than 2 entries, or there is no
// defined Transform of a unit to or from the baseUnit.
//
// If the baseUnit is present in the unit list it will also panic because that transform will already be registered.
// If no transform exists between two units, then a transform will be created converting via the baseUnit.
func newTransformations(baseUnit *Unit, units ...*Unit) {
	if len(units) < 2 {
		panic(errors.New("must supply 2 or more Unit's to link to base"))
	}

	for _, src := range units {
		if src.Equals(baseUnit) {
			panic(fmt.Errorf("base unit %q included in list of units", baseUnit.Name()))
		}

		srcToBase, err := GetTransform(src, baseUnit)
		if err != nil {
			panic(err)
		}

		for _, dest := range units {
			if dest.Equals(baseUnit) {
				panic(fmt.Errorf("base unit %q included in list of units", baseUnit.Name()))
			}

			if !src.Equals(dest) {
				baseToDest, err := GetTransform(baseUnit, dest)
				if err != nil {
					panic(err)
				}

				if !TransformAvailable(src, dest) {
					NewTransform(src, dest, srcToBase.Then(baseToDest))
				}
			}
		}
	}
}

// NewBiTransform registers two transforms, one for u1->u2 and one for u2->u1
func NewBiTransform(u1, u2 *Unit, u1ToU2, u2ToU1 Transformer) {
	NewTransform(u1, u2, u1ToU2)
	NewTransform(u2, u1, u2ToU1)
}

// GetTransform returns the Transformer that will transform between two units.
// An error is returned if the requested transform has not been defined.
func GetTransform(from, to *Unit) (Transformer, error) {
	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	if t, exists := transformers[n]; exists {
		return t, nil
	}

	return nil, fmt.Errorf("transform %q not defined", n)
}

// TransformAvailable returns true if it's possible to transfrom between two units.
func TransformAvailable(from, to *Unit) bool {
	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	_, exists := transformers[n]
	return exists
}

// TransformsAvailable returns true if it's possible
// to transform between two units in either direction.
func TransformsAvailable(a, b *Unit) bool {
	return TransformAvailable(a, b) && TransformAvailable(b, a)
}

// AssertTransformsAvailable returns an error if it's not possible
// to transform between two units in either direction.
func AssertTransformsAvailable(a, b *Unit) error {
	_, err := GetTransform(a, b)
	if err == nil {
		_, err = GetTransform(b, a)
	}
	return err
}

// Transform will transform a float64 between two units.
// An error is returned if the requested transform has not been defined.
func Transform(f float64, from, to *Unit) (float64, error) {
	// No transform required
	if from.Equals(to) {
		return f, nil
	}

	t, err := GetTransform(from, to)
	if err != nil {
		return 0, err
	}
	return t(f)
}

// BasicTransform returns a Transformer that multiplies a value with a constant
// conversion factor.
//
// This is used for transforms like meters per second to kilometers per hour
func BasicTransform(factor float64) Transformer {
	return func(f float64) (float64, error) {
		return f * factor, nil
	}
}

// BasicInverseTransform returns a Transformer that divides a value with
// a constant conversion factor.
//
// This is the same as BasicTransform(1.0/factor)
func BasicInverseTransform(factor float64) Transformer {
	return func(f float64) (float64, error) {
		return f / factor, nil
	}
}

// NewBasicTransform is the same as NewTransform(from, to, BasicTransform(factor))
func NewBasicTransform(from, to *Unit, factor float64) {
	NewTransform(from, to, BasicTransform(factor))
}

// NewBasicBiTransform registers two transforms using a constant conversion factor.
// This allows for conversions between two units.
func NewBasicBiTransform(u1, u2 *Unit, factor float64) {
	NewTransform(u1, u2, BasicTransform(factor))
	NewTransform(u2, u1, BasicInverseTransform(factor))
}

type TransformDef struct {
	From string
	To   string
}

func GetTransforms() []TransformDef {
	mutex.Lock()
	defer mutex.Unlock()

	var defs []TransformDef
	for k, _ := range transformers {
		s := strings.Split(k, "->")
		if len(s) == 2 {
			defs = append(defs, TransformDef{
				From: s[0],
				To:   s[1],
			})
		}
	}
	return defs
}

// NopTransformer does nothing.
func NopTransformer(f float64) (float64, error) {
	return f, nil
}
