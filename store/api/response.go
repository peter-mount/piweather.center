package api

import (
	"github.com/peter-mount/go-kernel/v2/rest"
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

// Submit the Response to the client.
// It will also ensure that the Response.Results and Response.Metrics slices
// are sorted in metric/time order
func (r Response) Submit(rs *rest.Rest) error {

	// Sort results by time
	sort.SliceStable(r.Results, func(i, j int) bool {
		return r.Results[i].Time.Before(r.Results[j].Time)
	})

	// Sort metrics by metric id then by time
	sort.SliceStable(r.Metrics, func(i, j int) bool {
		a, b := r.Metrics[i], r.Metrics[j]
		return a.Metric < b.Metric || a.Time.Before(b.Time)
	})

	// If from and to defined then ensure that from.Before(to) holds.
	if r.From != nil && r.To != nil && r.From.After(*r.To) {
		r.From, r.To = r.To, r.From
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
