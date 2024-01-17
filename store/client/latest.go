package client

import (
	"github.com/peter-mount/piweather.center/store/api"
	"strings"
	"time"
)

func (c *Client) queryMetric(metric string, query ...string) (*api.Response, error) {
	path := "/metric"
	if metric != "" {
		path = path + "/" + strings.ReplaceAll(metric, ".", "/")
	}
	if len(query) > 0 {
		path = path + "?" + strings.Join(query, "&")
	}

	resp := &api.Response{}
	if found, err := c.get(path, resp); err != nil {
		return nil, err
	} else if found {
		return resp, nil
	} else {
		return nil, nil
	}
}

func (c *Client) LatestMetrics() (*api.Response, error) {
	return c.Metric("")
}

func (c *Client) LatestMetricsAt(t time.Time) (*api.Response, error) {
	return c.MetricAt("", t)
}

func (c *Client) Metric(m string) (*api.Response, error) {
	return c.queryMetric(m, "latest")
}

func (c *Client) MetricToday(m string) (*api.Response, error) {
	return c.queryMetric(m, "today")
}

func (c *Client) MetricTodayUTC(m string) (*api.Response, error) {
	return c.queryMetric(m, "todayUTC")
}

func (c *Client) MetricYesterday(m string) (*api.Response, error) {
	return c.queryMetric(m, "yesterday")
}

func (c *Client) MetricYesterdayUTC(m string) (*api.Response, error) {
	return c.queryMetric(m, "yesterdayUTC")
}

func (c *Client) MetricAt(m string, t time.Time) (*api.Response, error) {
	return c.queryMetric(m, "at="+t.Format(time.RFC3339))
}

func (c *Client) Between(m string, f, t time.Time) (*api.Response, error) {
	return c.queryMetric(m, "from="+f.Format(time.RFC3339), "to="+t.Format(time.RFC3339))
}
