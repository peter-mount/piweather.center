package util

import "time"

// TimeBetween returns true if start <= t <= end
func TimeBetween(t, start, end time.Time) bool {
	return !(t.Before(start) || t.After(end))
}

// NormalizeTime ensures that time a is before b
func NormalizeTime(a, b time.Time) (time.Time, time.Time) {
	if a.After(b) {
		return b, a
	}
	return a, b
}
