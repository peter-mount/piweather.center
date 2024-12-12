package strings

import (
	"strings"
)

// Match tests s against pattern pat.
//
// If p starts and ends with '*' then this means contains the text between them.
//
// If p starts with '*' only then means match the end of the pattern
//
// If p ends with '*' only then means match the start of the pattern
//
// A '|' allows for multiple patterns
//
// A pattern of "" means always match
func Match(s, pat string) bool {
	pat = strings.TrimSpace(pat)
	if pat == "" {
		return true
	}

	for _, p := range strings.Split(pat, "|") {
		prefix := strings.HasPrefix(p, "*")
		suffix := strings.HasSuffix(p, "*")

		match := false
		switch {
		case prefix && suffix:
			match = strings.Contains(s, p[1:len(p)-1])

		case prefix:
			match = strings.HasSuffix(s, p[1:])

		case suffix:
			match = strings.HasPrefix(s, p[:len(p)-1])

		default:
			match = s == p
		}

		if match {
			return true
		}
	}

	return false
}
