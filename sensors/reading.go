package sensors

import (
	"encoding/json"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// Reading contains the readings from a Device
type Reading struct {
	// Time of the reading
	Time time.Time

	// ID of the sensor, optional
	ID string

	// Readings from the sensor
	Readings map[string]value.Value

	// Device model
	Device string
}

// Get the named value from Reading.Readings
func (r *Reading) Get(n string) value.Value {
	if r.Readings == nil {
		return value.Value{}
	}
	return r.Readings[n]
}

// Set the named value in Reading.Readings
func (r *Reading) Set(n string, v value.Value) {
	// If Time is not set then set it once
	if r.Time.IsZero() {
		r.Time = time.Now()
	}

	if r.Readings == nil {
		r.Readings = make(map[string]value.Value)
	}
	r.Readings[n] = v
}

func (r *Reading) MarshalJSON() ([]byte, error) {
	var buf []byte
	buf = append(buf, `{"time":"`...)
	buf = append(buf, r.Time.Format(time.RFC3339)...)

	if r.ID != "" {
		buf = append(buf, `","id":"`...)
		buf = append(buf, r.ID...)
	}

	buf = append(buf, `","readings":`...)
	if r.Readings == nil {
		buf = append(buf, "null"...)
	} else {
		b1, err := json.Marshal(r.Readings)
		if err != nil {
			return nil, err
		}
		buf = append(buf, b1...)
	}

	buf = append(buf, `,"device":"`...)
	buf = append(buf, r.Device...)

	buf = append(buf, `"}`...)
	return buf, nil
}

func NewReading(dev Device) *Reading {
	return &Reading{
		Time:     time.Now().UTC(), // default to now but usually the device will override it
		Device:   dev.Info().Model,
		Readings: make(map[string]value.Value),
	}
}

type Difference struct {
	// Name of field that is different
	Name string
	// Old Value
	Old value.Value
	// New Value
	New value.Value
	// Time of New Value
	Time time.Time
}

// Differences returns any Difference's between this result and a previous Result.
// This allows us to only submit values when they differ from a previous result
func (r *Reading) Differences(b *Reading) []Difference {
	if r.Time.IsZero() || r.Readings == nil || b == nil || b.Time.IsZero() || b.Readings == nil {
		return nil
	}

	var ret []Difference
	for k, newVal := range r.Readings {
		var different bool

		oldVal, exists := b.Readings[k]

		switch {
		case exists && oldVal.IsValid() && newVal.IsValid():
			// We have an old value so if both are valid then include it
			if notEqual, err := newVal.NotEqual(oldVal); err == nil {
				different = notEqual
			}

		case exists && !oldVal.IsValid():
			// Old value exists but invalid so check new value
			different = newVal.IsValid()

		case newVal.IsValid():
			// Value only in new Result so include if it's valid
			different = true
		}

		if different {
			ret = append(ret, Difference{
				Name: k,
				Old:  oldVal,
				New:  newVal,
				Time: r.Time,
			})
		}
	}

	return ret
}
