package julian

// Year (symbol: a or aj) is a unit of measurement of time defined as
// exactly 365.25 days of 86400 SI seconds each.
//
// The length of the Julian year is the average length of the year in
// the Julian calendar that was used in Western societies until the
// adoption of the Gregorian Calendar, and from which the unit is named.
// Nevertheless, because astronomical Julian years are measuring duration
// rather than designating dates, this Julian year does not correspond
// to years in the Julian calendar or any other calendar.
// Nor does it correspond to the many other ways of defining a year.
type Year float64

func (y Year) Float() float64 {
	return float64(y)
}

func (y Year) ToDay() Day {
	return Day(float64(y) * JulianYear)
}

func (y Year) ToYear() Year {
	return y
}
