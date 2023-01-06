package util

import (
	"encoding/xml"
)

func String(v interface{}) string {
	if v == nil {
		return "nil"
	}

	b, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}
