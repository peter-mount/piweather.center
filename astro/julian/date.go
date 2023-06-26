package julian

// Date is the methods Day, Year and Century implement to allow
// conversions between each unit.
type Date interface {
	ToDay() Day
	ToYear() Year
	ToCentury() Century
}
