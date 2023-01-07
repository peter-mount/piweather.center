package ephemeris

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
)

type Ephemeris struct {
	Name    string        `xml:"name"`
	Range   *julian.Range `xml:"range"`
	Meta    Meta          `xml:"meta"`
	Entries []Entry       `xml:"entries>entry"`
}

type Meta struct {
	LatLong coord.LatLong `xml:"location"`
}

type Entry struct {
	Date       julian.Day        `xml:"date,attr"`
	Equatorial *coord.Equatorial `xml:"equatorial,omitempty"`
	RiseSet    *coord.RiseSet    `xml:"riseSet,omitempty"`
}

// Include ensures that the Range includes the specific Day
func (e *Ephemeris) Include(d julian.Day) *Ephemeris {
	e.Range = e.Range.Include(d)
	return e
}

// IncludePeriod includes 2 dates. This is equivalent to calling Include() twice.
func (e *Ephemeris) IncludePeriod(a, b julian.Day) *Ephemeris {
	e.Range = e.Range.IncludePeriod(a, b)
	return e
}

// IncludeRange will include another Range into this one.
// If the range to be included is not Valid this Range is unchanged.
func (e *Ephemeris) IncludeRange(b *julian.Range) *Ephemeris {
	e.Range = e.Range.IncludeRange(b)
	return e
}

// IncludeDays will include an arbitrary number of Day's to the Range.
// If including one or two days only, use Include() or IncludePeriod() instead as
// they will be more efficient.
func (e *Ephemeris) IncludeDays(a ...julian.Day) *Ephemeris {
	e.Range = e.Range.IncludeDays(a...)
	return e
}

// AppendDuration adds a specified amount of time to the end of the Range,
// extending it further into the future
// This has no effect if the Range does not have any current Date(s).
// A negative duration will have no effect on the start but may affect the end if it overlaps.
func (e *Ephemeris) AppendDuration(duration float64) *Ephemeris {
	e.Range = e.Range.AppendDuration(duration)
	return e
}

// PrependDuration adds a specified amount of time to the start of the Range,
// extending it further into the past.
// This has no effect if the Range does not have any current Date(s).
// A negative duration will have no effect on the end but may affect the start if it overlaps.
func (e *Ephemeris) PrependDuration(duration float64) *Ephemeris {
	e.Range = e.Range.PrependDuration(duration)
	return e
}

func (e *Ephemeris) Append(entry Entry) {
	e.Entries = append(e.Entries, entry)
}

func (e *Ephemeris) IsEmpty() bool {
	return len(e.Entries) == 0
}
