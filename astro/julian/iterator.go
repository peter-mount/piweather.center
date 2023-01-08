package julian

import (
	"errors"
)

// Iterator allows for iteration between two Days.
// In use, you create one via Iterate providing a start, end Day and the step size in days.
// Then you call HasNext to see if another Day exists then call Next to get that Day.
// The ForEach function allows you to run through the Iterator without the boilerplate code.
//
// A nil Iterator will act as if the iterator is empty.
type Iterator struct {
	day       Day       // Current Day in iterator
	end       Day       // The end Day to stop the iterator
	step      float64   // The step in days between each iteration
	predicate Predicate // The Predicate to determine when to continue
	first     bool      // overrides predicate if true, used for when we want a single run to be performed
}

// Iterate will create an Iterator to iterate between two Day's with a specific step size.
// If step is positive then the iterator will run from the start to the end.
// If step is negative then the iterator will run from the end to the start.
// If step is 0 then it will run only once for the start Day. The end Day will be ignored.
func Iterate(a, b Day, step float64) *Iterator {
	// Ensure a Before(b) is always true
	if a.After(b) {
		a, b = Swap(a, b)
	}

	i := &Iterator{step: step, first: true}

	// Configure predicate depending on the step size
	switch {
	// increasing time start -> end every step
	case step > 0:
		i.day = a
		i.end = b
		i.predicate = BeforeEqual
	// reverse time end -> start every step
	case step < 0:
		i.day = b
		i.end = a
		i.predicate = AfterEqual
	// Run once, step==0
	default:
		i.day = a
		i.end = a
		i.predicate = False
	}

	return i
}

// HasNext returns true if there is another iteration possible
func (i *Iterator) HasNext() bool {
	return i != nil && (i.first || i.predicate(i.day, i.end))
}

// Next returns the next Day in the iteration.
// If the iterator does not have another Day (e.g. HasNext() returns false) this will panic.
func (i *Iterator) Next() Day {
	if i.HasNext() {
		r := i.day
		i.first = false
		i.day = i.day.Add(i.step)
		return r
	}

	// Iterator has no more values so panic as HasNext() should have returned false.
	// Caused by the iterator being called incorrectly and this should prevent infinite loops.
	panic(iteratorCompleted)
}

// iteratorCompleted is the error in the panic when Next() is called when there is no next value
var iteratorCompleted = errors.New("iterator competed")

// IteratorHandler is a Function called by ForEach
type IteratorHandler func(Day) error

// ForEach will call an IteratorHandler for each remaining iteration.
// If the handler returns an error it will stop at that point.
func (i *Iterator) ForEach(f IteratorHandler) error {
	for i.HasNext() {
		err := f(i.Next())
		if err != nil {
			return err
		}
	}
	return nil
}

// Slice will return a slice of all Day's in the iterator.
func (i *Iterator) Slice() []Day {
	var s []Day
	for i.HasNext() {
		s = append(s, i.Next())
	}
	return s
}
