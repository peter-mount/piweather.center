package util

import (
	"encoding/json"
	"sort"
)

type StringSet map[string]interface{}

func NewStringSet() StringSet {
	return make(map[string]interface{})
}

func (a *StringSet) Add(s string) {
	if a != nil {
		(*a)[s] = true
	}
}

func (a *StringSet) AddAll(s ...string) {
	if a != nil {
		for _, e := range s {
			(*a)[e] = true
		}
	}
}

func (a *StringSet) Contains(s string) bool {
	if a == nil {
		return false
	}
	_, exists := (*a)[s]
	return exists
}

func (a *StringSet) Size() int {
	if a == nil {
		return 0
	}
	return len(*a)
}

func (a *StringSet) ForEach(f func(string) error) error {
	if a != nil {
		for k, _ := range *a {
			if err := f(k); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *StringSet) MarshalJSON() ([]byte, error) {
	var r []string
	for k, _ := range *a {
		r = append(r, k)
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return json.Marshal(r)
}

func (a *StringSet) UnmarshalJSON(data []byte) error {
	var r []string
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	a.AddAll(r...)
	return nil
}
