package model

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/store/api"
)

// Processor represents a type that can process an api.Metric
type Processor interface {
	Process(api.Metric, *Response)
}

type Response struct {
	Uuid    string                      `json:"uuid"`    // UUID of Dashboard
	Metric  api.Metric                  `json:"metric"`  // Inbound metric
	Actions map[string]map[string][]int `json:"actions"` // Actions for this metric
}

type Action struct {
	ID    string `json:"id"`
	Index []int  `json:"index"`
}

func (a Action) IsValid() bool {
	return a.ID != "" && len(a.Index) > 0
}

func (a Action) Add(i int) Action {
	a.Index = append(a.Index, i)
	return a
}

func (r *Response) Add(t string, a Action) {
	if a.IsValid() {
		if r.Actions == nil {
			r.Actions = make(map[string]map[string][]int)
		}
		id := r.Actions[t]
		if id == nil {
			id = make(map[string][]int)
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
