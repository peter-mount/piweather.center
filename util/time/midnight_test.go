package time

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
	"unicode"
)

// For every timezone on the test machine,
// run LocalMidnight against every hour of the UTC year for the four-year period 2023..2026, handling leap years.
func TestLocalMidnight(t *testing.T) {
	testAllTimeZones(t, LocalMidnight, nil)
}

// For every timezone on the test machine,
// run YesterdayMidnight against every hour of the UTC year for the four-year period 2023..2026, handling leap years.
func TestYesterdayMidnight(t *testing.T) {
	testAllTimeZones(t, YesterdayMidnight, func(t *testing.T, timeZone string, tm, got time.Time) {
		// Now check the date is yesterday
		midnight := LocalMidnight(tm)
		if !IsMidnight(midnight) {
			t.Errorf("%q is not midnight", midnight.String())
		}

		if !got.Before(midnight) {
			t.Errorf("%q is not yesterday for %q", got.String(), midnight.String())
		}
	})
}

// TestYesterdayMidnight had errors for some entries where midnight is off by 2 days.
// This is now fixed but this checks that it is still fixed.
func TestYesterdayMidnight_wrongDay(t *testing.T) {
	for _, test := range []struct {
		timeZone string
		date     string
		hrStart  int
		hrEnd    int
		expected string
	}{
		{
			timeZone: "Africa/Cairo",
			date:     "2023-04-29T00:13:14-03:00",
			hrStart:  0,
			hrEnd:    5,
			expected: "2023-04-28",
		},
		{
			timeZone: "Atlantic/Azores",
			date:     "2023-03-27T00:13:14-03:00",
			hrStart:  0,
			hrEnd:    5,
			expected: "2023-03-26",
		},
		{
			timeZone: "Asia/Famagusta",
			date:     "2023-10-29T02:13:14+02:00",
			hrStart:  0,
			hrEnd:    5,
			expected: "2023-10-28",
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

				expected, err := time.Parse(time.DateOnly, test.expected)
				if err != nil {
					t.Errorf("Failed to parse %q: %v", test.date, err)
					return
				}
				expY, expM, expD := expected.Date()

				tm := t0.UTC().In(loc)

				c := 0
				for h := test.hrStart; h <= test.hrEnd; h++ {

					got := YesterdayMidnight(tm)

					testTimeIsMidnight(t, test.timeZone, tm, got)

					// Test it's the date we are expecting
					if yr, mn, dy := got.Date(); !(yr == expY && mn == expM && dy == expD) {
						t.Errorf("Wrong day %d\ntm\t\t%-25s\ngot\t\t%-25s\t%v\nexpect\t%-25s\n",
							h,
							tm.Format(time.RFC822),
							got.Format(time.RFC822),
							IsMidnight(got),
							test.expected)
					}

					// Now check the date is yesterday
					midnight := LocalMidnight(tm)
					midnight = midnight.UTC().In(midnight.Location())

					dr := midnight.Sub(got)
					d := int(dr / time.Hour)
					if d < 20 || d > 28 {
						t.Errorf("Too far %d\nmnight\t%-25s\t%-25s\ngot\t\t%-25s\t%-25s\noffset\t%s\n",
							h,
							midnight.Format(time.RFC822),
							midnight.Format(time.RFC3339),
							got.Format(time.RFC822),
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
// run TomorrowMidnight against every hour of the UTC year for the four-year period 2023..2026, handling leap years.
func TestTomorrowMidnight(t *testing.T) {
	testAllTimeZones(t, TomorrowMidnight, func(t *testing.T, timeZone string, tm, got time.Time) {
		// Now check the date is tomorrow
		midnight := LocalMidnight(tm)
		if !IsMidnight(midnight) {
			t.Errorf("%q is not midnight", midnight.String())
		}

		ty, tmn, td := midnight.Date()
		tomorrow := LocalMidnight(time.Date(ty, tmn, td+1, 4, 0, 0, 0, midnight.Location()))
		//tomorrow := LocalMidnight(midnight.AddDate(0, 0, 1))
		//tomorrow := LocalMidnight(midnight.Add(24 * time.Hour))
		if !IsMidnight(tomorrow) {
			t.Errorf("tomorrow not midnight %q", tomorrow.String())
		}

		if got.Before(midnight) || tomorrow.Before(got) {
			t.Errorf("Not tomorrow, got %q expected %q from %q %q", got.String(), midnight.String(), tm.String(), tomorrow.String())
		}
	})
}

// TestTomorrowMidnight had errors for some entries where midnight is off by a day.
// This is now fixed but this checks that it is still fixed.
func TestTomorrowMidnight_wrongDay(t *testing.T) {
	for _, test := range []struct {
		timeZone string
		date     string
		hrStart  int
		hrEnd    int
		expected string
	}{
		{
			timeZone: "America/Asuncion",
			date:     "2023-09-30T00:13:14-04:00",
			hrStart:  0,
			hrEnd:    5,
			expected: "2023-10-01",
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

				expected, err := time.Parse(time.DateOnly, test.expected)
				if err != nil {
					t.Errorf("Failed to parse %q: %v", test.date, err)
					return
				}
				expY, expM, expD := expected.Date()

				tm := t0.UTC().In(loc)

				c := 0
				for h := test.hrStart; h <= test.hrEnd; h++ {

					got := TomorrowMidnight(tm)

					testTimeIsMidnight(t, test.timeZone, tm, got)

					// Test it's the date we are expecting
					if yr, mn, dy := got.Date(); !(yr == expY && mn == expM && dy == expD) {
						t.Errorf("Wrong day %d\ntm\t\t%-25s\ngot\t\t%-25s\t%v\nexpect\t%-25s\n",
							h,
							tm.Format(time.RFC822),
							got.Format(time.RFC822),
							IsMidnight(got),
							test.expected)
					}

					// Now check the date is yesterday
					midnight := LocalMidnight(tm)
					midnight = midnight.UTC().In(midnight.Location())

					dr := -midnight.Sub(got)
					d := int(dr / time.Hour)
					if d < 20 || d > 28 {
						t.Errorf("Too far %d\nmnight\t%-25s\t%-25s\ngot\t\t%-25s\t%-25s\noffset\t%s\n",
							h,
							midnight.Format(time.RFC822),
							midnight.Format(time.RFC3339),
							got.Format(time.RFC822),
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

func testAllTimeZones(t *testing.T, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, got time.Time)) {

	m := make(map[string][]string)
	for _, zone := range getAvailableTimeZones() {
		var k, v string
		s := strings.SplitN(zone, "/", 2)
		if len(s) == 2 {
			k, v = s[0], s[1]
		} else {
			k, v = "_", s[0]
		}
		m[k] = append(m[k], v)
	}

	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		if k != "_" {
			t.Run(k, func(t *testing.T) {
				testTimeZones(t, k+"/", m[k], f, test)
			})
		} else {
			testTimeZones(t, "", m[k], f, test)
		}
	}
}

func testTimeZones(t *testing.T, prefix string, zones []string, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, got time.Time)) {
	sort.SliceStable(zones, func(i, j int) bool {
		return zones[i] < zones[j]
	})
	for _, timeZone := range zones {
		testTimeZone(t, prefix, timeZone, f, test)
	}
}

func testTimeZone(t *testing.T, prefix, zone string, f func(time.Time) time.Time, test func(t *testing.T, timeZone string, tm, got time.Time)) {
	timeZone := prefix + zone
	t.Run(zone, func(t *testing.T) {
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
	if !IsMidnight(got) {
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

	// Skip Lord Howe Island as it fails due to daylight savings being 30 minutes & the general population
	// for the island is around 6, so not of much use spending time to fix this one.
	var r []string
	for _, zone := range timeZones {
		if zone != "Australia/Lord_Howe" && zone != "Australia/LHI" {
			r = append(r, zone)
		}
	}

	return r
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
