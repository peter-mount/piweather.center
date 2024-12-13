package strings

import "strings"

// In returns true if s is one of the strings in p.
// This is a case-insensitive test.
func In(s string, p ...string) bool {
	s = strings.ToLower(s)
	for _, e := range p {
		if s == strings.ToLower(e) {
			return true
		}
	}
	return false
}
