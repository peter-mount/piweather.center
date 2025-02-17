package strings

import (
	"math"
	"reflect"
	"strconv"
	"strings"
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

// FormatFloatN returns a float with n decimal places but this will remove any trailing zeros, or even the
// decimal point if f is an integer. This is used in places like SVG where we want to reduce the size of the output
// by not having too many unnecessary characters.
func FormatFloatN(f float64, precision int) string {
	s := strconv.FormatFloat(f, 'f', precision, 64)
	s = strings.TrimRight(s, "0")
	return strings.TrimSuffix(s, ".")
}

// FormatFloatDown is similar to strconv.FormatFloat with type 'f' or "%f" in Sprintf except that this will round down
// the result instead of up which is the default in Go.
//
// digits the number of digits for the integer portion. If this is greater than 1 then the integer portion will be
// padded by zeros to ensure it has at least that number of digits.
//
// precision the number of decimal places
func FormatFloatDown(f float64, digits, precision int) string {
	fI, fF := math.Modf(f)
	s := Itoa(int(fI), digits)
	if precision > 0 {
		s = s + "." + Itoa(int(math.Floor(fF*math.Pow(10, float64(precision)))), precision)
	}
	return s
}
