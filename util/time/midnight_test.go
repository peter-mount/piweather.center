package time

import (
	"fmt"
	"testing"
	"time"
)

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
