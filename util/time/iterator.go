package time

import (
	"github.com/peter-mount/go-kernel/v2/util"
	"time"
)

type iterator struct {
	start      time.Time     // The start time
	end        time.Time     // The end time, Zero when in count mode
	step       time.Duration // Step size
	startCount int           // the initial count, ignored if start & end are not zero
	current    time.Time     // The current value returned by Next()
	next       time.Time     // The next value after current in the sequence
	remaining  int           // the remaining count, ignored if start & end are not zero
	predicate  func() bool   // Predicate in use for HasNext()
}

// HasNext returns true if the iterator has not completed.
func (it *iterator) HasNext() bool {
	r := it.predicate()
	if r {
		it.next = it.current.Add(it.step)
		it.remaining--
	}
	return r
}

// Next returns the next time.Time in the iteration.
func (it *iterator) Next() time.Time {
	r := it.current
	it.current = it.next
	return r
}

func (it *iterator) init() util.Iterator[time.Time] {
	startZero := it.start.IsZero()
	endZero := it.end.IsZero()
	countMode := it.startCount > 0

	switch {
	case startZero && endZero,
		!startZero && !endZero && it.start.Equal(it.end):
		// Do nothing if both zero or equal
		it.step = 0
		it.startCount = 0
		it.predicate = it.hasNone

	case endZero && countMode:
		// start is valid so run from there using count
		it.predicate = it.hasRemaining
		it.current = it.start

	case startZero && countMode:
		// start is zero so swap it with end
		it.start, it.end = it.end, it.start
		// negate step as we need to go in the reverse direction
		it.step = -it.step
		it.predicate = it.hasRemaining
		it.current = it.start

	case endZero:
		// start is valid so run from there using count
		it.predicate = it.hasRemaining
		it.current = it.start

	case startZero:
		// end is valid so use it as the start. should not happen but it's possible with ReverseIterator
		it.start, it.end = it.end, it.start
		it.predicate = it.hasRemaining
		it.current = it.start

	case it.start.Before(it.end):
		// Step must be positive to reach end
		it.step = it.step.Abs()
		it.predicate = it.hasPositive
		it.current = it.start

	case it.start.After(it.end):
		// Swap start & negate step
		it.start, it.end = it.end, it.start
		it.step = -it.step.Abs()
		it.predicate = it.hasNegative
		it.current = it.end

	}

	it.next = it.current
	it.remaining = it.startCount

	return it
}

func (it *iterator) hasPositive() bool {
	return it.current.Before(it.end)
}

func (it *iterator) hasNegative() bool {
	return it.current.After(it.start)
}

// HasNext() for when start==end
func (it *iterator) hasNone() bool {
	return false
}

func (it *iterator) hasRemaining() bool {
	return it.remaining > 0
}

func (it *iterator) ForEach(f func(time time.Time)) {
	for it.HasNext() {
		f(it.Next())
	}
}

func (it *iterator) ForEachAsync(f func(time time.Time)) {
	it.ForEach(f)
}

func (it *iterator) ForEachFailFast(f func(time time.Time) error) error {
	for it.HasNext() {
		if err := f(it.Next()); err != nil {
			return err
		}
	}
	return nil
}

func (it *iterator) clone() *iterator {
	return &iterator{
		start:      it.start,
		end:        it.end,
		step:       it.step,
		startCount: it.startCount,
	}
}

func (it *iterator) Iterator() util.Iterator[time.Time] {
	return it.clone().init()
}

func (it *iterator) ReverseIterator() util.Iterator[time.Time] {
	r := it.clone()
	// Reverse the start/end times. init() will handle changes to step
	r.start, r.end = r.end, r.start
	return r.init()
}

// IterateBetween returns an Iterator which will iterate between two time.Time's with the
// specified time.Duration between each step.
//
// This implementation will only include whole step's in its run.
// If the last step ends after the end time then it will not be included.
func IterateBetween(start, end time.Time, step time.Duration) util.Iterator[time.Time] {
	it := iterator{start: start, end: end, step: step}
	return it.init()
}

// IterateFrom returns an Iterator which will iterate from a time.Time count times, adding step between
// each step.
func IterateFrom(start time.Time, step time.Duration, count int) util.Iterator[time.Time] {
	it := iterator{start: start, startCount: count, step: step}
	return it.init()
}
