package value

import (
	"fmt"
	"strings"
)

// Transformer is an operation that will transform a float64.
// This is the core of how a Unit can be transformed to another Unit.
type Transformer func(f float64) (float64, error)

// Then allows for one transformer to pass its result to another one.
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

// Chain allows for a sequence of Transformer's to be chained together to form a new Transform.
//
// An example is value.Chain(fahrenheitCelsius, celsiusKelvin) which creates a Transform
// that converts Fahrenheit to Kelvin by converting it to Celsius first.
func Chain(transforms ...Transformer) Transformer {
	var t Transformer
	for _, e := range transforms {
		t = t.Then(e)
	}
	return t
}

var transformers = make(map[string]Transformer)

func transformName(from, to Unit) string {
	return strings.ToLower(from.name + "->" + to.name)
}

// NewTransform registers a Transformer that will handle transforming from one Unit to Another.
// If a specific transform has already been registered or if both from and to represent the same Unit
// then this will panic.
func NewTransform(from, to Unit, t Transformer) {
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

// GetTransform returns the Transformer that will transform between two units.
// An error is returned if the requested transform has not been defined.
func GetTransform(from, to Unit) (Transformer, error) {
	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	if t, exists := transformers[n]; exists {
		return t, nil
	}

	return nil, fmt.Errorf("transform %q not defined", n)
}

// Transform will transform a float64 between two units.
// An error is returned if the requested transform has not been defined.
func Transform(f float64, from Unit, to Unit) (float64, error) {
	// No transform required
	if from == to {
		return f, nil
	}

	t, err := GetTransform(from, to)
	if err != nil {
		return 0, err
	}
	return t(f)
}
