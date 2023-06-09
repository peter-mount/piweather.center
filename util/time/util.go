package time

import (
	"fmt"
	"strings"
	"time"
)

// Between returns true if start <= t <= end
func Between(t, start, end time.Time) bool {
	return !(t.Before(start) || t.After(end))
}

// NormalizeTime ensures that time a is before b
func NormalizeTime(a, b time.Time) (time.Time, time.Time) {
	if a.After(b) {
		return b, a
	}
	return a, b
}

// Zone returns the timezone of a time.Time.
//
// For example, if BST then this returns "BST (UTC+1)"
// If in UTC then only returns "UTC". If GMT then returns "GMT" as that's UTC+0
func Zone(t time.Time) string {
	zone, offset := t.Zone()
	ts := strings.TrimSuffix(
		strings.TrimSuffix(fmt.Sprintf("%.2f", float64(offset)/3600.0), "0"),
		".0")
	if ts == "0" {
		ts = ""
	} else if offset < 0 {
		ts = " (UTC-" + ts + ")"
	} else {
		ts = " (UTC+" + ts + ")"
	}
	return zone + ts
}

// LocalMidnight returns the time of midnight in the local time zone before
// the provided time.
//
// Note: truncate to hour then subtract hours to get the start.
// It might look weird when you could truncate to day, but that truncate
// seems to set it to 0h UTC, so if we are in BST (UTC+1) then the day
// starts at 0100 and not 0000 midnight.
//
// TODO check this works for other timezones
func LocalMidnight(t time.Time) time.Time {
	today := t.Truncate(time.Hour)
	return today.Add(time.Hour * time.Duration(-today.Hour()))
}
