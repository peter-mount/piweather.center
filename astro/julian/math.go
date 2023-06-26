package julian

func (d Day) Before(b Day) bool {
	return d < b
}

func (d Day) After(b Day) bool {
	return d > b
}

func (d Day) Add(days float64) Day {
	return d + Day(days)
}

func Swap(a, b Day) (Day, Day) {
	return b, a
}
