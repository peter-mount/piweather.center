package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/util/source"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"sort"
	"strings"
	"testing"
)

func Test_voltage_transforms(t *testing.T) {
	testConversions(t, []conversionTest{
		newConversionTest(Volt.Value(1), MilliVolt.Value(1000.0), false),
		newConversionTest(Volt.Value(1), MicroVolt.Value(1000000.0), false),
		// dBV reference
		newConversionTest(Volt.Value(1), DecibelVolt.Value(0), false),
		// V -> dBV
		newConversionTest(MilliVolt.Value(100), DecibelVolt.Value(-20), false),
		newConversionTest(MilliVolt.Value(1), DecibelVolt.Value(-60), false),
		newConversionTest(Volt.Value(10), DecibelVolt.Value(20), false),
		newConversionTest(Volt.Value(100), DecibelVolt.Value(40), false),
	})
}

// Used by multiple tests, test we can convert between two values in both directions
type conversionTest struct {
	file    source.File
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

func newConversionTest(from value.Value, to value.Value, wantErr bool) conversionTest {
	return conversionTest{
		file:    source.SourceFileN(1),
		from:    from,
		to:      to,
		wantErr: wantErr,
	}
}

func testConversions(t *testing.T, tests []conversionTest) {
	// Form the actual tests for from->to and to->from, order by the keys
	type testUnit struct {
		conversionTest
		fromName  string
		toName    string
		fromValue value.Value
		toValue   value.Value
	}

	var ary []testUnit
	for _, test := range tests {
		ary = append(ary,
			// from -> to
			testUnit{
				fromName: test.from.Unit().Name(), fromValue: test.from,
				toName: test.to.Unit().Name(), toValue: test.to,
				conversionTest: test,
			},

			// to -> from
			testUnit{
				fromName: test.to.Unit().Name(), fromValue: test.to,
				toName: test.from.Unit().Name(), toValue: test.from,
				conversionTest: test,
			},
		)
	}

	sort.SliceStable(ary, func(i, j int) bool {
		from, to := ary[i], ary[j]
		c := strings.Compare(from.fromName, to.fromName)
		if c == 0 {
			c = strings.Compare(from.toName, to.toName)
			if c == 0 {
				c = strings.Compare(from.fromValue.String(), to.fromValue.String())
			}
		}
		return c < 0
	})

	var lastFrom string
	for _, testFrom := range ary {
		if lastFrom != testFrom.fromName {
			lastFrom = testFrom.fromName
			t.Run(testFrom.fromName,
				func(t *testing.T) {
					var lastTo string
					for _, testTo := range ary {
						if testTo.fromName == lastFrom && testTo.toName != lastTo {
							lastTo = testTo.toName

							t.Run(testTo.toName,
								func(t *testing.T) {
									for _, test := range ary {
										if test.fromName == lastFrom && test.toName == lastTo {

											t.Run(fmt.Sprintf("[%s] %s->%s", test.file, test.fromValue.PlainString(), test.toValue.PlainString()),
												func(t *testing.T) {
													got, err := test.fromValue.As(test.toValue.Unit())
													if err != nil {
														if !test.wantErr {
															t.Errorf("error = %v, wantErr %v", err, test.wantErr)
														}
														return
													}

													// Limit precision due to rounding errors
													if math.Abs(got.Float()-test.toValue.Float()) > 1e-4 {
														t.Errorf("from %s got = %.4f, want %.4f", test.fromValue.String(), got.Float(), test.toValue.Float())
													}
												})

										}
									}
								})

						}
					}
				})

		}
	}
}
