package strings

import (
	"reflect"
	"strconv"
)

func ToFloat64(v interface{}) (float64, bool) {
	if v == nil {
		return 0, false
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		return float64(v.(int)), true
	case reflect.Int64:
		return float64(v.(int64)), true
	case reflect.Float64:
		return v.(float64), true
	case reflect.String:
		s := v.(string)
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f, true
		}
		if i, err := strconv.Atoi(s); err == nil {
			return float64(i), true
		}
	}
	return 0, false
}

func FloatDefault(a *float64, b float64) float64 {
	if a != nil {
		return *a
	}
	return b
}
