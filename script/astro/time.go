package astro

import (
	"github.com/peter-mount/piweather.center/astro/julian"
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
	return julian.FromDate(y, m, d, h, m, s)
}
