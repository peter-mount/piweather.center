package payload

import "net/url"

func UnmarshalPost(b []byte, m *map[string]interface{}) error {
	values, err := url.ParseQuery(string(b))
	if err != nil {
		return err
	}

	// Use Get so we get the first associated value, not an array
	for k, _ := range values {
		(*m)[k] = values.Get(k)
	}

	return nil
}
