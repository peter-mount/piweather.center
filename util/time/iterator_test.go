package time

import (
	"github.com/peter-mount/go-kernel/v2/util"
	"testing"
	"time"
)

// testIteratorCount counts the number of entries within an iterator and fails if it isn't the expected number
func testIteratorCount(t *testing.T, name string, it util.Iterator[time.Time], expectedCount int, expectedEnd, expectedReverse time.Time) {
	count := 0
	var lastTime time.Time
	for it.HasNext() {
		lastTime = it.Next()
		count++
	}
	testAssertExpected(t, name, count, expectedCount, lastTime, expectedEnd, expectedReverse)
}

func testAssertExpected(t *testing.T, name string, count, expectedCount int, lastTime, expectedEnd, expectedReverse time.Time) {
	if count != expectedCount {
		t.Errorf("%s returned %d entries, expected %d", name, count, expectedCount)
	}

	le1 := lastTime.Equal(expectedEnd)
	le2 := lastTime.Equal(expectedReverse)
	if !(!le1 && le2 || le1 && !le2) {
		t.Errorf("%s returned %q, expected %q", name, lastTime.Format(RFC3339), expectedEnd.Format(RFC3339))
	}
}

type testIteratorFactory func() util.Iterator[time.Time]

// testIteratorCountAll tests the iterator using all the methods within the util.Interator interface
func testIteratorCountAll(t *testing.T, name string, f testIteratorFactory, expectedCount int, expectedEnd, expectedReverse time.Time) {
	testIteratorCountAllInner(t, name, f, expectedCount, expectedEnd, expectedReverse, true)
}

func testIteratorCountAllInner(t *testing.T, name string, f testIteratorFactory, expectedCount int, expectedEnd, expectedReverse time.Time, nest bool) {
	t.Run("raw", func(t *testing.T) {
		testIteratorCount(t, name, f(), expectedCount, expectedEnd, expectedReverse)
	})

	if nest {
		// Run the same tests again but for the functions returning an Iterator

		t.Run("Iterator", func(t *testing.T) {
			testIteratorCountAllInner(t, name+".Iterator()",
				func() util.Iterator[time.Time] {
					return f().Iterator()
				},
				expectedCount, expectedEnd, expectedReverse, false)
		})

		t.Run("ReverseIterator", func(t *testing.T) {
			testIteratorCountAllInner(t, name+".ReverseIterator()",
				func() util.Iterator[time.Time] {
					return f().ReverseIterator()
				},
				expectedCount,
				// these are also reversed
				expectedEnd, expectedReverse,
				false)
		})
	}

	t.Run("ForEach", func(t *testing.T) {
		count := 0
		var lastTime time.Time
		f().ForEach(func(t time.Time) {
			count++
			lastTime = t
		})
		testAssertExpected(t, name+".ForEach()", count, expectedCount, lastTime, expectedEnd, expectedReverse)
	})

	t.Run("ForEachAsync", func(t *testing.T) {
		count := 0
		var lastTime time.Time
		f().ForEachAsync(func(t time.Time) {
			count++
			lastTime = t
		})
		testAssertExpected(t, name+".ForEachAsync()", count, expectedCount, lastTime, expectedEnd, expectedReverse)
	})

	t.Run("ForEachFailFast", func(t *testing.T) {
		count := 0
		var lastTime time.Time
		_ = f().ForEachFailFast(func(t time.Time) error {
			count++
			lastTime = t
			return nil
		})
		testAssertExpected(t, name+".ForEachFailFast()", count, expectedCount, lastTime, expectedEnd, expectedReverse)
	})

}

