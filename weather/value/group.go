package value

import (
	"fmt"
	"strings"
)

// Group represents a set of Unit's which are related.
// e.g. Units Celsius, Fahrenheit and Kelivn are members of the Temperature group.
type Group struct {
	name  string
	units []*Unit
	err   error
}

// NewGroup creates a group of Units.
//
// This will panic if there are no Transform's available between
// any sibling Unit and the base Unit, or if it cannot create a Transform
// between any sibling and another via the base Unit.
func NewGroup(name string, base *Unit, siblings ...*Unit) *Group {
	n := strings.ToLower(name)
	if _, exists := GetGroup(n); exists {
		panic(fmt.Errorf("unit %q already registered", n))
	}

	g := &Group{
		name:  name,
		err:   fmt.Errorf("not a %s", name),
		units: append([]*Unit{base}, siblings...),
	}

	// Validate the group Unit's can transform between each other
	switch len(siblings) {
	case 0:
		// Do nothing as this is a group of just 1 unit

	case 1:
		// If 1 sibling then test that a transform exists between the two.
		// If it doesn't exist then panic as a sibling must transform with
		// the base in both directions.
		if err := AssertTransformsAvailable(base, siblings[0]); err != nil {
			panic(err)
		}

	default:
		// Ensure that there's a transform between all the units using
		// newTransformations to test and fill in with the base
		newTransformations(base, siblings...)
	}

	// Finally link each Unit to the Group
	for _, unit := range g.units {
		if unit.group != nil {
			panic(fmt.Errorf("unit %q is already a member of group %q", unit.ID(), unit.group.Name()))
		}
		unit.group = g
	}

	// Register the group.
	// We cannot lock before this point as the Transformation checks share the
	// same mutex and we would deadlock ourselves
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := groups[n]; exists {
		// Should never happen but as we cannot be locked until this point
		// we should check just in-case...
		panic(fmt.Errorf("race condition: group %q created during it's own creation", n))
	}
	groups[n] = g

	return g
}

// Name of the Group
func (g *Group) Name() string {
	if g == nil {
		return "nil"
	}
	return g.name
}

// Units returns the units forming this group
func (g *Group) Units() []*Unit {
	if g == nil {
		return nil
	}
	return g.units
}

// IsUnit returns true if the unit is a member of this Group.
// If Group or Unit is nil then this returns false.
func (g *Group) IsUnit(u *Unit) bool {
	if g != nil && u != nil {
		for _, m := range g.units {
			if m.Equals(u) {
				return true
			}
		}
	}
	return false
}

// IsValue returns true if the Value's Unit is a member of this Group.
// If Group is nil then this returns false.
func (g *Group) IsValue(v Value) bool {
	return g.IsUnit(v.Unit())
}

// IsError returns true if the error was from either Assert or AssertUnit
// from this Group.
// If Group is nil then this returns false.
func (g *Group) IsError(err error) bool {
	return g != nil && g.err == err
}

// AssertUnit will return an error if the supplied Unit is not a member
// of this Group.
// If either Group or Unit are nil then this returns nil.
func (g *Group) AssertUnit(u *Unit) error {
	if g == nil || u == nil || g.IsUnit(u) {
		return nil
	}
	return g.err
}

// AssertValue will return an error if the supplied Value is not a member
// of this Group.
// If Group is nil then this returns nil.
func (g *Group) AssertValue(v Value) error {
	return g.AssertUnit(v.Unit())
}

// GetGroup returns a registered Unit based on its name.
// If the unit is not registered then this returns (nil,false).
// Names are case insensitive.
func GetGroup(id string) (*Group, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	g, e := groups[strings.ToLower(id)]
	return g, e
}

// GetGroupIDs returns a slice of all registered group ID's
func GetGroupIDs() []string {
	var r []string
	mutex.Lock()
	defer mutex.Unlock()
	for k, _ := range groups {
		r = append(r, k)
	}
	return r
}

// GetGroups returns a slice of all registered group's
func GetGroups() []*Group {
	var r []*Group
	mutex.Lock()
	defer mutex.Unlock()
	for _, g := range groups {
		r = append(r, g)
	}
	return r
}
