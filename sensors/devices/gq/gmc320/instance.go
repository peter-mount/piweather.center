package gmc320

import (
	device2 "github.com/peter-mount/piweather.center/sensors/device"
	"time"
)

type gmc320 struct {
	device2.BasicSerialDevice
	// Readings in realtime mode
	realtimeReadings []cpmReading
	// The last CPS value, used to reduce the number of entries published
	lastCps int
}

type cpmReading struct {
	Time time.Time // Time of reading
	CPS  int       // Count per second
	CPM  int       // Count per minute
}

func (i *gmc320) record(t time.Time, cps int) cpmReading {
	// Truncate t to the Second as that's the DB resolution.
	t = t.Truncate(time.Second)

	// if the time is the same as the latest entry then modify that entry to match.
	// Required as the counter can issue entries not quite on the second, and we would lose data in the db
	// so it's safer to aggregate
	if len(i.realtimeReadings) > 0 {
		l := len(i.realtimeReadings) - 1
		r := i.realtimeReadings[l]
		if r.Time.Equal(t) {
			// add cps to the existing entry
			r.CPS = r.CPS + cps

			// Only update CPM if we have a value - to account for entries before we have enough data to calculate it
			if r.CPM > 0 {
				r.CPM = r.CPM + cps
			}

			// Update slice as its pass by value here
			i.realtimeReadings[l] = r

			return r
		}
	}

	r := cpmReading{Time: t, CPS: cps}

	// Remove entries older than 59 seconds, or we have more than 59 readings.
	cutOff := r.Time.Add(-(time.Minute - time.Second))
	for len(i.realtimeReadings) > 0 && (i.realtimeReadings[0].Time.Before(cutOff) || len(i.realtimeReadings) > 59) {
		i.realtimeReadings = i.realtimeReadings[1:]
	}

	// Only calculate CPM if we have 59 entries (as we don't add 60th entry until we have a CPM)
	if len(i.realtimeReadings) == 59 {
		r.CPM = r.CPS
		for _, e := range i.realtimeReadings {
			r.CPM += e.CPS
		}
	}

	// Add this result once we have finished setting it up as it's pass by value here
	i.realtimeReadings = append(i.realtimeReadings, r)

	return r
}
