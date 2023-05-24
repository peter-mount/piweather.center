package util

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	var f string
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f = "%d"
	case reflect.Float32, reflect.Float64:
		f = "%.3f"
	case reflect.String:
		f = "%s"
	default:
		f = "%v"
	}
	return fmt.Sprintf(f, v)
}

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

func StringDefault(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

func FloatDefault(a *float64, b float64) float64 {
	if a != nil {
		return *a
	}
	return b
}
