package client

import "github.com/peter-mount/piweather.center/store/api"

func (c *Client) LatestMetrics() (*api.Response, error) {
	resp := &api.Response{}
	if found, err := c.get("/metric?latest", resp); err != nil {
		return nil, err
	} else if found {
		return resp, nil
	} else {
		return nil, nil
	}
}
