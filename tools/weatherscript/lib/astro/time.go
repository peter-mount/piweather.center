package astro

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/soniakeys/unit"
	"time"
)

type Time struct{}

func (_ Time) JDNow() julian.Day {
	return julian.Now()
}

func (_ Time) StartOfToday() julian.Day {
	return julian.StartOfToday()
}

func (_ Time) FromTime(t0 time.Time) julian.Day {
	return julian.FromTime(t0)
}

func (_ Time) FromDate(y, m, d, h, min, s int) julian.Day {
	return julian.FromDate(y, m, d, h, min, s)
}

func (_ Time) HourDMSString(t unit.Time) string {
	return util.HourDMSString(t)
}

func (_ Time) HourDMSStringExt(d float64) string {
	return util.HourDMSStringExt(d)
}

func (_ Time) DegDMS(d float64) (int, int, float64) {
	return util.DegDMS(d)
}

func (_ Time) DegDMSString(d float64, sign bool) string {
	return util.DegDMSString(d, sign)
}

func (_ Time) DegDMSStringExt(d float64, sign bool, p, m string) string {
	return util.DegDMSStringExt(d, sign, p, m)
}

func (_ Time) LocalMidnight(t time.Time) time.Time {
	return time2.LocalMidnight(t)
}
