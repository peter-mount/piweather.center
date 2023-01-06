package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetBodyBytes returns the body of the response as a byte slice
func GetBodyBytes(response *http.Response) ([]byte, error) {
	return ioutil.ReadAll(response.Body)
}

// GetBodyJson returns the body of the response unmarshalled into a dest object
func GetBodyJson(response *http.Response, dest interface{}) error {
	body, err := GetBodyBytes(response)
	if err == nil {
		err = json.Unmarshal(body, dest)
	}
	return err
}
