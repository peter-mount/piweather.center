package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	// The url prefix, e.g. "http://localhost:8080" of the remote service
	// Note no trailing "/" as the client will add a patch starting with "/"
	Url string
}

// Make a GET to a remote service
// path - path of rest endpoint
// v - instance of object to unmarshal
// returns (true, nil) if found and v contains data
// (false, nil) if not found or (false, error ) on error
func (c *Client) get(path string, v interface{}) (bool, error) {

	if resp, err := http.Get(c.Url + path); err != nil {
		return false, err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			return false, nil
		}

		if body, err := io.ReadAll(resp.Body); err != nil {
			return false, err
		} else {
			err = json.Unmarshal(body, v)
			return err == nil, err
		}
	}
}
