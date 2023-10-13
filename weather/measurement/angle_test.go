package measurement

import (
	"math"
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
