package api

import (
	time2 "github.com/peter-mount/piweather.center/util/time"
	"testing"
	"time"
)

// testTimeIteratorCount counts the number of entries within an iterator and fails if it isn't the expected number
func testTimeIteratorCount(t *testing.T, name string, it *TimeIterator, expected int) {
	count := 0
	for it.HasNext() {
		_ = it.Next()
		count++
	}
	if count != expected {
		t.Errorf("%s returned %d entries, expected %d", name, count, expected)
	}
}

func TestRangeBetween(t *testing.T) {
	tests := []struct {
		name     string
		start    string
		end      string
		every    string
		expected int
	}{
		{name: "h", start: "2024-11-01 06:00:00", end: "2024-11-01T07:00:00", every: "1m", expected: 60},
		{name: "24h July", start: "2024-07-01", end: "2024-08-01", every: "24h", expected: 31},
		{name: "24h August", start: "2024-08-01", end: "2024-09-01", every: "24h", expected: 31},
		{name: "24h September", start: "2024-09-01", end: "2024-10-01", every: "24h", expected: 30},
		{name: "24h October", start: "2024-10-01", end: "2024-11-01", every: "24h", expected: 31},
		{name: "24h November", start: "2024-11-01", end: "2024-12-01", every: "24h", expected: 30},
		{name: "24h December", start: "2024-12-01", end: "2025-01-01", every: "24h", expected: 31},
		{name: "24h January", start: "2025-01-01", end: "2025-02-01", every: "24h", expected: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			every, err := time.ParseDuration(tt.every)
			if err != nil {
				t.Fatal(err)
			}

			r := RangeBetween(time2.ParseTimeIn(tt.start, time.UTC), time2.ParseTimeIn(tt.end, time.UTC))
			r.Every = every
			testTimeIteratorCount(t, "RangeBetween().Iterator()", r.Iterator(), tt.expected)
		})
	}
}

func TestRangeFrom(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		duration string
		every    string
		expected int
	}{
		{name: "1h", from: "2024-11-01T05:00:00Z", duration: "1h", every: "1m", expected: 60},
		{name: "24h", from: "2024-11-01", duration: "24h", every: "1h", expected: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur, err := time.ParseDuration(tt.duration)
			if err != nil {
				t.Fatal(err)
			}

			every, err := time.ParseDuration(tt.every)
			if err != nil {
				t.Fatal(err)
			}

			r := RangeFrom(time2.ParseTimeIn(tt.from, time.UTC), dur)
			r.Every = every
			testTimeIteratorCount(t, "RangeFrom().Iterator()", r.Iterator(), tt.expected)
		})
	}
}

func TestIterate(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		step     time.Duration
		expected int
	}{
		{
			name:     "1h",
			start:    time2.ParseTimeIn("2024-11-01T05:00:00Z", time.UTC),
			end:      time2.ParseTimeIn("2024-11-01T06:00:00Z", time.UTC),
			step:     time.Minute,
			expected: 60,
		},
		{
			name:     "24h by hr",
			start:    time2.ParseTimeIn("2024-11-01", time.UTC),
			end:      time2.ParseTimeIn("2024-11-02", time.UTC),
			step:     time.Hour,
			expected: 24, // 1 day with data every hour
		},
		{
			name:     "24h by mn",
			start:    time2.ParseTimeIn("2024-11-01", time.UTC),
			end:      time2.ParseTimeIn("2024-11-02", time.UTC),
			step:     time.Minute,
			expected: 24 * 60, // 1 day with data every minute
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTimeIteratorCount(t, "Iterate()", Iterate(tt.start, tt.end, tt.step), tt.expected)
		})
	}
}
