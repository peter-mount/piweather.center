package julian

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/sidereal"
	"github.com/soniakeys/unit"
	"math"
	"time"
)

type Day float64

// JD returns the Julian Date of a day
// Deprecated
func (d Day) JD() float64 {
	return float64(d)
}

// Float is useful to convert a Day to a Float
func (d Day) Float() float64 {
	return float64(d)
}

func (d Day) Day() Day {
	return d
}

func (d Day) Year() Year {
	return Year(float64(d) / JulianYear)
}

func (d Day) Century() Century {
	return d.Year().Century()
}

func (y Year) Century() Century {
	return Century(float64(y) / JulianCentury)
}

const (
	// JulianYear (a) = 365.25 days
	JulianYear    = 365.25
	JulianCentury = 100 * JulianYear
)

// IsGregorian returns true if the date is in the Gregorian Calendar.
// Here we take the original date of the reform so dates on or before 1582 Oct 4 is
// in the Julian calendar (returns false) and those on or after 1582 Oct 15 are in
// the Gregorian Calendar (returns true).
//
// Although invalid, dates 15823 Oct 5..14 do not exist but return true here.
//
// Note: not all countries changed on that date, Great Britain changed in 1752 and Turkey
// in 1927, but we stick to the original dates here.
func IsGregorian(d, m, y int) bool {
	switch {
	// Julian Calendar
	case y < 1582,
		y == 1582 && m < 10,
		y == 1582 && m == 10 && d <= 4:
		return false

	// Gregorian Calendar
	case y > 1582,
		y == 1582 && m > 10,
		y == 1582 && m == 10 && d >= 15:
		return true

	// Invalid date
	default:
		return true
	}
}

// Now returns a Day for the current time
func Now() Day {
	return FromTime(time.Now().UTC())
}

// StartOfToday returns a Day for the start of the current day (00:00 UTC)
func StartOfToday() Day {
	return FromTime(time.Now().UTC().Truncate(time.Hour * 24))
}

// FromTime returns a Day from a time.Time
func FromTime(t0 time.Time) Day {
	// Enforce the time into UTC
	t := t0.UTC()
	Y := t.Year()
	M := int(t.Month())
	D := t.Day()
	F := (float64(t.Hour()) + (float64(t.Minute()) / 60.0) + (float64(t.Second()) / 3600.0)) / 24.0

	if M < 3 {
		Y = Y - 1
		M = M + 12
	}

	B := 0
	if IsGregorian(D, M, Y) {
		A := Y / 100
		B = 2 - A + (A / 4)
	}

	return Day(math.Floor(365.25*float64(Y+4716)) + math.Floor(30.6001*float64(M+1)) + float64(D+B) - 1524.5 + F)
}

// FromDate returns a Day from a calendar date
func FromDate(y, m, d, h, min, s int) Day {
	return FromTime(time.Date(y, time.Month(m), d, h, min, s, 0, time.UTC))
}

// Time returns the time.Time of a Day
func (d Day) Time() time.Time {
	z, f := math.Modf(d.Float() + 0.5)

	a := int(z)
	if z >= 2299161 {
		a1 := int(math.Floor((float64(z) - 1867216.25) / 36524.25))
		a = a + 1 + a1 - (a1 / 4)
	}

	b := float64(a + 1524)
	c := math.Floor((b - 122.1) / 365.25)
	d1 := math.Floor(365.25 * c)
	e := math.Floor((b - d1) / 30.6001)

	month := e - 1
	if e >= 14 {
		month = e - 13
	}

	year := int(c) - 4715
	if month > 2 {
		year = int(c) - 4716
	}

	day, h, m, s := util.FdayToHMS(b - d1 - math.Floor(30.6001*e) + f)

	return time.Date(year, time.Month(int(month)), day, h, m, s, 0, time.UTC)
}

// CenturiesJ2k returns the number of Julian Centuries since J2000.0 (2000 Jan 1.5 TD)
func (d Day) CenturiesJ2k() float64 {
	return (d.Float() - 2451545.0) / 36525.0
}

func (d Day) String() string {
	return fmt.Sprintf("%f", d.Float())
}

func (d Day) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// JDMidnight returns the Day at 0h
func (d Day) JDMidnight() Day {
	// Julian days start at midday hence subtract 0.5 before converting to integer then adding 0.5 back.
	return Day(math.Floor(d.Float()-0.5) + 0.5)
}

// Apparent returns the apparent sidereal time at Greenwich for this Day
func (d Day) Apparent() unit.Time {
	return sidereal.Apparent(d.Float())
}

// Apparent0UT returns the apparent sidereal time at Greenwich at 0UT on this Day
func (d Day) Apparent0UT() unit.Time {
	return sidereal.Apparent0UT(d.Float())
}

// Mean returns the mean sidereal time at Greenwich for this Day
func (d Day) Mean() unit.Time {
	return sidereal.Mean(d.Float())
}

// Mean0UT returns the mean sidereal time at Greenwich at 0UT on this Day
func (d Day) Mean0UT() unit.Time {
	return sidereal.Mean0UT(d.Float())
}

func (d Day) IsGregorian() bool {
	t1 := d.Time()
	return IsGregorian(t1.Day(), int(t1.Month()), t1.Year())
}
