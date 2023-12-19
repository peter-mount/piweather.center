package api

import (
	"fmt"
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

func RangeBetween(s, e time.Time) Range {
	if s.After(e) {
		s, e = e, s
	}
	return Range{From: s, To: e}
}

func RangeAt(s time.Time) Range {
	return RangeBetween(s, s)
}

func RangeFrom(s time.Time, d time.Duration) Range {
	return RangeBetween(s, s.Add(d))
}

func (r Range) Contains(t time.Time) bool {
	return !t.Before(r.From) && t.Before(r.To)
}

func (r Range) String() string {
	return fmt.Sprintf("%s to %s every %s",
		r.From.Format(time.RFC3339),
		r.To.Format(time.RFC3339),
		r.Every.String(),
	)
}

func (r Range) IsValid() bool {
	return !(r.From.IsZero() || r.To.IsZero() || r.From.After(r.To))
}

func (r Range) Equals(b Range) bool {
	return r.IsValid() && b.IsValid() && r.From == b.From && r.To == b.To
}

func expand(t time.Time, d time.Duration, f func(time.Time) bool) time.Time {
	nt := t.Add(d)
	if f(nt) {
		return nt
	}
	return t
}

func (r Range) Expand(min, max time.Duration) Range {
	r.From = expand(r.From, min, r.From.After)
	r.To = expand(r.To, max, r.To.Before)
	return r
}

func (r Range) Iterator() *TimeIterator {
	return Iterate(r.From, r.To, r.Every)
}

type TimeIterator struct {
	t time.Time
	e time.Time
	s time.Duration
}

func Iterate(t, e time.Time, s time.Duration) *TimeIterator {
	if t.After(e) {
		t, e = e, t
	}

	if s < 0 {
		s = -s
	}
	if s == 0 {
		s = time.Minute
	}
	if s < time.Second {
		s = time.Second
	}
	return &TimeIterator{t: t, e: e, s: s}
}

func (it *TimeIterator) HasNext() bool {
	return it.t.Before(it.e)
}

func (it *TimeIterator) Next() time.Time {
	t := it.t
	it.t = it.t.Add(it.s)
	return t
}
