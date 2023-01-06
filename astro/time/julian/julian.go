package julian

import (
	"math"
	"time"
)

type Day float64

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

func FromTime(t time.Time) Day {
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

func FromDate(y, m, d, h, min, s int) Day {
	return FromTime(time.Date(y, time.Month(m), d, h, min, s, 0, time.UTC))
}

func (t Day) Time() time.Time {
	z, f := math.Modf(float64(t) + 0.5)

	a := int(z)
	if z >= 2299161 {
		a1 := int(math.Floor((float64(z) - 1867216.25) / 36524.25))
		a = a + 1 + a1 - (a1 / 4)
	}

	b := float64(a + 1524)
	c := math.Floor((b - 122.1) / 365.25)
	d := math.Floor(365.25 * c)
	e := math.Floor((b - d) / 30.6001)

	month := e - 1
	if e >= 14 {
		month = e - 13
	}

	year := int(c) - 4715
	if month > 2 {
		year = int(c) - 4716
	}

	day, h, m, s := FdayToHMS(b - d - math.Floor(30.6001*e) + f)

	return time.Date(year, time.Month(int(month)), day, h, m, s, 0, time.UTC)
}
