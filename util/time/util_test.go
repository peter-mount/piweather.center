package time

import (
	"os"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestLocalMidnight2(t *testing.T) {
	testTimeZone(t, "Europe/London", LocalMidnight, func(t *testing.T, timeZone string, tm, localTime, got time.Time) {
		tmDayId := DayId(tm.In(localTime.Location()))
		gotDayId := DayId(got)
		if tmDayId != gotDayId {
			t.Errorf("Not same day %d %d %s %s",
				tmDayId, gotDayId,
				tm.String(),
				got.String(),
			)
		}
	})
}

// For every timezone on the test machine, run local midnight against every hour
// of the UTC year.
//
// This then tests things like Daylight Savings etc
func TestLocalMidnight(t *testing.T) {
	testAllTimeZones(t, LocalMidnight, func(t *testing.T, timeZone string, tm, localTime, got time.Time) {
		// Check that midnight is on the same day as localTime.
		// Note: tm.In() here is so that it handles DST
		tmDayId := DayId(tm.In(got.Location()))
		gotDayId := DayId(got)
		if tmDayId != gotDayId {
			t.Errorf("Not same day %d %d %s %s",
				tmDayId, gotDayId,
				tm.String(),
				got.String(),
			)
		}
	})
}

func TestYesterdayMidnight(t *testing.T) {
	testAllTimeZones(t, YesterdayMidnight, func(t *testing.T, timeZone string, tm, localTime, got time.Time) {
		// Now check the date is yesterday
		midnight := LocalMidnight(localTime)
		yesterday := LocalMidnight(midnight.Add(-time.Hour))

		if !yesterday.Before(midnight) {
			t.Errorf("%q is not yesterday for %q", got.String(), midnight.String())
		}

		// yesterday should match what the actual function have us
		if !yesterday.Equal(got) {
			t.Errorf("Not same date, %q yesterday %q got %q",
				tm.String(),
				got.String(),
				yesterday.String())
		}
	})
}

func TestTomorrowMidnight(t *testing.T) {
	testAllTimeZones(t, TomorrowMidnight, func(t *testing.T, timeZone string, tm, localTime, got time.Time) {
		// Now check the date is tomorrow
		midnight := LocalMidnight(localTime)
		mYd := DayId(midnight)

		// Lookup tomorrow, checking for days which are 23, 24 or 25 hours long
		tomorrow := midnight
		for i := 23; i < 26; i++ {
			tomorrow = LocalMidnight(midnight.Add(time.Duration(i) * time.Hour))
			if DayId(midnight) > mYd {
				break
			}
		}

		if !tomorrow.After(localTime) {
			t.Errorf("%q %q is not tomorrow for %q", tm.String(), got.String(), localTime.String())
		}

		// tomorrow should match what the actual function have us
		if !tomorrow.Equal(got) {
			t.Errorf("Not same date, %q tomorrow %q got %q",
				tm.String(),
				got.String(),
				tomorrow.String())
		}
	})
}

func testAllTimeZones(t *testing.T, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, localTime, got time.Time)) {
	for _, timeZone := range getAvailableTimeZones() {
		testTimeZone(t, timeZone, f, test)
	}
}

func testTimeZone(t *testing.T, timeZone string, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, localTime, got time.Time)) {
	t.Run(timeZone, func(t *testing.T) {
		loc, err := time.LoadLocation(timeZone)
		if err != nil {
			t.Errorf("Failed to load %q: %v", timeZone, err)
			return
		}

		// test across a four-year period where one, 2024, is a leap year.
		// Note: This time span is in UTC not the timezone being tested
		for year := 2023; year <= 2026; year++ {
			tm := time.Date(year, 1, 1, 0, 13, 14, 0, time.UTC)

			for tm.Year() == year {
				localTime := tm.In(loc)

				got := f(localTime)

				testTimeIsMidnight(t, timeZone, localTime, got)

				if test != nil {
					test(t, timeZone, tm, localTime, got)
				}

				tm = tm.Add(time.Hour)
			}

		}
	})
}

// testTimeIsMidnight tests that got is pointing to Midnight.
// This accounts for some Time Zones where when DST occurs and there is no Midnight when the DST transition occurs.
func testTimeIsMidnight(t *testing.T, timeZone string, localTime, got time.Time) {
	// We would expect midnight to occur at 00:00:00
	failed := !IsMidnight(got)

	// If it's not 0 then check special cases
	if failed {
		switch timeZone {
		// These Time Zones switch to DST at midnight,
		// so the start of this single local day is 01:00 and not 00:00
		// e.g. 23:59:59 is followed by 01:00:00
		case "Africa/Cairo",
			"America/Asuncion",
			"America/Havana",
			"America/Santiago",
			"America/Scoresbysund",
			"Asia/Beirut",
			"Atlantic/Azores",
			"Chile/Continental",
			"Cuba",
			"Egypt":
			failed = got.Hour() != 1
		}
	}

	//if !failed {
	//	failed = !IsMidnight(got)
	//}

	if failed {
		t.Errorf("%s got %s for %q", localTime.Format(time.RFC3339), got.Format(time.RFC3339), timeZone)
	}

}

func getAvailableTimeZones() []string {
	var timeZones []string
	for _, zd := range []string{
		// Update path according to your OS
		"/usr/share/zoneinfo/",
		"/usr/share/lib/zoneinfo/",
		"/usr/lib/locale/TZ/",
	} {
		timeZones = walkTzDir(zd, timeZones)

		for idx, zone := range timeZones {
			timeZones[idx] = strings.ReplaceAll(zone, zd+"/", "")
		}
	}
	return timeZones
}

func walkTzDir(path string, zones []string) []string {
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		return zones
	}

	isAlpha := func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return true
	}

	for _, info := range fileInfos {
		if info.Name() != strings.ToUpper(info.Name()[:1])+info.Name()[1:] {
			continue
		}

		if !isAlpha(info.Name()[:1]) {
			continue
		}

		newPath := path + "/" + info.Name()

		if info.IsDir() {
			zones = walkTzDir(newPath, zones)
		} else {
			zones = append(zones, newPath)
		}
	}

	return zones
}
