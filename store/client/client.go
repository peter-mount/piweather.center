package client

import (
	"bytes"
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"io"
	"net/http"
)

type Client struct {
	// The url prefix, e.g. "http://localhost:8080" of the remote service
	// Note no trailing "/" as the client will add a patch starting with "/"
	Url string
	// If true then use internal protocol rather than JSON
	Internal bool
}

// Make a GET to a remote service
// path - path of rest endpoint
// v - instance of object to unmarshal
// returns (true, nil) if found and v contains data
// (false, nil) if not found or (false, error ) on error
func (c *Client) get(path string, v interface{}) (bool, error) {
	resp, err := http.Get(c.Url + path)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return c.getResult(resp, v, false)
}

func (c *Client) post(path string, req []byte, v interface{}) (bool, error) {

	contentType := "application/json"
	if c.Internal {
		contentType = "application/binary"
	}

	resp, err := http.Post(c.Url+path, contentType, bytes.NewReader(req))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return c.getResult(resp, v, c.Internal)
}

func (c *Client) getResult(resp *http.Response, v interface{}, internal bool) (bool, error) {
	if resp.StatusCode == 404 {
		return false, nil
	}

	defer resp.Body.Close()

	// If requesting internal, and it's a Readable result then use it
	if internal {
		if r, ok := v.(api.Readable); ok {
			err := r.Read(resp.Body)
			if err != nil {
				log.Printf("error %v", err)
				return false, err
			}
			return true, nil
		}
	}

	// Default JSON
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return true, json.Unmarshal(body, v)
}
