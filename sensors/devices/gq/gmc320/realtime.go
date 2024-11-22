package gmc320

import (
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"time"
)

func (i *gmc320) RunDevice(pub publisher.Publisher) error {
	return i.Run(func() error {
		b := make([]byte, 2)

		for {
			n, err := i.Read(b)
			if err != nil {
				return err
			}

			// Only record if we have 2 bytes from the Geiger counter
			if n == 2 {
				rec := i.record(
					time.Now().UTC().Truncate(time.Second),
					toInt16(b)&0x3fff,
				)

				if err := i.publish(pub, rec); err != nil {
					return err
				}
			}

			// Clear buffer in case we have a race condition, or we have invalid data
			// in the stream from the Geiger counter
			err = i.ResetInputBuffer()
			if err != nil {
				return err
			}
		}
	})
}

func (i *gmc320) publish(pub publisher.Publisher, rec cpmReading) error {
	reading := i.NewReading()

	if rec.CPS > 0 {
		reading.SetInt("cps", measurement.CountPerSecond, rec.CPS)
	}

	if rec.CPM > 0 {
		reading.SetInt("cpm", measurement.CountPerMinute, rec.CPM)
	}

	return pub.Do(reading)
}

// toInt16 converts a 2 byte value to an unsigned integer
func toInt16(b []byte) int {
	return (int(b[0]) * 256) + int(b[1])
}
