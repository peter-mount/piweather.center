package catalogue

import (
	"fmt"
	"github.com/peter-mount/piweather.center/util/strings"
	"github.com/soniakeys/unit"
)

// Entry represents an entry within a Catalog
type Entry struct {
	ra  int32 // RA J2000 unit = 1 arc sec
	dec int32 // Dec J2000 unit = 1 arc sec
	mag int16 // Visual Magnitude, unit=0.01 mag
}

// NewEntry creates a new Entry for the specified object
func NewEntry(ra unit.RA, dec unit.Angle, mag float64) Entry {
	return Entry{
		ra:  int32(ra.Hour() * 3600.0),
		dec: int32(dec.Deg() * 3600.0),
		mag: int16(mag * 100.0),
	}
}

// RA The Right Ascension of the Entry
func (e Entry) RA() unit.RA {
	return unit.RAFromHour(float64(e.ra) / 3600.0)
}

// Dec The Declination of the Entry
func (e Entry) Dec() unit.Angle {
	return unit.AngleFromDeg(float64(e.dec) / 3600.0)
}

// Mag the Magnitude of the Entry
func (e Entry) Mag() float64 {
	return float64(e.mag) / 100.0
}

// String the Entry in a fixed format
func (e Entry) String() string {
	return fmt.Sprintf("[%q,%q,%.2f]",
		strings.HourDMSStringExt(e.RA().Hour()),
		strings.DegDMSString(e.Dec().Deg(), true),
		e.Mag(),
	)
}

// IsValid returns true if the entry is valid.
// Specifically it is not a 0 Magnitude object at RA 0.0 Dec 0.0
func (e Entry) IsValid() bool {
	return !(e.mag == 0 && e.ra == 0 && e.dec == 0)
}

// Equals returns true of the two Entries are identical
func (e Entry) Equals(b Entry) bool {
	//return e.String() == b.String()
	return e.ra == b.ra && e.dec == b.dec && e.mag == b.mag
}
