package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"strings"
	"time"
)

// Request holds the various optional query parameters to allow
// for filtering or specifying time ranges
type Request struct {
	Metric string    // The metric to work on, "" for all
	At     time.Time // Specific time required
	From   time.Time // Start time
	To     time.Time // End time
	Filter string    // Filter metrics with pattern, "" for match all
}

// GetRequest returns a Request populated with values from the inbound request
func GetRequest(r *rest.Rest) Request {
	query := r.Request().URL.Query()
	return Request{
		Metric: strings.ReplaceAll(r.Var(METRIC), "/", "."),
		At:     unit.ParseTime(query.Get(AT)).Truncate(time.Second),
		From:   unit.ParseTime(query.Get(FROM)).Truncate(time.Second),
		To:     unit.ParseTime(query.Get(TO)).Truncate(time.Second),
		Filter: query.Get(FILTER),
	}
}

// Response returns an api.Response prepopulated with values from the Request
func (r Request) Response() api.Response {
	resp := api.Response{
		Metric: r.Metric,
	}

	if !r.At.IsZero() {
		resp.Time = &r.At
	}

	if !r.From.IsZero() {
		resp.From = &r.From
	}

	if !r.To.IsZero() {
		resp.To = &r.To
	}

	return resp
}

// Match returns true if the metric passed matches the Filter pattern
func (r Request) Match(s string) bool {
	return util.Match(s, r.Filter)
}
