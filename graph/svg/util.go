package svg

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

// AttrMerge taks a var ags list of strings and appends another slice to it
func AttrMerge(a []string, s ...string) []string {
	return append(s, a...)
}

// Attr returns a formatted attribute
func Attr(n, f string, a ...interface{}) string {
	s := strings.TrimSpace(fmt.Sprintf(f, a...))
	if s == "" {
		return s
	}
	return n + "=\"" + s + "\""
}

func AttrN(n string, f float64) string { return Attr(n, Number(f)) }

// Number formats f to 0, 1 or 2 decimal places ensuring we do not have trailing 0's
func Number(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	switch {
	// Special case
	case value.Equal(f, 0):
		return "0"

	// xx.00 -> xx
	case strings.HasSuffix(s, ".00"):
		return s[:len(s)-3]
	// xx.x0 -> xx.x
	case strings.HasSuffix(s, "0"):
		return s[:len(s)-1]
	// xx.xx
	default:
		return s
	}
}

func CData(f string) string {
	return "<![CDATA[" + f + "]]>"
}
