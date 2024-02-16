package location

import (
	"github.com/alecthomas/participle/v2"
	"strings"
	"sync"
)

type Map interface {
	GetLocation(n string) *Location
	GetLocations() []*Location
	SetLocation(l *Location) (Map, bool)
	Merge(b Map) (Map, error)
}

type basicMap struct {
	mutex     sync.Mutex
	locations map[string]*Location
}

func NewMap() Map {
	return newMap()
}

func newMap() *basicMap {
	return &basicMap{
		locations: make(map[string]*Location),
	}
}

func (s *basicMap) GetLocation(n string) *Location {
	if s == nil {
		return nil
	}

	n = strings.ToLower(n)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.locations[n]
}

func (s *basicMap) GetLocations() []*Location {
	if s == nil {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	var r []*Location
	for _, l := range s.locations {
		r = append(r, l)
	}
	return r
}

func (s *basicMap) SetLocation(l *Location) (Map, bool) {
	if l == nil {
		return s, false
	}

	if s == nil {
		return NewMap().SetLocation(l)
	}

	l.Name = strings.ToLower(l.Name)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exists := s.locations[l.Name]; exists {
		return s, false
	}
	s.locations[l.Name] = l
	return s, true
}

func (s *basicMap) Merge(b Map) (Map, error) {
	if b == nil {
		return s, nil
	}

	if s == nil {
		s = newMap()
	}

	// Merge the state, dealing with id clashes
	for _, l := range b.GetLocations() {
		if _, added := s.SetLocation(l); !added {
			return nil, participle.Errorf(l.Pos, "location %q already defined", l.Name)
		}
	}

	return s, nil
}
