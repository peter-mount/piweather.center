package lang

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"time"
)

type Range struct {
	From  time.Time     // Start time
	To    time.Time     // End time
	Every time.Duration // Step duration
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
	return !(r.From.After(t) || r.To.Before(t))
}

func (r Range) Add(t time.Time) time.Time {
	return t.Add(r.Every)
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
	log.Printf("Expand %v %v", min, max)
	r.From = expand(r.From, min, r.From.After)
	r.To = expand(r.To, max, r.To.Before)
	return r
}

func (r Range) Iterator() *TimeIterator {
	if r.Every == 0 {
		r.Every = time.Minute
	}
	if r.Every < time.Second {
		r.Every = time.Second
	}
	return &TimeIterator{
		t: r.From,
		e: r.To,
		s: r.Every,
	}
}

type TimeIterator struct {
	t time.Time
	e time.Time
	s time.Duration
}

func (it *TimeIterator) HasNext() bool {
	return it.t.Before(it.e)
}

func (it *TimeIterator) Next() time.Time {
	t := it.t
	it.t = it.t.Add(it.s)
	return t
}
