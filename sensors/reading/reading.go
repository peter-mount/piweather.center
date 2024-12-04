package reading

import (
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

	keys []string

	// Device model
	Device string

	// Source of the message, only used by http
	Source string
}

// IsEmpty returns true if the Reading is nil, it has no Time set, or it has no readings
func (r *Reading) IsEmpty() bool {
	return r == nil || len(r.Readings) == 0 || r.Time.IsZero()
}

// Get the named value from Reading.Readings
func (r *Reading) Get(n string) value.Value {
	if r.Readings == nil {
		return value.Value{}
	}
	return r.Readings[n]
}

// Set the named Value
func (r *Reading) Set(n string, v value.Value) {
	// If Time is not set then set it once
	if r.Time.IsZero() {
		r.Time = time.Now()
	}

	if r.Readings == nil {
		r.Readings = make(map[string]value.Value)
	}

	if _, exists := r.Readings[n]; !exists {
		r.keys = append(r.keys, n)
	}
	r.Readings[n] = v
}

// SetInt sets the named value
func (r *Reading) SetInt(n string, unit *value.Unit, v int) {
	r.SetFloat64(n, unit, float64(v))
}

// SetFloat64 sets the named value
func (r *Reading) SetFloat64(n string, unit *value.Unit, v float64) {
	r.Set(n, unit.Value(v))
}

func (r *Reading) MarshalJSON() ([]byte, error) {
	var buf []byte
	buf = append(buf, `{"time":"`...)
	buf = append(buf, r.Time.Format(time.RFC3339)...)
	buf = append(buf, '"')

	if r.ID != "" {
		buf = append(buf, `,"id":"`...)
		buf = append(buf, r.ID...)
		buf = append(buf, '"')
	}

	buf = r.marshallMap(buf, "readings", func(v value.Value) string {
		return v.PlainString()
	})

	buf = r.marshallMap(buf, "units", func(v value.Value) string {
		return `"` + v.Unit().ID() + `"`
	})

	buf = append(buf, `,"device":"`...)
	buf = append(buf, r.Device...)

	buf = append(buf, `"}`...)
	return buf, nil
}

func (r *Reading) marshallMap(buf []byte, name string, f func(k value.Value) string) []byte {
	buf = append(buf, `,"`...)
	buf = append(buf, name...)
	buf = append(buf, `":`...)

	if r.Readings == nil {
		return append(buf, "null"...)
	}

	buf = append(buf, '{')

	sep := false
	for _, k := range r.keys {
		if v, exists := r.Readings[k]; exists {
			if sep {
				buf = append(buf, ',')
			} else {
				sep = true
			}
			buf = append(buf, '"')
			buf = append(buf, k...)
			buf = append(buf, '"', ':')
			buf = append(buf, f(v)...)
		}
	}

	return append(buf, '}')
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
	if r.IsEmpty() || b.IsEmpty() {
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
