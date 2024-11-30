package state

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"strconv"
)

// ResponseType is a map of ResponseComponent keyed by component type
type ResponseType map[string]*ResponseComponentId

func (r *ResponseType) ResponseType(k string) *ResponseComponentId {
	v, exists := (*r)[k]
	if !exists {
		v = &ResponseComponentId{}
		(*r)[k] = v
	}
	return v
}

// ResponseComponentId is a map of ResponseMetric keyed by componentId
type ResponseComponentId map[string]*ResponseMetric

func (r *ResponseComponentId) ResponseComponentId(k string) *ResponseMetric {
	v, exists := (*r)[k]
	if !exists {
		v = &ResponseMetric{}
		(*r)[k] = v
	}
	return v
}

// ResponseMetric is a map of api.Metric keyed by index which is an int and a string suffix
type ResponseMetric map[string]api.Metric

func (r *ResponseMetric) Set(i int, s string, v api.Metric) {
	(*r)[strconv.Itoa(i)+s] = v
}

type ResponseComponent interface {
	station.ComponentId
	station.ComponentType
}

type Response struct {
	Uid     string       `json:"uid"`     // UID of Dashboard
	Actions ResponseType `json:"actions"` // Actions for this metric
}

func (r *Response) SetComponent(c ResponseComponent, i ComponentEntryIndex, m api.Metric) {
	r.Set(c.GetType(), c.GetID(), i.Index, i.Suffix, m)
}

// Set sets an api.Metric against componentType t with componentId id to the field i with suffix s
func (r *Response) Set(t, id string, i int, s string, m api.Metric) {
	if r.Actions == nil {
		r.Actions = make(ResponseType)
	}
	r.Actions.ResponseType(t).ResponseComponentId(id).Set(i, s, m)
}

func (r *Response) Json() ([]byte, bool) {
	if r.Actions != nil && len(r.Actions) > 0 {
		b, err := json.Marshal(r)
		return b, err == nil
	}
	return nil, false
}
