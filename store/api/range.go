package api

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"time"
)

type Range struct {
	// Start time
	From time.Time `json:"from"`
	// End time
	To time.Time `json:"to"`
	// Step duration
	Every time.Duration `json:"every"`
}

func (r *Range) write(w *writer) error {
	err := w.time(r.From)
	if err == nil {
		err = w.time(r.To)
	}
	if err == nil {
		err = w.duration(r.Every)
	}
	return err
}

func (r *Range) read(rd *reader) error {
	var err error
	r.From, err = rd.time()
	if err == nil {
		r.To, err = rd.time()
	}
	if err == nil {
		r.Every, err = rd.duration()
	}
	return err
}

// RangeBetween returns a Range between two Times
func RangeBetween(s, e time.Time) Range {
	if s.After(e) {
		s, e = e, s
	}
	return Range{From: s, To: e}
}

// RangeAt returns a Range consisting of a single Time.
func RangeAt(s time.Time) Range {
	return RangeBetween(s, s)
}

// RangeFrom returns a Range from one time and a duration from it.
func RangeFrom(s time.Time, d time.Duration) Range {
	return RangeBetween(s, s.Add(d))
}

// Contains returns true if the Range contains a specific time.
func (r *Range) Contains(t time.Time) bool {
	return r != nil && !t.Before(r.From) && t.Before(r.To)
}

// Duration returns the duration of the Range between the Start and End times
func (r *Range) Duration() time.Duration {
	if r == nil {
		return 0
	}
	return r.To.Sub(r.From)
}

// IsZero returns true if either of the From or To times IsZero
func (r *Range) IsZero() bool {
	return r != nil && (r.From.IsZero() || r.To.IsZero())
}

// Add adds two ranges together so that the new Range will contain both Ranges.
// If they do not intersect, the new Range will contain all time between them.
// If the Range being added is zero this is a no-operation.
// If the source Range is zero then the Range being added will be returned.
func (r *Range) Add(b Range) Range {
	if r == nil || r.IsZero() {
		return b
	}
	if b.IsZero() {
		return *r
	}
	if b.From.Before(r.From) {
		r.From = b.From
	}
	if b.To.After(r.To) {
		r.To = b.To
	}
	return *r
}

func (r *Range) String() string {
	if r == nil {
		return ""
	}

	return fmt.Sprintf("%s to %s every %s",
		r.From.Format(time.RFC3339),
		r.To.Format(time.RFC3339),
		r.Every.String(),
	)
}

func (r *Range) IsValid() bool {
	return r != nil && !(r.From.IsZero() || r.To.IsZero() || r.From.After(r.To))
}

func (r *Range) Equals(b Range) bool {
	return r.IsValid() && b.IsValid() && r.From == b.From && r.To == b.To
}

func expand(t time.Time, d time.Duration, f func(time.Time) bool) time.Time {
	nt := t.Add(d)
	if f(nt) {
		return nt
	}
	return t
}

func (r *Range) Expand(min, max time.Duration) Range {
	r.From = expand(r.From, min, r.From.After)
	r.To = expand(r.To, max, r.To.Before)
	return *r
}

func (r *Range) Iterator() util.Iterator[time.Time] {
	return Iterate(r.From, r.To, r.Every)
}

// Iterate returns an Iterator which will iterate between two time.Time's with the
// specified time.Duration between each step.
//
// This is the same as time.IterateBetween except this will default the step to time.Minute if 0,
// or time.Second if step is less than 1 second.
func Iterate(start, end time.Time, step time.Duration) util.Iterator[time.Time] {
	if start.After(end) {
		start, end = end, start
	}

	step = step.Abs()
	if step == 0 {
		step = time.Minute
	}
	if step < time.Second {
		step = time.Second
	}

	return time2.IterateBetween(start, end, step)
}
