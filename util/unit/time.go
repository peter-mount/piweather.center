package unit

import (
	time2 "github.com/peter-mount/piweather.center/util/time"
	"strconv"
	"strings"
	"time"
)

const (
	RFC3339            = "2006-01-02T15:04:05Z07:00" // for convenience
	RFC3339Zulu        = "2006-01-02T15:04:05Z"
	RFC3339NoTC        = "2006-01-02T15:04:05"
	WeatherUnderground = "2006-01-02 15:04:05" // mysql format, used by Weather Underground
	Ecowitt            = "2006-01-02+15:04:05" // Common from ecowitt devices
	Minute             = "2006-01-02T15:04Z07:00"
	MinuteZulu         = "2006-01-02T15:04Z"
	MinuteNoTC         = "2006-01-02T15:04"
	Hour               = "2006-01-02T15Z07:00"
	HourZulu           = "2006-01-02T15Z"
	HourNoTC           = "2006-01-02T15"
	Day                = "2006-01-02Z07:00"
	DayZulu            = "2006-01-02Z"
	DayNoTC            = "2006-01-02"
)

// Time formats in order of precedence.
//
// If formats start with the same characters then the longer one must be before the shorter one!
var timeFormats = []string{
	Ecowitt,            // Common from ecowitt devices
	WeatherUnderground, // mysql format, used by Weather Underground
	RFC3339NoTC,
	RFC3339Zulu,
	RFC3339NoTC,
	RFC3339,
	Minute,
	MinuteZulu,
	MinuteNoTC,
	Hour,
	HourZulu,
	HourNoTC,
	Day,
	DayZulu,
	DayNoTC,
}

// ParseTime parses a string to get a time. This will accept the following aliases (case-insensitive):
//
// "now" - the current time
//
// "midnight" - midnight in the local timezone before now
//
// "midnightutc" - midnight in UTC before now
//
// "yesterday" - midnight in the local timezone of the start of yesterday
//
// "yesterdayutc" - midnight in UTC of the start of yesterday
//
// "tomorrow" - midnight in the local timezone of the start of tomorrow
//
// "tomorrowutc" - midnight in UTC of the local timezone of the start of tomorrow
//
// # If not it will attempt to parse the time based on common formats used by various systems:
//
// "2006-01-02T15:04:05Z07:00" RFC3339
//
// "2006-01-02T15:04:05Z" RFC 3339 with just Z for UTC
//
// "2006-01-02T15:04:05" RFC 3339 with no timezone, uses local time zone
//
// "2006-01-02 15:04:05" mysql format, used by Weather Underground
//
// "2006-01-02+15:04:05" used by EcoWitt devices
//
// If none of the above then this will try to parse as an integer for Unix time
// (number of seconds since 1970 Jan 1)
//
// If the string cannot be passed then a zero time.Time is returned
func ParseTime(s string) time.Time {
	return ParseTimeIn(s, time.Local)
}

func ParseTimeIn(s string, loc *time.Location) time.Time {

	switch strings.ToLower(s) {
	case "now":
		return NowIn(loc)

	case "midnight", "today":
		return time2.LocalMidnight(NowIn(loc))

	case "midnightutc", "todayutc":
		return time2.LocalMidnight(NowUTC())

	case "yesterday":
		return time2.LocalMidnight(NowIn(loc).Add(-24 * time.Hour))

	case "yesterdayutc":
		return time2.LocalMidnight(NowUTC().Add(-24 * time.Hour))

	case "tomorrow":
		return time2.LocalMidnight(NowIn(loc).Add(24 * time.Hour))

	case "tomorrowutc":
		return time2.LocalMidnight(NowUTC().Add(24 * time.Hour))

	default:
		// Parse time using one of our formats
		for _, tf := range timeFormats {
			if t, err := time.ParseInLocation(tf, s, loc); err == nil {
				return t
			}
		}

		// Unix time
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return time.Unix(i, 0)
		}

		return time.Time{}
	}
}

func NowIn(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}
