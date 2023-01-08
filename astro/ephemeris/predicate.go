package ephemeris

// Predicate performs a test against an Entry
type Predicate func(*Entry) bool

// Do invokes a Predicate. If it is nil then this always returns true.
// This is the preferred way to invoke a predicate as this does the nil check.
func (p Predicate) Do(e *Entry) bool {
	if p != nil {
		return p(e)
	}
	return true
}

func True(_ *Entry) bool { return true }

func False(_ *Entry) bool { return false }

// And logical and of two predicates.
// If either is nil then this returns the other one as a nil predicate is always true
func And(a, b Predicate) Predicate {
	switch {
	case a == nil:
		return b
	case b == nil:
		return a
	default:
		// No need to use Do(e) here as both a and b are not nil.
		return func(e *Entry) bool {
			return a(e) && b(e)
		}
	}
}

// Or logical and of two predicates.
// If either is nil then this returns the other one as a nil predicate is always true
func Or(a, b Predicate) Predicate {
	switch {
	case a == nil:
		return b
	case b == nil:
		return a
	default:
		// No need to use Do(e) here as both a and b are not nil.
		return func(e *Entry) bool {
			return a(e) || b(e)
		}
	}
}

// Neg negates a Predicate
func Neg(a Predicate) Predicate {
	if a == nil {
		// As a nil predicate is always true then just return False
		return False
	}
	return func(e *Entry) bool {
		return !a(e)
	}
}
