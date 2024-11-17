package gmc320

import (
	device2 "github.com/peter-mount/piweather.center/sensors/device"
	"time"
)

type gmc320 struct {
	device2.BasicSerialDevice
	// Readings in realtime mode
	realtimeReadings []CpmReading
}

type CpmReading struct {
	Time time.Time // Time of reading
	CPS  int       // Count per second
	CPM  int       // Count per minute
}

func (i *gmc320) record(t time.Time, cps int) CpmReading {
	r := CpmReading{Time: t, CPS: cps}

	// Remove entries older than 59 seconds
	cutOff := r.Time.Add(-(time.Minute - time.Second))
	for len(i.realtimeReadings) > 0 && i.realtimeReadings[0].Time.Before(cutOff) {
		i.realtimeReadings = i.realtimeReadings[1:]
	}

	// Stop processing here if no entries or oldest entry is after the cutoff
	cutOff = cutOff.Add(time.Second)
	if len(i.realtimeReadings) > 0 && !i.realtimeReadings[0].Time.After(cutOff) {

		// Calculate CPM as all in recent time including new value
		r.CPM = r.CPS
		for _, e := range i.realtimeReadings {
			r.CPM += e.CPS
		}
	}

	// Add this result once we have finished setting it up as it's pass by value here
	i.realtimeReadings = append(i.realtimeReadings, r)

	return r
}
