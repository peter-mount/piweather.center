package time

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestXXX(t *testing.T) {
	//timeZone := "Africa/Addis_Ababa"
	//timeZone := "Africa/Cairo"
	timeZone := "Europe/London"
	//timeZone := "Australia/Lord_Howe"

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		t.Errorf("Failed to load %q: %v", timeZone, err)
		return
	}

	for d := 1; d <= 31; d++ {
		t.Run(fmt.Sprintf("%s %02d", timeZone, d),
			func(t *testing.T) {
				t0 := time.Date(2024, 4, d-1, 23, 0, 0, 0, loc).
					Add(time.Hour)
				tm := t0

				c := 0
				for h := 0; tm.Day() == t0.Day(); h++ {

					got := LocalMidnight(tm)

					testTimeIsMidnight(t, timeZone, tm, got)

					//if h == 0 {
					fmt.Printf("%02d/%02d\t%s\t%v\t%s\t%v\n",
						d, h,
						tm.Format(time.RFC822),
						IsMidnight(tm),
						got.Format(time.RFC822),
						IsMidnight(got))
					//}

					tm = tm.Add(time.Hour)
					c++
				}
				fmt.Printf("hours %d\n", c)
			})
	}
}

func TestLocalMidnight2(t *testing.T) {
	testTimeZone(t, "Europe/London", LocalMidnight, nil)
	testTimeZone(t, "Africa/Cairo", LocalMidnight, nil)
}

func TestIsMidnight(t *testing.T) {
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		t.Errorf("Failed to load location %v", err)
	}

	dt := time.Date(2023, 3, 25, 23, 13, 14, 0, loc)
	for h := 0; h < 5; h++ {
		mid := LocalMidnight(dt)
		fmt.Printf("dt %v %v midnight %v %v\n",
			dt.String(),
			IsMidnight(dt),
			mid.String(),
			IsMidnight(mid))

		if mid.Hour() != 0 {
			t.Errorf("Not midnight %v got %v",
				dt.String(),
				mid.String(),
			)
		}

		dt = dt.Add(time.Hour)
	}
}

// For every timezone on the test machine,
// run LocalMidnight against every hour of the UTC year
// for the four-year period 2023..2026, handling leap years.
func TestLocalMidnight(t *testing.T) {
	testAllTimeZones(t, LocalMidnight, nil)
}

// For every timezone on the test machine,
// run YesterdayMidnight against every hour of the UTC year
// for the four-year period 2023..2026, handling leap years.
func TestYesterdayMidnight(t *testing.T) {
	testAllTimeZones(t, YesterdayMidnight, func(t *testing.T, timeZone string, tm, got time.Time) {
		// Now check the date is yesterday
		midnight := LocalMidnight(tm)
		midnight = midnight.UTC().In(midnight.Location())

		if !got.Before(midnight) {
			t.Errorf("%q is not yesterday for %q", got.String(), midnight.String())
		}

		if !IsMidnight(midnight) {
			t.Errorf("%q is not midnight", midnight.String())
		}

		dr := midnight.Sub(got)
		d := int(dr / time.Hour)
		if d < 20 || d > 27 {
			t.Errorf("Not same date, %q midnight %q got %q offset %s",
				tm.String(),
				midnight.String(),
				got.String(),
				dr.String())
		}
	})
}

