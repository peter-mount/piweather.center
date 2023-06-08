package state

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client interface {
	GetState(id string) (*Station, error)
}

type client struct {
	url string
}

func New(url string) Client {
	return &client{url: url}
}

// Make a GET to a remote service
// path - path of rest endpoint
// v - instance of object to unmarshal
// returns (true, nil) if found and v contains data
// (false, nil) if not found or (false, error ) on error
func (c *client) get(path string, v interface{}) (bool, error) {
	if resp, err := http.Get(c.url + path); err != nil {
		return false, err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			return false, nil
		}

		body, err := io.ReadAll(resp.Body)
		if err == nil {
			err = json.Unmarshal(body, v)
		}
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (c *client) GetState(id string) (*Station, error) {
	msg := &Station{}
	found, err := c.get("/api/station/"+id+".json", msg)
	if err != nil {
		return nil, err
	}

	if found {
		return msg, nil
	} else {
		return nil, nil
	}
}
