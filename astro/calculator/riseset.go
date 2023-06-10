package calculator

import (
	"fmt"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"time"
)

// RiseSet represents the time an object rises or sets on the horizon.
// It also records the time of transit (highest altitude)
type RiseSet struct {
	JD           julian.Day
	Location     *globe.Coord
	Rise         AltAz
	Set          AltAz
	UpperTransit AltAz
	LowerTransit AltAz
	Duration     time.Duration
}

// NotVisible is true when the object is never above the horizon
func (r RiseSet) NotVisible() bool {
	return !r.UpperTransit.IsVisible()
}

// Circumpolar is true if the object is always above the horizon
func (r RiseSet) Circumpolar() bool {
	return r.UpperTransit.IsVisible() && r.LowerTransit.IsVisible()
}

// AltAz represents the position of an object in the local sky at a specific time
type AltAz struct {
	Time     time.Time
	Altitude value.Value
	Azimuth  value.Value
}

func (a AltAz) IsValid() bool {
	return !a.Time.IsZero() || (a.Azimuth.IsValid() && a.Altitude.IsValid())
}

func (a AltAz) IsVisible() bool {
	return a.IsValid() && a.Altitude.IsValid() && value.GreaterThan(a.Altitude.Float(), 0)
}

func (a AltAz) String() string {
	if a.IsValid() {
		return fmt.Sprintf("%s %s %s", a.Time.Format(time.RFC3339), a.Altitude, a.Azimuth)
	}
	return "N/A"
}
