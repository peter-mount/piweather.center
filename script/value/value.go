package astro

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"time"
)

func init() {
	packages.Register("value", &Value{})
}

type Value struct{}

func (_ Value) PlainTime(t time.Time) value.Time {
	return value.PlainTime(t)
}

func (_ Value) BasicTime(t time.Time, loc *globe.Coord, alt float64) value.Time {
	return value.BasicTime(t, loc, alt)
}
