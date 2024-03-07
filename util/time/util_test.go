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

			year := 2023
			tm := time.Date(year, 1, 1, 0, 13, 14, 0, time.UTC)

			for tm.Year() == year {
				localTime := tm.In(loc)

				got := LocalMidnight(localTime)

				if got.Hour() != 0 {
					t.Errorf("%s got %s for %q", localTime.Format(time.RFC3339), got.Format(time.RFC3339), timeZone)
				}

				tm = tm.Add(time.Hour)
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
