package client

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/store/api"
	"time"
)

func (c *Client) Record(m string, t time.Time, val float64, unit string) (*api.Response, error) {
	return c.RecordMetric(api.Metric{
		Metric: m,
		Time:   t,
		Unit:   unit,
		Value:  val,
		Unix:   t.Unix(),
	})
}

func (c *Client) RecordMetric(m api.Metric) (*api.Response, error) {
	b, err := json.Marshal(&m)
	if err != nil {
		return nil, err
	}

	resp := &api.Response{}
	if found, err := c.post("/record", b, resp); err != nil {
		return nil, err
	} else if found {
		return resp, nil
	} else {
		return nil, nil
	}
}
