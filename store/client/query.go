package client

import (
	"github.com/peter-mount/piweather.center/store/api"
)

func (c *Client) Query(s string) (*api.Result, error) {
	resp := &api.Result{}
	if found, err := c.post("/query", []byte(s), resp); err != nil {
		return nil, err
	} else if found {
		return resp, nil
	} else {
		return nil, nil
	}
}
