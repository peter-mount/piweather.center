package model

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/store/api"
	"strconv"
)

// Processor represents a type that can process an api.Metric
type Processor interface {
	Process(api.Metric, *Response)
}

type Init interface {
	Init(db string)
}

type Response struct {
	Uid     string                                      `json:"uid"`     // UID of Dashboard
	Actions map[string]map[string]map[string]api.Metric `json:"actions"` // Actions for this metric
}

type Action struct {
	ID    string                `json:"id"`
	Index map[string]api.Metric `json:"index"`
}

// IsValid returns true if the Action should be pushed to clients
func (a Action) IsValid() bool {
	return a.ID != "" && a.Index != nil && len(a.Index) > 0
}

// Add an indexed metric
func (a Action) Add(i int, m api.Metric) Action {
	return a.AddSuffix(i, "", m)
}

// AddSuffix adds an indexed metric with the supplied suffix to its id.
func (a Action) AddSuffix(i int, s string, m api.Metric) Action {
	if a.Index == nil {
		a.Index = make(map[string]api.Metric)
	}
	a.Index[strconv.Itoa(i)+s] = m
	return a
}

func (r *Response) Add(t string, a Action) {
	if a.IsValid() {
		if r.Actions == nil {
			r.Actions = make(map[string]map[string]map[string]api.Metric)
		}
		id := r.Actions[t]
		if id == nil {
			id = make(map[string]map[string]api.Metric)
		}
		id[a.ID] = a.Index
		r.Actions[t] = id
	}
}

func (r *Response) Json() ([]byte, bool) {
	if r.Actions != nil && len(r.Actions) > 0 {
		b, err := json.Marshal(r)
		return b, err == nil
	}
	return nil, false
}
