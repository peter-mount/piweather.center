package value

import (
	"fmt"
	"strings"
)

type Transformer func(f float64) (float64, error)

var transformers = make(map[string]Transformer)

func transformName(from, to Unit) string {
	return strings.ToLower(from.name + "->" + to.name)
}

func NewTransform(from, to Unit, t Transformer) {
	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	if _, exists := transformers[n]; exists {
		panic(fmt.Errorf("transform %q already defined", n))
	}
	transformers[n] = t
}

func GetTransform(from, to Unit) (Transformer, error) {
	mutex.Lock()
	defer mutex.Unlock()
	n := transformName(from, to)
	if t, exists := transformers[n]; exists {
		return t, nil
	}

	return nil, fmt.Errorf("transform %q not defined", n)
}

func Transform(f float64, from Unit, to Unit) (float64, error) {
	t, err := GetTransform(from, to)
	if err != nil {
		return 0, err
	}
	return t(f)
}
