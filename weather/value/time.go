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
