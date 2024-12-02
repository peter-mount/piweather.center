package station

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"strconv"
	"time"
)

type ResponseComponent interface {
	station.ComponentId
	station.ComponentType
}

type Response struct {
	Station   string                                      `json:"-"`       // Station this response is for
	Dashboard string                                      `json:"-"`       // Dashboard this response is for
	Uid       string                                      `json:"uid"`     // UID of Dashboard
	Actions   map[string]map[string]map[string]api.Metric `json:"actions"` // Actions for this metric
}

func (r *Response) IsValid() bool {
	return r != nil && r.Actions != nil
}

func (r *Response) SetComponent(c ResponseComponent, i ComponentEntryIndex, m api.Metric) {
	// Suffix "T" means we want the time
	if i.Suffix == "T" {
		// TODO force this to the dashboard TimeZone and not UTC
		m.Formatted = m.Time.Format(time.RFC3339)
	}
	r.Set(c.GetType(), c.GetID(), i.Index, i.Suffix, m)
}

// Set sets an api.Metric against componentType t with componentId id to the field i with suffix s
func (r *Response) Set(t, id string, i int, s string, m api.Metric) {
	if r.Actions == nil {
		r.Actions = make(map[string]map[string]map[string]api.Metric)
	}
	typeMap, exists := r.Actions[t]
	if !exists {
		typeMap = make(map[string]map[string]api.Metric)
		r.Actions[t] = typeMap
	}
	compMap, exists := typeMap[id]
	if !exists {
		compMap = make(map[string]api.Metric)
		typeMap[id] = compMap
	}
	compMap[strconv.Itoa(i)+s] = m
}

func (r *Response) Json() ([]byte, bool) {
	if r.Actions != nil && len(r.Actions) > 0 {
		b, err := json.Marshal(r)
		return b, err == nil
	}
	return nil, false
}
