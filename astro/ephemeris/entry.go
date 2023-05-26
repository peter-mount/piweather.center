package ephemeris

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"sort"
)

// Entry represents the calculated value for an object
type Entry struct {
	Date       julian.Day        `xml:"date,attr"`            // Date of this entry, used in sorting and display
	Ord        int               `xml:"ord,attr,omitempty"`   // Ordinal, default 0, used in sorting
	Name       string            `xml:"name,attr,omitempty"`  // Name of object in entry, default "", used in sorting and display
	Equatorial *coord.Equatorial `xml:"equatorial,omitempty"` // Equatorial coordinates
	RiseSet    *coord.RiseSet    `xml:"riseSet,omitempty"`    // Rise/Transit/Set times
}

// EntryHandler handles an Entry, used by ForEach
type EntryHandler func(*Entry) error

// Append an Entry to the Ephemeris
func (e *Ephemeris) Append(entry *Entry) {
	if entry != nil {
		e.Entries = append(e.Entries, entry)
	}
}

// ForEach will call an EntryHandler for each Entry within an Ephemeris
func (e *Ephemeris) ForEach(f EntryHandler) error {
	if e != nil {
		for _, entry := range e.Entries {
			if err := f(entry); err != nil {
				return err
			}
		}
	}
	return nil
}

// ForEachFiltered will call an EntryHandler for each Entry within an Ephemeris which passes a Predicate
func (e *Ephemeris) ForEachFiltered(f EntryHandler, p Predicate) error {
	if e != nil {
		for _, entry := range e.Entries {
			if p.Do(entry) {
				if err := f(entry); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (e *Ephemeris) Sort() *Ephemeris {
	sort.SliceStable(e.Entries, func(i, j int) bool {
		a := e.Entries[i]
		b := e.Entries[j]
		if a.Date.Before(b.Date) {
			return true
		}
		if a.Ord < b.Ord {
			return true
		}
		return a.Name < b.Name
	})
	return e
}

func (e *Ephemeris) CalculateRiseSetTimes() error {
	return e.ForEachFiltered(func(entry *Entry) error {
		th0 := entry.Date.Apparent0UT()
		eq := entry.Equatorial
		rs := eq.RiseSet(*e.Meta.LatLong.Coord(), th0, eq.Diameter)
		entry.RiseSet = &rs
		return nil
	}, func(entry *Entry) bool {
		// Only apply if we don't have an entry already, we have equatorial coordinates
		return entry.RiseSet == nil && entry.Equatorial != nil
	})
}
