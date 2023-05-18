package util

import "sort"

// EndpointEntry contains info exposed to the templates
type EndpointEntry struct {
	Id       string
	Name     string
	Endpoint string
	Method   string
	Protocol string
}

type Endpoints struct {
	e []EndpointEntry
}

func (s *Endpoints) Add(e EndpointEntry) {
	s.e = append(s.e, e)
}

func (s *Endpoints) Get() []EndpointEntry {
	sort.SliceStable(s.e, func(i, j int) bool {
		return s.e[i].Id < s.e[j].Id
	})
	return s.e
}
