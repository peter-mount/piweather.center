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
	Metric  api.Metric          `json:"metric"`
	Actions map[string][]string `json:"actions"`
}

func (r *Response) Add(t, id string) {
	if r.Actions == nil {
		r.Actions = make(map[string][]string)
	}
	r.Actions[t] = append(r.Actions[t], id)
}

func (r *Response) Json() ([]byte, bool) {
	if r.Actions != nil && len(r.Actions) > 0 {
		b, err := json.Marshal(r)
		return b, err == nil
	}
	return nil, false
}
