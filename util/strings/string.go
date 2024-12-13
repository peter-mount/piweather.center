package strings

import (
	"fmt"
	"reflect"
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

func StringDefault(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
