package api

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"net/http"
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

func (r *Response) Close() error {
	if r != nil {
		r.Results = nil
		r.Result = nil
		r.Metrics = nil
	}
	return nil
}

// Submit the Response to the client.
// It will also ensure that the Response.Results and Response.Metrics slices
// are sorted in metric/time order
func (r *Response) Submit(rs *rest.Rest) error {
	var p time2.Period

	if len(r.Results) > 0 {
		// Sort results by time
		sort.SliceStable(r.Results, func(i, j int) bool {
			return r.Results[i].Time.Before(r.Results[j].Time)
		})

		// Results is for 1 metric so as it's sorted in time order
		// we can just add the first/last entry to the Period
		p = p.Add(r.Results[0].Time).
			Add(r.Results[len(r.Results)-1].Time)
	}

	if len(r.Metrics) > 0 {
		// Sort metrics by metric id then by time
		sort.SliceStable(r.Metrics, func(i, j int) bool {
			a, b := r.Metrics[i], r.Metrics[j]
			return a.Metric < b.Metric || a.Time.Before(b.Time)
		})

		// We have to scan Metrics as it's ordered by metric then time
		for _, e := range r.Metrics {
			p = p.Add(e.Time)
		}
	}

	if !p.IsZero() {
		f, t := p.Range()
		r.From = &f
		r.To = &t
	}

	// If not set then set Status to 200 OK
	if r.Status == 0 {
		r.Status = http.StatusOK
	}

	rs.Status(r.Status).
		ContentType(rs.GetHeader("Content-Type")).
		Value(r)

	return nil
}
