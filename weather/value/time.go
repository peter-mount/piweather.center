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
	SetTime(time.Time)
	// Location on Earth's surface
	Location() *globe.Coord
	// Altitude at Location()
	Altitude() float64
	// Clone this Time.
	Clone() Time
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

func (b *basicTime) SetTime(t time.Time) { b.t = t }

func (b *basicTime) Location() *globe.Coord { return b.loc }

func (b *basicTime) Altitude() float64 { return b.alt }

func (b *basicTime) Clone() Time { return &timeWrapper{t: b.t, p: b} }

type timeWrapper struct {
	t time.Time
	p Time
}

func (b *timeWrapper) Time() time.Time { return b.t }

func (b *timeWrapper) SetTime(t time.Time) { b.t = t }

func (b *timeWrapper) Location() *globe.Coord { return b.p.Location() }

func (b *timeWrapper) Altitude() float64 { return b.p.Altitude() }

func (b *timeWrapper) Clone() Time { return b.p.Clone() }
