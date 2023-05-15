package unit

import (
	"time"
)

const (
	RFC3339            = time.RFC3339 // for convenience
	RFC3339NoTC        = "2006-01-02T15:04:05"
	WeatherUnderground = "2006-01-02 15:04:05" // mysql format, used by Weather Underground
	Ecowitt            = "2006-01-02+15:04:05" // Common from ecowitt devices
)

// Time formats in order of precedence. If formats start with the same characters then
var timeFormats = []string{
	"2006-01-02+15:04:05", // Common from ecowitt devices
	"2006-01-02 15:04:05", // mysql format, used by Weather Underground
	time.RFC3339,
	"2006-01-02T15:04:05", // RFC3339 without timezone
}

func ParseTime(s string) time.Time {
	for _, tf := range timeFormats {
		if t, err := time.Parse(tf, s); err == nil {
			return t
		}
	}

	return time.Time{}
}
