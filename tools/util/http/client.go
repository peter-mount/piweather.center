package http

import (
	"net/http"
)

type Client struct {
	Url    string // Url to connect to
	Method Method // Method to use
}

type ResponseHandler func(response *http.Response) error

func (c *Client) get(respHandler ResponseHandler) error {
	resp, err := http.Get(c.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return respHandler(resp)
}

func (c *Client) post(body []byte, respHandler ResponseHandler) error {

}