func TestIterateBetween(t *testing.T) {

	tzUK, err := time.LoadLocation("Europe/London")
	if err != nil {
		t.Fatalf("Cannot load UK timezone %v", err)
	}

	const (
		stepDay   = 24 * time.Hour
		stepHour  = time.Hour
		month31   = 31
		month30   = 30
		febLeap   = 29
		febNormal = 28
		dayNormal = 24 // Normal day, 24 hours long
		dayToBST  = 23 // GMT->BST during day so only 23 hours long
		dayToGMT  = 25 // BST->GMT during day so 25 hours long
	)

	tests := []struct {
		name            string
		start           string
		end             string
		step            time.Duration
		expected        int            // Expected number of iterations
		expectedEnd     string         // Expected value to be returned on the last iteration
		expectedReverse string         // Expected value to be returned on the last iteration on reversed iterations
		tz              *time.Location // TimeZone
	}{
		// 1 hour during the day
		{name: "1h", start: "2024-11-01 06:00:00", end: "2024-11-01T07:00:00", expectedEnd: "2024-11-01T06:59:00", expectedReverse: "2024-11-01T06:01:00", expected: 60, tz: time.UTC, step: time.Minute},
		// 24 hours in UTC
		{name: "24h Jul 2024", start: "2024-07-01", end: "2024-08-01", expectedEnd: "2024-07-31", expectedReverse: "2024-07-02", expected: month31, tz: time.UTC, step: stepDay},
		{name: "24h Aug 2024", start: "2024-08-01", end: "2024-09-01", expectedEnd: "2024-08-31", expectedReverse: "2024-08-02", expected: month31, tz: time.UTC, step: stepDay},
		{name: "24h Sep 2024", start: "2024-09-01", end: "2024-10-01", expectedEnd: "2024-09-30", expectedReverse: "2024-09-02", expected: month30, tz: time.UTC, step: stepDay},
		{name: "24h Oct 2024", start: "2024-10-01", end: "2024-11-01", expectedEnd: "2024-10-31", expectedReverse: "2024-10-02", expected: month31, tz: time.UTC, step: stepDay},
		{name: "24h Nov 2024", start: "2024-11-01", end: "2024-12-01", expectedEnd: "2024-11-30", expectedReverse: "2024-11-02", expected: month30, tz: time.UTC, step: stepDay},
		{name: "24h Dec 2024", start: "2024-12-01", end: "2025-01-01", expectedEnd: "2024-12-31", expectedReverse: "2024-12-02", expected: month31, tz: time.UTC, step: stepDay},
		{name: "24h Jan 2024", start: "2025-01-01", end: "2025-02-01", expectedEnd: "2025-01-31", expectedReverse: "2025-01-02", expected: month31, tz: time.UTC, step: stepDay},
		{name: "24h Feb 2024", start: "2024-02-01", end: "2024-03-01", expectedEnd: "2024-02-29", expectedReverse: "2024-02-02", expected: febLeap, tz: time.UTC, step: stepDay},
		{name: "24h Feb 2024", start: "2025-02-01", end: "2025-03-01", expectedEnd: "2025-02-28", expectedReverse: "2025-02-02", expected: febNormal, tz: time.UTC, step: stepDay},
		//
		// Iterate over a day in Europe/London covering dates around daylight saving changes.
		//
		// So for the day GMT->BST will be 23 hours long and BST->GMT will be 25 hours long.
		//
		// This applies to other time zones (but on different dates),
		// except for Lord Howe Island. We don't talk about Lord Howe Island.
		//
		// Source for the 2025, 2026 and 2027 dates: https://www.gov.uk/when-do-the-clocks-change
		//
		// 2025 BST starts 30 March and ends 26 October
		//
		// GMT -> BST
		{name: "29 Mar 2025", start: "2025-03-29", end: "2025-03-30", expectedEnd: "2025-03-29T23:00:00+00:00", expectedReverse: "2025-03-29T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "30 Mar 2025", start: "2025-03-30", end: "2025-03-31", expectedEnd: "2025-03-30T23:00:00+01:00", expectedReverse: "2025-03-30T01:00:00+00:00", expected: dayToBST, tz: tzUK, step: stepHour},
		{name: "31 Mar 2025", start: "2025-03-31", end: "2025-04-01", expectedEnd: "2025-03-31T23:00:00+01:00", expectedReverse: "2025-03-31T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		// BST -> GMT
		{name: "25 Oct 2025", start: "2025-10-25", end: "2025-10-26", expectedEnd: "2025-10-25T23:00:00+01:00", expectedReverse: "2025-10-25T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "26 Oct 2025", start: "2025-10-26", end: "2025-10-27", expectedEnd: "2025-10-26T23:00:00+00:00", expectedReverse: "2025-10-26T01:00:00+01:00", expected: dayToGMT, tz: tzUK, step: stepHour},
		{name: "27 Oct 2025", start: "2025-10-27", end: "2025-10-28", expectedEnd: "2025-10-27T23:00:00+00:00", expectedReverse: "2025-10-27T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
		//
		// 2026 BST starts 29 March and ends 25 October
		//
		// GMT -> BST
		{name: "28 Mar 2026", start: "2026-03-28", end: "2026-03-29", expectedEnd: "2026-03-28T23:00:00+00:00", expectedReverse: "2026-03-28T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "29 Mar 2026", start: "2026-03-29", end: "2026-03-30", expectedEnd: "2026-03-29T23:00:00+01:00", expectedReverse: "2026-03-29T01:00:00+00:00", expected: dayToBST, tz: tzUK, step: stepHour},
		{name: "30 Mar 2026", start: "2026-03-30", end: "2026-03-31", expectedEnd: "2026-03-30T23:00:00+01:00", expectedReverse: "2026-03-30T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		// BST -> GMT
		{name: "24 Oct 2026", start: "2026-10-24", end: "2026-10-25", expectedEnd: "2026-10-24T23:00:00+01:00", expectedReverse: "2026-10-24T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "25 Oct 2026", start: "2026-10-25", end: "2026-10-26", expectedEnd: "2026-10-25T23:00:00+00:00", expectedReverse: "2026-10-25T01:00:00+01:00", expected: dayToGMT, tz: tzUK, step: stepHour},
		{name: "26 Oct 2026", start: "2026-10-26", end: "2026-10-27", expectedEnd: "2026-10-26T23:00:00+00:00", expectedReverse: "2026-10-26T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
		//
		// 2025 BST starts 28 March and ends 31 October
		//
		// GMT -> BST
		{name: "27 Mar 2027", start: "2027-03-27", end: "2027-03-28", expectedEnd: "2027-03-27T23:00:00+00:00", expectedReverse: "2027-03-27T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "28 Mar 2027", start: "2027-03-28", end: "2027-03-29", expectedEnd: "2027-03-28T23:00:00+01:00", expectedReverse: "2027-03-28T01:00:00+00:00", expected: dayToBST, tz: tzUK, step: stepHour},
		{name: "29 Mar 2027", start: "2027-03-29", end: "2027-03-30", expectedEnd: "2027-03-29T23:00:00+01:00", expectedReverse: "2027-03-29T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		// BST -> GMT
		{name: "30 Oct 2027", start: "2027-10-30", end: "2027-10-31", expectedEnd: "2027-10-30T23:00:00+01:00", expectedReverse: "2027-10-30T01:00:00+01:00", expected: dayNormal, tz: tzUK, step: stepHour},
		{name: "31 Oct 2027", start: "2027-10-31", end: "2027-11-01", expectedEnd: "2027-10-31T23:00:00+00:00", expectedReverse: "2027-10-31T01:00:00+01:00", expected: dayToGMT, tz: tzUK, step: stepHour},
		{name: "01 Nov 2027", start: "2027-11-01", end: "2027-11-02", expectedEnd: "2027-11-01T23:00:00+00:00", expectedReverse: "2027-11-01T01:00:00+00:00", expected: dayNormal, tz: tzUK, step: stepHour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := ParseTimeIn(tt.start, tt.tz)
			end := ParseTimeIn(tt.end, tt.tz)
			expectedEnd := ParseTimeIn(tt.expectedEnd, tt.tz)
			expectedReverse := ParseTimeIn(tt.expectedReverse, tt.tz)

			// test start->end
			t.Run("start<end", func(t *testing.T) {
				testIteratorCountAll(t, "IterateBetween()",
					func() util.Iterator[time.Time] {
						return IterateBetween(start, end, tt.step)
					},
					tt.expected, expectedEnd, expectedReverse)
			})

			// test end->start
			t.Run("start>end", func(t *testing.T) {
				testIteratorCountAll(t, "IterateBetween()",
					func() util.Iterator[time.Time] {
						return IterateBetween(end, start, tt.step)
					},
					tt.expected, expectedEnd, expectedReverse)
			})
		})
	}
}

func TestIterateFrom(t *testing.T) {

	//tzUK, err := time.LoadLocation("Europe/London")
	//if err != nil {
	//	t.Fatalf("Cannot load UK timezone %v", err)
	//}

	const (
		stepMinute = time.Minute
		stepDay    = 24 * time.Hour
		stepHour   = time.Hour
	)

	tests := []struct {
		name            string
		start           string
		step            time.Duration
		count           int
		expectedCount   int
		expectedEnd     string
		expectedReverse string
		tz              *time.Location
	}{
		// The iterator should run once per minute
		{name: "1h", start: "2024-11-01T05:00:00Z", expectedEnd: "2024-11-01T05:59:00Z", expectedReverse: "2024-11-01T04:01:00Z", expectedCount: 60, step: stepMinute, tz: time.UTC},
		{name: "24h by hr", start: "2024-11-01", expectedEnd: "2024-11-01T23:00:00Z", expectedReverse: "2024-10-31T01:00:00Z", expectedCount: 24, step: stepHour, tz: time.UTC},
		{name: "24h by mn", start: "2024-11-01", expectedEnd: "2024-11-01T23:59:00Z", expectedReverse: "2024-10-31T00:01:00Z", expectedCount: 24 * 60, step: stepMinute, tz: time.UTC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := ParseTimeIn(tt.start, tt.tz)
			expectedEnd := ParseTime(tt.expectedEnd)
			expectedReverse := ParseTime(tt.expectedReverse)

			testIteratorCountAll(t, "IterateFrom",
				func() util.Iterator[time.Time] {
					return IterateFrom(start, tt.step, tt.expectedCount)
				},
				tt.expectedCount, expectedEnd, expectedReverse)
		})
	}
}
