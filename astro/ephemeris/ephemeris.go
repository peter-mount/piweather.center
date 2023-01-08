package ephemeris

import (
	"encoding/xml"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
)

// Ephemeris is a container of results where calculations for an object over a period of time are stored
type Ephemeris struct {
	XMLName xml.Name      `xml:"https://piweather.center/xml/ephemeris ephemeris" json:"-" yaml:"-"` // XML only
	Name    string        `xml:"name,omitempty"`                                                     // Name or title of this Ephemeris. This can be anything.
	Range   *julian.Range `xml:"range"`                                                              // Range of dates this Ephemeris covers
	Meta    Meta          `xml:"meta"`                                                               // Meta data
	Entries []*Entry      `xml:"entries>entry"`                                                      // Entries within the Ephemeris
}

// Meta contains static data used by the Ephemeris
type Meta struct {
	LatLong coord.LatLong `xml:"location"` // Location on Earth this Ephemeris applies to
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

// IsEmpty returns true if the Ephemeris is empty
func (e *Ephemeris) IsEmpty() bool {
	return len(e.Entries) == 0
}
