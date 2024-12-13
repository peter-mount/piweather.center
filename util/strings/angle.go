package strings

import (
	"github.com/soniakeys/unit"
	"math"
	"strings"
)

const (
	toRad = math.Pi / 180.0
	toDeg = 180.0 / math.Pi
)

// Deg2Rad converts an angle in degrees to radians
func Deg2Rad(d float64) float64 {
	return d * toRad
}

// Rad2Deg converts an angle in radians to degrees
func Rad2Deg(r float64) float64 {
	return r * toDeg
}

func DegRange(d float64) float64 {
	for d < 0.0 {
		d = d + 360.0
	}
	for d >= 360.0 {
		d = d - 360.0
	}
	return d
}

func DegDMS(d float64) (int, int, float64) {
	deg, fDeg := math.Modf(d)
	m, s := math.Modf(fDeg * 60)
	s = s * 60.0
	return int(deg), int(m), s
}

func DegDMSString(d float64, sign bool) string {
	degDigits := 3
	if sign {
		degDigits = 2
	}

	return DegDMSStringExt(d, sign, "+", "-", degDigits, 0)
}

// DegDMSStringExt converts a float64 into a string consisting of degrees, minutes and seconds.
// This is the same as DegDMSStringExt2 with ":" separating the values and no suffix for seconds.
//
// d value to convert
//
// sign true if a sign should be included, false if not
//
// p & m the signs to use for positive and negative values, ignored if sign is false
//
// degDigits the number of digits for the degree value, usually 2 or 3
//
// precision the number of decimal places for seconds
func DegDMSStringExt(d float64, sign bool, p, m string, degDigits, precision int) string {
	return DegDMSStringExt2(d, sign, p, m, ":", ":", "", degDigits, precision)
}

// DegDMSStringExt2 converts a float64 into a string consisting of degrees, minutes and seconds
//
// d value to convert
//
// sign true if a sign should be included, false if not
//
// p & m the signs to use for positive and negative values, ignored if sign is false
//
// dSuffix the suffix for degrees but also the separator between degrees and minutes
//
// mSuffix the suffix for minutes but also the separator between minutes and seconds
//
// sSuffix the suffix for Seconds
//
// degDigits the number of digits for the degree value, usually 2 or 3
//
// precision the number of decimal places for seconds
func DegDMSStringExt2(d float64, sign bool, p, m, dSuffix, mSuffix, sSuffix string, degDigits, precision int) string {
	deg, minute, sec := DegDMS(math.Abs(d))

	var s []string
	if sign {
		if d < 0.0 {
			s = append(s, m)
		} else {
			s = append(s, p)
		}
	}

	// When formatting the output we cannot use strconv.FormatFloat or "%f" in Sprintf as those functions will
	// round the result up. When the value is very close to the upper bound of 360.0 (which is out of bounds) then
	// the seconds field ends up being 60 instead of 59.xxxx
	//
	// So we have to split seconds and then format using integers to ensure we don't get that result.
	//
	// This means that seconds effectively round down for all seconds but this is acceptable.
	s = append(s,
		Itoa(deg, degDigits),
		dSuffix,
		Itoa(minute, 2),
		mSuffix,
		FormatFloatDown(sec, 2, precision),
		sSuffix)

	return strings.TrimSpace(strings.Join(s, ""))
}

func HourDMSString(t unit.Time) string {
	return HourDMSStringExt(t.Hour())
}

func HourDMSStringExt(d float64) string {
	return DegDMSStringExt(d, false, "", "", 2, 0)
}
