package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
	"testing"
)

func Test_angle(t *testing.T) {
	testConversions(t, []conversionTest{
		// =========================
		// ArcMinute
		// =========================
		{ArcMinute.Value(1), Degree.Value(0.0166666666), false},
		{ArcMinute.Value(1), ArcSecond.Value(60), false},
		{ArcMinute.Value(1), HourAngle.Value(0.001111111), false},
		{ArcMinute.Value(1), Gradian.Value(0.018518518), false},
		{ArcMinute.Value(1), Radian.Value(0.000290888), false},
		{ArcMinute.Value(1), Turn.Value(0.000046296), false},
		// =========================
		// ArcSecond
		// =========================
		{ArcSecond.Value(1), Degree.Value(0.000277777), false},
		{ArcSecond.Value(1), ArcMinute.Value(0.0166666666), false},
		{ArcSecond.Value(1), HourAngle.Value(0.000018518), false},
		{ArcSecond.Value(1), Gradian.Value(0.000308641), false},
		{ArcSecond.Value(1), Radian.Value(0.000004848), false},
		{ArcSecond.Value(1), Turn.Value(0.000000771), false},
		// =========================
		// Degree
		// =========================
		{Degree.Value(1), ArcMinute.Value(60), false},
		{Degree.Value(1), ArcSecond.Value(3600), false},
		{Degree.Value(1), HourAngle.Value(0.0666666666), false},
		{Degree.Value(1), Gradian.Value(1.1111111111), false},
		{Degree.Value(1), Radian.Value(0.017453292), false},
		{Degree.Value(1), Turn.Value(0.002777777), false},
		// =========================
		// Gradian
		// =========================
		{Gradian.Value(1), Degree.Value(0.9), false},
		{Gradian.Value(1), ArcMinute.Value(54), false},
		{Gradian.Value(1), ArcSecond.Value(3240), false},
		{Gradian.Value(1), HourAngle.Value(0.06), false},
		{Gradian.Value(1), Radian.Value(0.015707963), false},
		{Gradian.Value(1), Turn.Value(0.0025), false},
		// =========================
		// Hour Angle
		// =========================
		{HourAngle.Value(1), Degree.Value(15), false},
		{HourAngle.Value(1), ArcMinute.Value(900), false},
		{HourAngle.Value(1), ArcSecond.Value(54000), false},
		{HourAngle.Value(1), Gradian.Value(16.666666667), false},
		{HourAngle.Value(1), Radian.Value(0.261799387), false},
		{HourAngle.Value(1), Turn.Value(0.041666666), false},
		// =========================
		// Radian
		// =========================
		{Radian.Value(1), ArcMinute.Value(3437.746771), false},
		{Radian.Value(1), ArcSecond.Value(206264.8062), false},
		{Radian.Value(1), Degree.Value(57.29577951), false},
		{Radian.Value(1), Gradian.Value(63.66197724), false},
		{Radian.Value(1), Turn.Value(0.159154943), false},
		// =========================
		// Turn
		// =========================
		{Turn.Value(1), ArcMinute.Value(21600), false},
		{Turn.Value(1), ArcSecond.Value(1296000), false},
		{Turn.Value(1), Degree.Value(360), false},
		{Turn.Value(1), Gradian.Value(400), false},
		{Turn.Value(1), HourAngle.Value(24), false},
		{Turn.Value(1), Radian.Value(math.Pi * 2), false},
	})
}

func Test_angle_bounds(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr string
	}{
		// ============================
		// RA min 0 max 23:59:59.999...
		// ============================
		// Valid
		{from: Turn.Value(0), to: RA.Value(0)},
		{from: Turn.Value(0.5), to: RA.Value(12)},
		// Invalid
		{from: Turn.Value(-1), to: RA.Value(-1), wantErr: "out of bounds"},
		{from: Turn.Value(1), to: RA.Value(24), wantErr: "out of bounds"},
		// ------------------------------------------------------------
		// These caused issues with seconds rounding up to 60
		// returned 23:59:59.9 which is correct
		{from: Degree.Value(15.0 * (24 - (0.1 / 3600.0))), to: RA.Value(0)},
		// returned 23:59:60.0 incorrect
		{from: Degree.Value(15.0 * (24 - (0.01 / 3600.0))), to: RA.Value(0)},
		// returned 23:59:60.0 incorrect
		{from: Degree.Value(15.0 * (24 - (0.001 / 3600.0))), to: RA.Value(0)},
		// ===========================
		// Declination min -90 max +90
		// ===========================
		// Valid
		{from: Degree.Value(10), to: Declination.Value(10)},
		{from: Degree.Value(89.9999), to: Declination.Value(89.9999)},
		{from: Degree.Value(-10), to: Declination.Value(-10)},
		{from: Degree.Value(-89.9999), to: Declination.Value(-89.9999)},
		{from: Degree.Value(90), to: Declination.Value(90)},
		{from: Degree.Value(-90), to: Declination.Value(-90)},
		// Invalid
		{from: Degree.Value(-90.0000001), to: Declination.Value(-90.0000001), wantErr: "out of bounds"},
		{from: Degree.Value(90.0000001), to: Declination.Value(90.0000001), wantErr: "out of bounds"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s %s %s", test.from.Unit().Name(), test.to.Unit().Name(), test.from)
		t.Run(testName, func(t *testing.T) {
			got, err := test.from.As(test.to.Unit())
			if err != nil {
				if test.wantErr == "" {
					t.Fatalf("unexpected error: %s", err)
				}
				if strings.Contains(err.Error(), test.wantErr) {
					// Stop test here
					return
				}
				t.Fatalf("got error %q want error=%q", err.Error(), test.wantErr)
			} else if test.wantErr != "" {
				t.Fatalf("wanted error %q but got none", test.wantErr)
			}

			if !got.IsValid() {
				t.Fatalf("got invalid value %v", got)
				return
			}

			equals, err := got.Equals(test.to)
			if err != nil {
				if test.wantErr == "" {
					t.Fatalf("unexpected error: %s", err)
				}
				if strings.Contains(err.Error(), test.wantErr) {
					// Stop test here
					return
				}
				t.Fatalf("got error %q want error=%q", err.Error(), test.wantErr)
			} else if test.wantErr != "" {
				t.Fatalf("wanted error %q but got none", test.wantErr)
			}

			if !equals {
				t.Errorf("got: %v want: %v", got, test.to)
			}
		})
	}
}
