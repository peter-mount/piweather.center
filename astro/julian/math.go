package julian

func (t Day) Before(b Day) bool {
	return t < b
}

func (t Day) After(b Day) bool {
	return t > b
}

func (t Day) Add(d float64) Day {
	return t + Day(d)
}

func Swap(a, b Day) (Day, Day) {
	return b, a
}
