package strings

import (
	"regexp"
	"strings"
)

// Expander is a function called by Expand to expand embedded strings in output.
type Expander func(s string) string

var (
	expandRegex = regexp.MustCompile(`\$\{[^}]+}`)
)

func stripExpand(s string) string {
	return strings.TrimPrefix(strings.TrimSuffix(s, "}"), "${")
}

// Expand replaces any instances of "${key}" in a string with the output of an Expander.
// The "${" and "}" wrapping the key are removed before being passed to the Expander.
func Expand(s string, e Expander) string {
	return string(expandRegex.ReplaceAllFunc([]byte(s), func(s []byte) []byte {
		return []byte(e(stripExpand(string(s))))
	}))
}

// Expansions returns a slice of keys that would be expanded if the string was passed to Expand()
func Expansions(s string) []string {
	var r []string
	for _, k := range expandRegex.FindAll([]byte(s), -1) {
		r = append(r, stripExpand(string(k)))
	}
	return r
}
