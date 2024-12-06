package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
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

	to := query.Get(TO)

	req := Request{
		Metric: strings.ReplaceAll(r.Var(METRIC), "/", "."),
		At:     time2.ParseTime(query.Get(AT)).Truncate(time.Second),
		From:   time2.ParseTime(query.Get(FROM)).Truncate(time.Second),
		To:     time2.ParseTime(to).Truncate(time.Second),
		Filter: query.Get(FILTER),
	}

	if !req.From.IsZero() {
		// End provided but not a valid time, try parsing it as a duration
		if to != "" && req.To.IsZero() {
			d, err := time.ParseDuration(to)
			if err == nil {
				req.To = req.From.Add(d)
			}
		}

		// Ensure from.Before(to)
		// Can happen if client sets them wrong, but also if to was a
		// negative duration
		if !req.To.IsZero() && req.To.Before(req.From) {
			req.From, req.To = req.To, req.From
		}
	}

	return req
}

// Response returns an api.Response prepopulated with values from the Request
func (r Request) Response() api.Response {
	resp := api.Response{
		Metric: r.Metric,
	}

	if !r.At.IsZero() {
		resp.Time = &r.At
	}

	return resp
}

// Match returns true if the metric passed matches the Filter pattern
func (r Request) Match(s string) bool {
	return util.Match(s, r.Filter)
}
