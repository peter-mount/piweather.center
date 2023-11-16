package api

import (
	"sort"
	"time"
)

type Response struct {
	Status  int           `json:"status,omitempty" xml:"status,attr,omitempty"`
	Time    *time.Time    `json:"time,omitempty" xml:"time,attr,omitempty"`
	From    *time.Time    `json:"from,omitempty" xml:"from,attr,omitempty"`
	To      *time.Time    `json:"to,omitempty" xml:"to,attr,omitempty"`
	Message string        `json:"message,omitempty" xml:"message,omitempty"`
	Source  string        `json:"source,omitempty" xml:"source,omitempty"`
	Metric  string        `json:"metric,omitempty" xml:"metric,omitempty"`
	Results []MetricValue `json:"results,omitempty" xml:"results,omitempty"`
	Result  *MetricValue  `json:"result,omitempty" xml:"result,omitempty"`
	Metrics []Metric      `json:"metrics,omitempty" xml:"metrics,omitempty"`
}

// Sort ensures the Response.Results and Response.Metrics slices
// are in time order
func (r Response) Sort() Response {

	// Sort results by time
	sort.SliceStable(r.Results, func(i, j int) bool {
		return r.Results[i].Time.Before(r.Results[j].Time)
	})

	// Sort metrics by metric id then by time
	sort.SliceStable(r.Metrics, func(i, j int) bool {
		a, b := r.Metrics[i], r.Metrics[j]
		return a.Metric < b.Metric || a.Time.Before(a.Time)
	})

	// Ensure from,to fields are valid
	if r.From != nil && r.To != nil && r.From.After(*r.To) {
		r.From, r.To = r.To, r.From
	}

	return r
}
