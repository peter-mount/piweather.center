package sen0575

import (
	"github.com/peter-mount/piweather.center/sensors"
	"github.com/peter-mount/piweather.center/sensors/bus/i2c"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type sen0575 struct {
	sensors.BasicI2CDevice
	lastReading     time.Time   // time of last reading
	lastBucketCount value.Value // total bucket tips from last reading
}

// Init the device by taking an initial reading first.
func (s *sen0575) Init() error {
	return s.readSensor(s.NewReading())
}

// ReadSensor reads the current values from the device
func (s *sen0575) ReadSensor() (*sensors.Reading, error) {
	ret := s.NewReading()
	return ret, s.readSensor(ret)
}

func (s *sen0575) readSensor(ret *sensors.Reading) error {
	return s.UseDevice(func(bus i2c.I2C) error {
		var err error

		// Set reading time here as we may have been blocked on the device
		ret.Time = time.Now()

		ret.Readings["total"], err = s.getCumulativeRainFall(bus)

		if err == nil {
			ret.Readings["hour"], err = s.getRainFall(bus, 1)
		}

		if err == nil {
			ret.Readings["day"], err = s.getRainFall(bus, 24)
		}

		if err == nil {
			ret.Readings["bucketCount"], err = s.getBucketCount(bus)
		}

		if err == nil {
			// Duration since last reading, accounting for initialisation
			if !s.lastReading.IsZero() {
				ret.Readings["duration"] = value.Integer.Value(ret.Time.Sub(s.lastReading).Seconds())
			}
			s.lastReading = ret.Time

			// Tips since last reading.
			ret.Readings["tips"], err = ret.Readings["bucketCount"].Subtract(s.lastBucketCount)
			s.lastBucketCount = ret.Readings["bucketCount"]
		}

		return err
	})
}
