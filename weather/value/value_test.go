package value

import (
	"fmt"
	"math"
	"testing"
)

type comparatorTest struct {
	a    float64
	b    float64
	want bool
}

func runComparatorTest(t *testing.T, f Comparator, tests []comparatorTest) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%.3f %.3f", tt.a, tt.b), func(t *testing.T) {
			if got := f(tt.a, tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equal(t *testing.T) {
	runComparatorTest(t,
		Equal,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, false},
			{math.MaxFloat64, 18.1, false},
			{18.1, -273.15, false},
			{42, 42, true},
			{42.1, 42.0, false},
			{12.0, 12.0000000001, true},
			{12.0, 11.9999999999, true},
			{12.0, 12.000000001, false},
			{12.0, 11.999999999, false},
		},
	)
}

func Test_notEqual(t *testing.T) {
	runComparatorTest(t,
		NotEqual,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, true},
			{math.MaxFloat64, 18.1, true},
			{18.1, -273.15, true},
			{42, 42, false},
			{42.1, 42.0, true},
			{12.0, 12.0000000001, false},
			{12.0, 11.9999999999, false},
			{12.0, 12.000000001, true},
			{12.0, 11.999999999, true},
		},
	)
}

func Test_lessThan(t *testing.T) {
	runComparatorTest(t,
		LessThan,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, true},
			{math.MaxFloat64, 18.1, false},
			{18.1, -273.15, false},
			{21, 0, false},
			{21, 30, true},
			{30, 21, false},
			{42, 42, false},
			{12.0000000001, 12.0, false},
			{11.9999999999, 12.0, false},
			{12.000000001, 12.0, false},
			{11.999999999, 12.0, true},
		},
	)
}

func Test_lessThanEqual(t *testing.T) {
	runComparatorTest(t,
		LessThanEqual,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, true},
			{math.MaxFloat64, 18.1, false},
			{18.1, -273.15, false},
			{21, 0, false},
			{21, 30, true},
			{30, 21, false},
			{42, 42, true},
			{12.0000000001, 12.0, true},
			{11.9999999999, 12.0, true},
			{12.000000001, 12.0, false},
			{11.999999999, 12.0, true},
		},
	)
}

func Test_greaterThan(t *testing.T) {
	runComparatorTest(t,
		GreaterThan,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, false},
			{math.MaxFloat64, 18.1, true},
			{18.1, -273.15, true},
			{21, 0, true},
			{21, 30, false},
			{30, 21, true},
			{42, 42, false},
			{12.0, 12.0000000001, false},
			{12.0, 11.9999999999, false},
			{12.0, 12.000000001, false},
			{12.0, 11.999999999, true},
		},
	)
}

func Test_greaterThanEqual(t *testing.T) {
	runComparatorTest(t,
		GreaterThanEqual,
		[]comparatorTest{
			{-math.MaxFloat64, 18.1, false},
			{math.MaxFloat64, 18.1, true},
			{18.1, -273.15, true},
			{21, 0, true},
			{21, 30, false},
			{30, 21, true},
			{42, 42, true},
			{12.0, 12.0000000001, true},
			{12.0, 11.9999999999, true},
			{12.0, 12.000000001, false},
			{12.0, 11.999999999, true},
		},
	)
}
