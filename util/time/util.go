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
// Note: this is NOT as simple as Truncate(24 * time.Hour) as that is only valid in UTC.
//
// So what we need to do is first do the Truncate, then adjust the result by the source
// Time's TimeZone offset, which will give us midnight based on the source Time.
//
// This works for most days, except for those where a switch between local Standard
// and Daylight Savings occurs. When that happens the result is not midnight but either
// 01:00 the same day or 23:00 the previous day - because the truncate has presumed
// that the day has 24 hours when in actual fact those local days are
// 23 hours long (Standard to Daylight Savings - e.g. "Spring Forward") or
// 25 hours long (Daylight Savings to Standard - e.g. "Fall back").
//
// So when this happens we then adjust the time again to account for the difference.
//
// TODO: This works for most Time Zones, except for 10 where it's still off.
func LocalMidnight(t time.Time) time.Time {
	// First truncate the time to 24 hours and add the time zone offset
	_, off := t.Zone()
	midnight := t.
		Truncate(24 * time.Hour).
		Add(-time.Duration(off) * time.Second) //.

	// If hour is still not zero then we have a Standard/Day-Light-Saving change
	// on this day so adjust the time accordingly, so if 1 then -1hour, 23 then add 1 hour
	if h := midnight.Hour(); h != 0 {
		if h >= 12 {
			h = h - 24
		}
		midnight = midnight.Add(time.Duration(-h) * time.Hour)
	}

	return midnight
}
