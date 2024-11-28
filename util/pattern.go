package util

import (
	"strings"
)

const (
	PatternEquals = iota
	PatternPrefix
	PatternSuffix
	PatternContains
	PatternAll
	PatternNone
)

type PatternType uint8

func (m PatternType) Match(s, p string) bool {
	switch m {
	case PatternAll:
		return true
	case PatternPrefix:
		// May look weird but *pattern means suffix=pattern
		return strings.HasSuffix(s, p)
	case PatternSuffix:
		// May look weird but pattern* means prefix=pattern
		return strings.HasPrefix(s, p)
	case PatternContains:
		return strings.Contains(s, p)
	case PatternEquals:
		return s == p
	case PatternNone:
		return false
	default:
		return false
	}
}

func ParsePatternType(s string) (PatternType, string) {
	s = strings.TrimSpace(s)
	prefix := strings.HasPrefix(s, "*")
	suffix := strings.HasSuffix(s, "*")

	var t PatternType
	switch {
	case s == "":
		t = PatternNone
	case s == "*", s == "**":
		t = PatternAll
		s = ""
	case prefix && suffix:
		t = PatternContains
		s = strings.TrimPrefix(strings.TrimSuffix(s, "*"), "*")
	case prefix:
		t = PatternPrefix
		s = strings.TrimPrefix(s, "*")
	case suffix:
		t = PatternSuffix
		s = strings.TrimSuffix(s, "*")
	default:
		t = PatternEquals
	}

	return t, s
}
