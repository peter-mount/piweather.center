package time

import (
	"testing"
	"time"
)

// Simply test that, when we request local midnight for a specific time zone
// we get that midnight and not that for UTC
func TestLocalMidnight(t *testing.T) {
	tests := []struct {
		name  string
		month time.Month
		want  int
	}{
		{name: "Europe/London", month: time.February},
		{name: "Europe/London", month: time.June},
		{name: "US/Eastern", month: time.February},
		{name: "US/Eastern", month: time.June},
	}
	for _, tt := range tests {
		t.Run(tt.name+" "+tt.month.String(), func(t *testing.T) {
			tt.want = 1
			loc, err := time.LoadLocation(tt.name)
			if err != nil {
				t.Errorf("Failed to load %q: %v", tt.name, err)
				return
			}

			tm := time.Date(2023, tt.month, 1, 12, 13, 14, 0, loc)
			got := LocalMidnight(tm)

			if got.Hour() != 0 {
				t.Errorf("LocalMidnight() = %v, want %v", got, 0)
			}
		})
	}
}
