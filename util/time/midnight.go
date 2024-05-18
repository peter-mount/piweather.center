package time

import (
	"time"
)

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
	return midnight(t, 1)
}

// YesterdayMidnight returns the time.Time for midnight of the previous day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func YesterdayMidnight(t time.Time) time.Time {
	return midnight(t, 2)
}

// TomorrowMidnight returns the time.Time for midnight of the following day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func TomorrowMidnight(t time.Time) time.Time {
	// Start from 19:00 in this time zone. This speeds up the search.
	// Any hour later will cause some DST/ST transitions to be missed
	t1 := LocalMidnight(t).Add(19 * time.Hour)
	y0, m0, d0 := t1.Date()

	// Search until we get a date change, that becomes our result
	y1, m1, d1 := y0, m0, d0
	for y0 == y1 && m0 == m1 && d0 == d1 {
		t1 = t1.Add(time.Hour)
		y1, m1, d1 = t1.Date()
	}
	return t1
}

// Locate the appropriate "midnight" from t.
// c the number of midnights to locate, 1=for today, 2=for yesterday.
// This only works for searching back in time.
func midnight(t time.Time, c int) time.Time {
	// move to the start of current hour
	h, m, s := t.Clock()
	t = t.Add(-(time.Duration(m) * time.Minute)).
		Add(-(time.Duration(s) * time.Second)).
		Add(-(time.Duration(t.Nanosecond()) * time.Nanosecond))

	// If h>3 then speed things up by trimming any hours above that from the time
	if h > 6 {
		t = t.Add(-(time.Duration(h-6) * time.Hour))
	}

	for j := 0; j < c; j++ {
		// when j>0 or we are seeking forward, and we are at midnight then move forward
		// a few steps otherwise we will just stop here
		if j > 0 {
			t = t.Add(-time.Hour * time.Duration(20))
		}

		for !IsMidnight(t) {
			t = t.Add(-time.Hour)
		}
	}

	return t
}

// IsMidnight returns true if t represents Midnight on the specific date in t.
// Note: This ignores Nanoseconds within t as we limit ourselves to a second resolution.
func IsMidnight(t time.Time) bool {
	// If we are pedantic, we should add t.Nanosecond()==0 here
	th, tm, ts := t.Clock()
	if !t.IsZero() && tm == 0 && ts == 0 {
		if th == 0 {
			return true
		}

		// comparing DST for t and 1 hour earlier allows us to check for DST changes
		// which happen at midnight - e.g. Cairo
		t1 := t.Add(-time.Hour)
		return t.IsDST() != t1.IsDST()
	}

	return false
}
