package time

import (
	"math"
	"time"
)

// Period represents a period of time
type Period struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

// PeriodOf creates a period based on two Time's
func PeriodOf(start, end time.Time) Period {
	start, end = NormalizeTime(start, end)
	return Period{
		start:    start,
		end:      end,
		duration: end.Sub(start),
	}
}

// Start returns the start of a period.
func (r Period) Start() time.Time { return r.start }

// End returns the end of a period.
func (r Period) End() time.Time { return r.end }

func (r Period) Range() (time.Time, time.Time) {
	return r.start, r.end
}

// Duration returns the duration of the period.
func (r Period) Duration() time.Duration { return r.duration }

// DurationMinutes is the duration of the period in minutes.
// This is useful when dealing with charts by using the minute as the
// real unit of an axis.
func (r Period) DurationMinutes() float64 { return r.duration.Minutes() }

// IsZero returns true if the Period is undefined.
// Specifically either its start or end values are zero.
func (r Period) IsZero() bool {
	return r.start.IsZero() || r.end.IsZero()
}

// Include returns a union of this period and another one.
// If the included period is zero then the period is unchanged.
// If the period is zero then the included period is returned.
func (r Period) Include(b Period) Period {
	// Handle if either span is zero
	if b.IsZero() {
		return r
	}
	if r.IsZero() {
		return b
	}

	start, end := r.start, r.end
	if b.start.Before(start) {
		start = b.start
	}
	if b.end.After(end) {
		end = b.end
	}
	return PeriodOf(start, end)
}

// Add includes a Time to the Period, expanding the period as required.
// If the Time is zero then the period is unchanged.
// If the Period is zero then a new Period is returned consisting of just
// the instant represented by Time.
func (r Period) Add(t time.Time) Period {
	// do not include t if it's zero
	if t.IsZero() {
		return r
	}

	// If r is zero then new span just of t
	if r.IsZero() {
		return PeriodOf(t, t)
	}

	start, end := r.start, r.end
	if t.Before(start) {
		start = t
	}
	if t.After(end) {
		end = t
	}
	return PeriodOf(start, end)
}

// Contains returns true if the Period contains a Time.
// If either the Period or Time are zero then this always returns false.
func (r Period) Contains(t time.Time) bool {
	// If either is zero then always false
	if r.IsZero() || t.IsZero() {
		return false
	}

	return !(t.Before(r.start) || t.After(r.end))
}

// MinutesFromStart returns the duration in minutes since the Period's
// start for a Time. This is useful when plotting a value within a Period.
// If either the Period or Time are zero then this returns NaN
func (r Period) MinutesFromStart(t time.Time) float64 {
	if t.IsZero() || r.IsZero() {
		return math.NaN()
	}
	return -r.start.Sub(t).Minutes()
}
