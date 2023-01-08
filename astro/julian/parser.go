package julian

import (
	"fmt"
	"strconv"
	"time"
)

// The time formats in order of precedence.
// The first one that matches is used.
var timeFormats = []string{
	time.RFC3339,
	"2006-01-02T15:04:05", // RFC3339 without time zone, UTC enforced
	"2006-01-02T15:04",    // RFC3339 without time zone, UTC enforced, seconds=0
	"2006-01-02",          // ISO8601 date
}

func Parse(s string) (Day, error) {
	// Handle decimal representation
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return Day(f), nil
	}

	for _, f := range timeFormats {
		if t, err := time.Parse(f, s); err == nil {
			return FromTime(t), nil
		}
	}

	return 0, fmt.Errorf("unable to parse Julian Day %q", s)
}
