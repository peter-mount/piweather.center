package strings

import (
	"strconv"
	"strings"
)

// Itoa wraps strconv.Itoa but ensures that we have at least p digits, prefixing the result with as many '0' as necessary.
func Itoa(i int, p int) string {
	s := strconv.Itoa(i)
	if p > len(s) {
		s = strings.Repeat("0", p) + s
		s = s[len(s)-p:]
	}
	return s
}
