package julian

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
)

type Range struct {
	valid bool // Range is valid. Always use Valid() and not read this directly
	start Day
	end   Day
}

// Valid true if the Range is valid, e.g. it's not nil and contains a Day
func (r *Range) Valid() bool { return r != nil && r.valid }

// Start Day in range.
// Returns 0 if the Range is empty or nil.
func (r *Range) Start() Day {
	if r == nil {
		return 0
	}
	return r.start
}

// End Day in range
// Returns 0 if the Range is empty or nil.
func (r *Range) End() Day {
	if r == nil {
		return 0
	}
	return r.end
}

// Duration returns the duration of the Range in days.
// Returns 0 if the Range is empty or nil.
func (r *Range) Duration() float64 {
	if r == nil {
		return 0
	}
	return float64(r.End() - r.Start())
}

// Include ensures that the Range includes the specific Day
func (r *Range) Include(d Day) *Range {
	return r.create().include(d)
}

// create ensures a Range exists, allows for r being nil when adding
func (r *Range) create() *Range {
	if r == nil {
		return &Range{}
	}
	return r
}

// include adds a Day to a Range. r cannot be nil here
func (r *Range) include(d Day) *Range {
	if r.valid {
		if d.Before(r.start) {
			r.start = d
		}

		if d.After(r.end) {
			r.end = d
		}
	} else {
		// Initial date
		r.start = d
		r.end = d
		r.valid = true
	}
	return r
}

// IncludePeriod includes 2 dates. This is equivalent to calling Include() twice.
func (r *Range) IncludePeriod(a, b Day) *Range { return r.create().include(a).include(b) }

// IncludeRange will include another Range into this one.
// If the range to be included is not Valid this Range is unchanged.
func (r *Range) IncludeRange(b *Range) *Range {
	if b.Valid() {
		return r.IncludePeriod(b.Start(), b.End())
	}
	return r
}

// IncludeDays will include an arbitrary number of Day's to the Range.
// If including one or two days only, use Include() or IncludePeriod() instead as
// they will be more efficient.
func (r *Range) IncludeDays(a ...Day) *Range {
	nr := r.create()
	for _, d := range a {
		nr = nr.include(d)
	}
	return nr
}

// AppendDuration adds a specified amount of time to the end of the Range,
// extending it further into the future
// This has no effect if the Range does not have any current Date(s).
// A negative duration will have no effect on the start but may affect the end if it overlaps.
func (r *Range) AppendDuration(duration float64) *Range {
	if r.Valid() {
		return r.Include(r.End().Add(duration))
	}
	return r
}

// PrependDuration adds a specified amount of time to the start of the Range,
// extending it further into the past.
// This has no effect if the Range does not have any current Date(s).
// A negative duration will have no effect on the end but may affect the start if it overlaps.
func (r *Range) PrependDuration(duration float64) *Range {
	if r.Valid() {
		return r.Include(r.Start().Add(-duration))
	}
	return r
}

// Iterator returns an Iterator that will iterate across the Range with a specific step size in days.
func (r *Range) Iterator(step float64) *Iterator {
	if !r.Valid() {
		// If invalid then return an empty Iterator
		return nil
	}

	return Iterate(r.Start(), r.End(), step)
}

// ForEach will call a function across the range using step days as the interval for each call.
// It's the same as calling Iterator(step).ForEach()
func (r *Range) ForEach(step float64, f IteratorHandler) error { return r.Iterator(step).ForEach(f) }

// Slice will return a slice of all Day's in the range with the specified step interval.
// It's the same as calling Iterator(step).Slice()
func (r *Range) Slice(step float64) []Day {
	return r.Iterator(step).Slice()
}

func (r *Range) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	b := util.NewXmlBuilder(encoder, start)

	if r.Valid() {
		b.AddAttribute(xml.Name{Local: "start"}, r.start.String()).
			AddAttribute(xml.Name{Local: "end"}, r.end.String())
	} else {
		b.AddBoolAttribute(xml.Name{Local: "empty"}, true)
	}

	return b.Build()
}
