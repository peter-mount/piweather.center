package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
	"sync"
)

// Factory returns a new struct of the appropriate type
type Factory func() Instance

type Instance interface {
	GetType() string
}

var (
	components = map[string]Factory{}
	mutex      sync.Mutex
)

// Register a named Factory.
// Factory names are case-insensitive.
func Register(n string, f Factory) {
	n = strings.ToLower(n)

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := components[n]; exists {
		panic(fmt.Errorf("dash component %q already registered", n))
	}
	components[n] = f
}

// Lookup a named Factory.
// Factory names are case-insensitive.
func Lookup(n string) Factory {
	n = strings.ToLower(n)

	mutex.Lock()
	defer mutex.Unlock()

	return components[n]
}

// Unmarshal using the correct Unmarshaller version.
func Unmarshal(b []byte, o any) error {
	return yaml.Unmarshal(b, o)
}

// Decode a yaml.Node and return either the correct type based on the As field or an error
func Decode(n *yaml.Node) (Instance, error) {
	// decode the node to get the As
	comp := struct {
		Type string `yaml:"type"`
	}{}
	err := n.Decode(&comp)
	if err != nil {
		return nil, err
	}

	// Now using that type, Lookup the correct type and decode it
	f := Lookup(comp.Type)
	if f == nil {
		return nil, fmt.Errorf("unknown component type %q", comp.Type)
	}

	o := f()
	return o, n.Decode(o)
}