// TestYesterdayMidnight has errors for some entries where midnight is off by 2 days
func TestYesterdayMidnight_wrongDay(t *testing.T) {
	for _, test := range []struct {
		timeZone string
		date     string
		hrStart  int
		hrEnd    int
	}{
		{
			timeZone: "Africa/Cairo",
			date:     "2023-04-29T00:13:14-03:00",
			hrStart:  0,
			hrEnd:    5,
		},
		{
			timeZone: "Atlantic/Azores",
			date:     "2023-03-27T00:13:14-03:00",
			hrStart:  0,
			hrEnd:    5,
		},
	} {
		t.Run(test.timeZone,
			func(t *testing.T) {

				// Verify timeZone is valid
				loc, err := time.LoadLocation(test.timeZone)
				if err != nil {
					t.Errorf("Failed to load %q: %v", test.timeZone, err)
					return
				}

				// Parse start time
				t0, err := time.Parse(time.RFC3339, test.date)
				if err != nil {
					t.Errorf("Failed to parse %q: %v", test.date, err)
					return
				}

				tm := t0.UTC().In(loc)

				c := 0
				for h := test.hrStart; h <= test.hrEnd; h++ {

					got := YesterdayMidnight(tm)

					testTimeIsMidnight(t, test.timeZone, tm, got)

					fmt.Printf("tm %s\tgot %s\t%v\n",
						tm.Format(time.RFC3339),
						got.Format(time.RFC3339),
						IsMidnight(got))

					// Now check the date is yesterday
					midnight := LocalMidnight(tm)
					midnight = midnight.UTC().In(midnight.Location())

					dr := midnight.Sub(got)
					d := int(dr / time.Hour)
					if d < 20 || d > 27 {
						t.Errorf("Not same date, midnight %s got %s offset %s",
							midnight.Format(time.RFC3339),
							got.Format(time.RFC3339),
							dr.String())
					}

					tm = tm.Add(time.Hour)
					c++
				}
				fmt.Printf("hours %d\n", c)
			})
	}
}

// For every timezone on the test machine,
// run TomorrowMidnight against every hour of the UTC year
// for the four-year period 2023..2026, handling leap years.
func TestTomorrowMidnight(t *testing.T) {
	testAllTimeZones(t, TomorrowMidnight, func(t *testing.T, timeZone string, tm, got time.Time) {
		// Now check the date is tomorrow
		midnight := LocalMidnight(tm)

		tomorrow := LocalMidnight(midnight.AddDate(0, 0, 1))

		if !IsMidnight(tomorrow) {
			t.Errorf("tomorrow not midnight %q", tomorrow.String())
		}

		//mYd := DayId(midnight)
		//
		//// Lookup tomorrow, checking for days which are 23, 24 or 25 hours long
		//tomorrow := midnight
		//for i := 23; i < 26; i++ {
		//	tomorrow = LocalMidnight(midnight.Add(time.Duration(i) * time.Hour))
		//	if DayId(midnight) > mYd {
		//		break
		//	}
		//}

		if !tomorrow.After(tm) {
			t.Errorf("%q %q is not tomorrow for %q", got.String(), tomorrow.String(), tm.String())
		}

		// tomorrow should match what the actual function have returned
		if tomorrow.Unix() != got.Unix() {
			t.Errorf("Not same date, %q tomorrow %q got %q",
				tm.String(),
				tomorrow.String(),
				got.String())
		}
	})
}

func testAllTimeZones(t *testing.T, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, got time.Time)) {
	for _, timeZone := range getAvailableTimeZones() {
		testTimeZone(t, timeZone, f, test)
	}
}

func testTimeZone(t *testing.T, timeZone string, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, got time.Time)) {
	t.Run(timeZone, func(t *testing.T) {
		loc, err := time.LoadLocation(timeZone)
		if err != nil {
			t.Errorf("Failed to load %q: %v", timeZone, err)
			return
		}

		// test across a four-year period where one, 2024, is a leap year.
		// Note: This time span is in UTC not the timezone being tested
		for year := 2023; year <= 2026; year++ {
			tm := time.Date(year, 1, 1, 0, 13, 14, 0, loc)

			for tm.Year() == year {
				got := f(tm)

				testTimeIsMidnight(t, timeZone, tm, got)

				if test != nil {
					test(t, timeZone, tm, got)
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

		default:
			failed = got.Hour() != 0
		}
	} else {
		failed = got.Hour() != 0
	}

	if failed {
		t.Errorf("%s got %s for %q",
			localTime.Format(time.RFC3339),
			got.Format(time.RFC3339),
			timeZone)
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
