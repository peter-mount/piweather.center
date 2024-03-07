package time

import (
	"os"
	"strings"
	"testing"
	"time"
	"unicode"
)

// For every timezone on the test machine, run local midnight against every hour
// of the UTC year.
//
// This then tests things like Daylight Savings etc
func TestLocalMidnight(t *testing.T) {
	for _, timeZone := range getAvailableTimeZones() {
		t.Run(timeZone, func(t *testing.T) {
			loc, err := time.LoadLocation(timeZone)
			if err != nil {
				t.Errorf("Failed to load %q: %v", timeZone, err)
				return
			}

			// test across a four-year period where one, 2024, is a leap year
			for year := 2023; year <= 2026; year++ {
				tm := time.Date(year, 1, 1, 0, 13, 14, 0, time.UTC)

				for tm.Year() == year {
					localTime := tm.In(loc)

					got := LocalMidnight(localTime)

					// We would expect midnight to occur at 00:00:00
					hour := got.Hour()
					failed := hour != 0

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
							failed = hour != 1
						}
					}

					if failed {
						t.Errorf("%s got %s for %q", localTime.Format(time.RFC3339), got.Format(time.RFC3339), timeZone)
					}

					tm = tm.Add(time.Hour)
				}

			}
		})
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
