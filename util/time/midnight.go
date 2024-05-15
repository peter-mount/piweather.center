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
	return midnight(t, 1, -time.Hour)
}

// YesterdayMidnight returns the time.Time for midnight of the previous day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func YesterdayMidnight(t time.Time) time.Time {
	return midnight(t, 2, -time.Hour)
}

// TomorrowMidnight returns the time.Time for midnight of the following day from t.
// This accounts for any changes within t's TimeZone, e.g. Daylight Saving
func TomorrowMidnight(t time.Time) time.Time {
	return midnight(t, 1, time.Hour)
}

func midnight(t time.Time, c int, offset time.Duration) time.Time {
	if c < 1 || offset.Abs() < time.Hour {
		panic("Invalid midnight() call")
	}

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
		if j > 0 || (offset > 0 && IsMidnight(t)) {
			t = t.Add(offset * time.Duration(20))
		}

		for !IsMidnight(t) {
			t = t.Add(offset)
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
		// Check to see if we are on the start date of the current zone, and if so
		// then return true if the time is the same as the start time.
		// If not then fall through to the usual check for hour 0.
		// This handles the odd time zones where, when they switch to DST Midnight is 01:00 and not 00:00
		zs, _ := t.ZoneBounds()
		if !zs.IsZero() {
			zyr, zmn, zdy := zs.Date()

			tyr, tmn, tdy := t.Date()

			if zyr == tyr && zmn == tmn && zdy == tdy {
				zh := zs.Hour()
				if zh == 0 {
					return zh == th
				}
			}
		}

		return th == 0
	}

	return false
}
