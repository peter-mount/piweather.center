package client

import (
	"bytes"
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

func (c *Client) post(path string, req []byte, v interface{}) (bool, error) {
	resp, err := http.Post(c.Url+path, "application/json", bytes.NewReader(req))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return false, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return true, json.Unmarshal(body, v)
}
