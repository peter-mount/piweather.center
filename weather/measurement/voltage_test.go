package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"sort"
	"strings"
	"testing"
)

func Test_voltage_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		{Volt.Value(1), MilliVolt.Value(1000.0), false},
		{Volt.Value(1), MicroVolt.Value(1000000.0), false},
		// dBV reference
		{Volt.Value(1), DecibelVolt.Value(0), false},
		// V -> dBV
		{MilliVolt.Value(100), DecibelVolt.Value(-20), false},
		{MilliVolt.Value(1), DecibelVolt.Value(-60), false},
		{Volt.Value(10), DecibelVolt.Value(20), false},
		{Volt.Value(100), DecibelVolt.Value(40), false},
	})
}

// Used by multiple tests, test we can convert between two values in both directions
type conversionTest struct {
	from    value.Value
	to      value.Value
	wantErr bool
}

func (a conversionTest) compare(b conversionTest) int {
	r := strings.Compare(a.from.Unit().ID(), b.from.Unit().ID())
	if r == 0 {
		r = strings.Compare(a.to.Unit().ID(), b.to.Unit().ID())
	}
	// Works as we know they are the same unit
	if r == 0 && a.from.Float() < b.from.Float() {
		r = -1
	}
	return r
}

func testConversions(t *testing.T, tests []conversionTest) {
	// Form the actual tests for from->to and to->from, order by the keys
	type testUnit struct {
		conversionTest
		name string
	}

	var ary []testUnit
	for _, test := range tests {
		ary = append(ary,
			testUnit{name: fmt.Sprintf("%s %s %s", test.from.Unit().Name(), test.to.Unit().Name(), test.from), conversionTest: test},
			testUnit{name: fmt.Sprintf("%s %s %s", test.to.Unit().Name(), test.from.Unit().Name(), test.to), conversionTest: test},
		)
	}

	sort.SliceStable(ary, func(i, j int) bool {
		return ary[i].name < ary[j].name
	})

	for _, test := range ary {
		t.Run(test.name, func(t *testing.T) {

			got, err := test.from.As(test.to.Unit())
			if err != nil {
				if !test.wantErr {
					t.Errorf("error = %v, wantErr %v", err, test.wantErr)
				}
				return
			}

			// Limit precision due to rounding errors
			if math.Abs(got.Float()-test.to.Float()) > 1e-4 {
				t.Errorf("from %s got = %.4f, want %.4f", test.from.String(), got.Float(), test.to.Float())
			}
		})
	}
}
