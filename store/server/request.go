package server

import (
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"strings"
	"time"
)

type Request struct {
	Metric string
	At     time.Time
	From   time.Time
	To     time.Time
	Filter string
}

func getRequest(r *rest.Rest) Request {
	query := r.Request().URL.Query()
	return Request{
		Metric: strings.ReplaceAll(r.Var(METRIC), "/", "."),
		At:     unit.ParseTime(query.Get(AT)).Truncate(time.Second),
		From:   unit.ParseTime(query.Get(FROM)).Truncate(time.Second),
		To:     unit.ParseTime(query.Get(TO)).Truncate(time.Second),
		Filter: query.Get(FILTER),
	}
}

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

func (r Request) Match(s string) bool {
	return util.Match(s, r.Filter)
}
