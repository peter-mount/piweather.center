package value

import (
	"github.com/soniakeys/meeus/v3/globe"
	"time"
)

type Time interface {
	// Time of some event
	Time() time.Time
	// SetTime sets the time, used so a single instance can be used
	// in iterators for the same location
	SetTime(time.Time) Time
	// Add a time.Duration to this Time
	Add(d time.Duration)
	// Location on Earth's surface
	Location() *globe.Coord
	// Altitude at Location()
	Altitude() float64
	// Clone this Time.
	Clone() Time
	// ForEach will call a function for each step for duration time
	ForEach(step, duration time.Duration, f func(Time) error) error
}

// PlainTime is a Time with no positional component
func PlainTime(t time.Time) Time {
	return BasicTime(t, nil, 0.0)
}

// BasicTime returns a time with the static coordinates
func BasicTime(t time.Time, loc *globe.Coord, alt float64) Time {
	return &basicTime{t: t, loc: loc, alt: alt}
}

type basicTime struct {
	t   time.Time
	loc *globe.Coord
	alt float64
}

func (b *basicTime) Time() time.Time { return b.t }

func (b *basicTime) SetTime(t time.Time) Time {
	b.t = t
	return b
}

func (b *basicTime) Add(d time.Duration) {
	b.SetTime(b.Time().Add(d))
}

func (b *basicTime) Location() *globe.Coord { return b.loc }

func (b *basicTime) Altitude() float64 { return b.alt }

func (b *basicTime) Clone() Time { return &timeWrapper{t: b.t, p: b} }

func (b *basicTime) ForEach(step, duration time.Duration, f func(Time) error) error {
	return forEachTime(b, step, duration, f)
}

type timeWrapper struct {
	t time.Time
	p Time
}

func (b *timeWrapper) Time() time.Time { return b.t }

func (b *timeWrapper) SetTime(t time.Time) Time {
	b.t = t
	return b
}

func (b *timeWrapper) Add(d time.Duration) {
	b.SetTime(b.Time().Add(d))
}

func (b *timeWrapper) Location() *globe.Coord { return b.p.Location() }

func (b *timeWrapper) Altitude() float64 { return b.p.Altitude() }

func (b *timeWrapper) Clone() Time { return &timeWrapper{t: b.t, p: b.p} }

func (b *timeWrapper) ForEach(step, duration time.Duration, f func(Time) error) error {
	return forEachTime(b, step, duration, f)
}

func forEachTime(t Time, step, duration time.Duration, f func(Time) error) error {
	// Preserve t at end
	ot := t.Time()
	defer func() {
		t.SetTime(ot)
	}()

	end := t.Time().Add(duration)

	for t.Time().Before(end) {
		if err := f(t); err != nil {
			return err
		}
		t.Add(step)
	}

	return nil
}
