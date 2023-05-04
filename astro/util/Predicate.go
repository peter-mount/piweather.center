package util

type Predicate func() bool

func (a Predicate) Not() Predicate {
	return func() bool {
		return !a()
	}
}

func (a Predicate) Or(b Predicate) Predicate {
	return func() bool {
		return a() || b()
	}
}

func (a Predicate) And(b Predicate) Predicate {
	return func() bool {
		return a() && b()
	}
}

func True() bool  { return true }
func False() bool { return false }
