package calendar

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/util/strings"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/soniakeys/unit"
	"time"
)

func init() {
	packages.RegisterPackage(&Calendar{})
}

type Calendar struct{}

func (_ Calendar) JDNow() julian.Day {
	return julian.Now()
}

func (_ Calendar) StartOfToday() julian.Day {
	return julian.StartOfToday()
}

func (_ Calendar) FromTime(t0 time.Time) julian.Day {
	return julian.FromTime(t0)
}

func (_ Calendar) FromDate(y, m, d, h, min, s int) julian.Day {
	return julian.FromDate(y, m, d, h, min, s)
}

func (_ Calendar) HourDMSString(t unit.Time) string {
	return strings.HourDMSString(t)
}

func (_ Calendar) HourDMSStringExt(d float64) string {
	return strings.HourDMSStringExt(d)
}

func (_ Calendar) DegDMS(d float64) (int, int, float64) {
	return strings.DegDMS(d)
}

func (_ Calendar) DegDMSString(d float64, sign bool) string {
	return strings.DegDMSString(d, sign)
}

func (_ Calendar) DegDMSStringExt(d float64, sign bool, p, m string) string {
	return strings.DegDMSStringExt(d, sign, p, m, 3, 0)
}

func (_ Calendar) LocalMidnight(t time.Time) time.Time {
	return time2.LocalMidnight(t)
}
