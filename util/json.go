package util

// Finds and returns a named value.
// This will return the object containing the value as well as the value
// or nil if the entry does not exist
func findJsonValue(r map[string]interface{}, n ...string) (map[string]interface{}, interface{}, bool) {
	var o map[string]interface{}
	var v interface{}
	for _, k := range n {
		if o == nil {
			o = r
		} else if a, ok := v.(map[string]interface{}); ok {
			o = a
		} else {
			// Not an object so give up
			return nil, nil, false
		}

		var e bool
		v, e = o[k]

		if !e || v == nil {
			return nil, nil, false
		}
	}
	return o, v, true
}

func GetJsonObjectValue(r map[string]interface{}, n ...string) (interface{}, bool) {
	_, v, exists := findJsonValue(r, n...)
	return v, exists
}

func GetJsonObject(o map[string]interface{}, n ...string) (map[string]interface{}, bool) {
	v, e := GetJsonObjectValue(o, n...)
	if e {
		if a, ok := v.(map[string]interface{}); ok {
			return a, true
		}
	}
	return nil, false
}

func GetJsonArray(o map[string]interface{}, n ...string) ([]interface{}, bool) {
	v, e := GetJsonObjectValue(o, n...)
	if e {
		if a, ok := v.([]interface{}); ok {
			return a, ok
		}
		var a []interface{}
		a = append(a, v)
		return a, true
	}
	return nil, false
}

// ForceJsonArray forces an entry to be a json array. If an entry is an object or value then
// it will be wrapped within a singular array
func ForceJsonArray(r map[string]interface{}, n ...string) {
	forceJsonArray(r, n, 0, len(n)-1)
}

func forceJsonArray(r map[string]interface{}, n []string, i, j int) {
	v, e := r[n[i]]
	if e {
		if i == j {
			if _, ok := v.([]interface{}); !ok {
				var a []interface{}
				a = append(a, v)
				r[n[len(n)-1]] = a
			}
			return
		}

		if a, ok := v.([]interface{}); ok {
			for _, e := range a {
				if o, ok := e.(map[string]interface{}); ok {
					forceJsonArray(o, n, i+1, j)
				}
			}
		} else if o, ok := v.(map[string]interface{}); ok {
			forceJsonArray(o, n, i+1, j)
		}
	}
}

// ForceJsonObject If an entry is "" then replace it with {} - seen in stations feed for InformationSystems
func ForceJsonObject(r map[string]interface{}, n ...string) {
	forceJsonObject(r, n, 0, len(n)-1)
}

func forceJsonObject(r map[string]interface{}, n []string, i, j int) {
	v, e := r[n[i]]
	if e {
		if i == j {
			if s, ok := v.(string); ok {
				if s == "" {
					r[n[i]] = make(map[string]interface{})
				}
			}
			return
		}

		if a, ok := v.([]interface{}); ok {
			for _, e := range a {
				if o, ok := e.(map[string]interface{}); ok {
					forceJsonObject(o, n, i+1, j)
				}
			}
		} else if o, ok := v.(map[string]interface{}); ok {
			forceJsonObject(o, n, i+1, j)
		}
	}
}
