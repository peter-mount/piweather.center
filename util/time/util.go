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
	// Reduce the time to midnight by removing the time component via subtraction.
	// Truncate(time.Hour) works for most time zones, except for those with fractional
	// hours in their zone offsets which causes the minutes not to be 0.
	//
	// Also DO NOT Truncate(24*time.Hour) here as not all days are 24 hours long!
	if !t.IsZero() {
		// A loop is necessary for some Australian Time Zones for some weird reason
		// where on the first subtraction when they switch from DST to ST they still
		// have some minutes remaining
		//for i := 0; i < 2 && !IsMidnight(t); i++ {
		t = t.Add(-(time.Duration(t.Hour()) * time.Hour)).
			Add(-(time.Duration(t.Minute()) * time.Minute)).
			Add(-(time.Duration(t.Second()) * time.Second))

		// If hour is still not zero then we have a Standard/Day-Light-Saving change
		// on this day so adjust the time accordingly, so if 1 then -1hour, 23 then add 1 hour
		/*if h := t.Hour(); h >= 12 {
			loc := t.Location()
			tz := t.UTC().Add(time.Hour)
			t = tz.In(loc)
			fmt.Printf("tz %s t %s %s\n",
				tz.Location().String(),
				t.Location().String(),
				t.Format(time.RFC3339))
		}*/
		if h := t.Hour(); h != 0 {
			if h >= 12 {
				h = h - 24
				t = t.Add(time.Duration(-h) * time.Hour)
			}
		}
		//}
	}
	return t
}

func IsMidnight(t time.Time) bool {
	return !t.IsZero() && t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0
}

// YesterdayMidnight returns the time.Time for midnight of the previous day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func YesterdayMidnight(t time.Time) time.Time {
	return LocalMidnight(LocalMidnight(t).Add(-time.Hour))
}

// TomorrowMidnight returns the time.Time for midnight of the following day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func TomorrowMidnight(t time.Time) time.Time {
	midnight := LocalMidnight(t)
	dayId := DayId(midnight)

	// Check tomorrow allowing for days of 23, 24 or 25 hours in length
	tomorrow := midnight
	for h := 23; h < 26; h++ {
		tomorrow = LocalMidnight(midnight.Add(time.Duration(h) * time.Hour))
		if DayId(tomorrow) > dayId {
			break
		}
	}
	return tomorrow
}

// DayId returns a unique id for a specific date.
// It's equivalent to (Year*366) + YearDay. 366 is used to allow for leap years,
// so for non-leap years there will be a gap in the sequence between 31-Dec and 1-Jan.
func DayId(t time.Time) int {
	return (t.Year() * 366) + t.YearDay()
}
