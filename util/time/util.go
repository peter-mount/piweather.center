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
// Note: this is NOT as simple as Truncate(24 * time.Hour) as not all days are 24 hours long.
// Where a Time Zone uses daylight saving, then the days they switch can be either
// 23 hours long (Standard to Daylight Savings - e.g. "Spring Forward") or
// 25 hours long (Daylight Savings to Standard - e.g. "Fall back").
//
// Note: Local Midnight is usually 00:00:00 but for some TimeZones, on the day they switch to Daylight Savings
// the do so at midnight, so this returns 01:00:00 for those TimeZones for that specific day.
func LocalMidnight(t time.Time) time.Time {
	// Truncate to the current hour, then subtract the remaining hours.
	// Do not truncate to 24*time.Hour here as not all days are 24 hours long!
	midnight := t.Truncate(time.Hour)
	midnight = midnight.Add(-time.Duration(midnight.Hour()) * time.Hour)

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
