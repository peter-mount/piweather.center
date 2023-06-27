package unit

import (
	"strconv"
	"time"
)

const (
	RFC3339            = "2006-01-02T15:04:05Z07:00" // for convenience
	RFC3339Zulu        = "2006-01-02T15:04:05Z"
	RFC3339NoTC        = "2006-01-02T15:04:05"
	WeatherUnderground = "2006-01-02 15:04:05" // mysql format, used by Weather Underground
	Ecowitt            = "2006-01-02+15:04:05" // Common from ecowitt devices
)

// Time formats in order of precedence. If formats start with the same characters then
var timeFormats = []string{
	Ecowitt,            // Common from ecowitt devices
	WeatherUnderground, // mysql format, used by Weather Underground
	RFC3339NoTC,
	RFC3339Zulu,
	RFC3339NoTC,
}

func ParseTime(s string) time.Time {
	// Parse time using one of our formats
	for _, tf := range timeFormats {
		if t, err := time.Parse(tf, s); err == nil {
			return t
		}
	}

	// Unix time
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Unix(i, 0)
	}

	return time.Time{}
}
