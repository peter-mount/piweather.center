package geiger

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"time"
)

type RealtimeStats struct {
}

type CpmReading struct {
	Time time.Time // Time of reading
	CPS  int       // Count per second
	CPM  int       // Count per minute
}

func (m *Geiger) realtime() error {
	b := make([]byte, 2)
	for {
		n, err := m.port.Read(b)
		if err != nil {
			return err
		}
		if *m.Debug {
			log.Printf("Read %d/%d %02x%02x", n, len(b), b[0], b[1])
		}

		// Only record if we have 2 bytes from the Geiger counter
		if n == 2 {
			m.realtimeRecord(CpmReading{
				Time: time.Now().UTC().Truncate(time.Second),
				CPS:  toInt16(b) & 0x3fff,
			})
		}

		// Clear buffer in case we have a race and we have invalid data
		// in the stream from the Geiger counter
		err = m.port.ResetInputBuffer()
		if err != nil {
			return err
		}
	}
}

func (m *Geiger) realtimeRecord(r CpmReading) {
	// Remove entries older than 59 seconds
	cutOff := r.Time.Add(-(time.Minute - time.Second))
	for len(m.realtimeReadings) > 0 && m.realtimeReadings[0].Time.Before(cutOff) {
		m.realtimeReadings = m.realtimeReadings[1:]
	}

	// Stop processing here if no entries or oldest entry is after the cutoff
	cutOff = cutOff.Add(time.Second)
	if len(m.realtimeReadings) > 0 && !m.realtimeReadings[0].Time.After(cutOff) {

		// Calculate CPM as all in recent time including new value
		r.CPM = r.CPS
		for _, e := range m.realtimeReadings {
			r.CPM += e.CPS
		}
	}

	// Add this result once we have finished setting it up as it's pass by value here
	m.realtimeReadings = append(m.realtimeReadings, r)

	// Finally publish the result, based on the report time.
	if (r.CPS > 0 || r.CPM > 0) && (*m.Realtime == 1 || r.Time.Second()%*m.Realtime == 0) {
		// Only record CPS when we have it
		if r.CPS > 0 {
			m.publish("cps", r.Time, float64(r.CPS), "CountPerSecond")
		}

		// Only record CPM when we have it, usually the first minute of operation until we
		// have a minutes worth of data
		if r.CPM > 0 {
			m.publish("cpm", r.Time, float64(r.CPM), "CountPerMinute")
		}
	}
}
