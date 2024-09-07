package sensors

import (
	"encoding/json"
	"time"
)

type Reading struct {
	// Time of the reading
	Time time.Time

	// ID of the sensor, optional
	ID string

	// Readings from the sensor
	Readings any

	// Device model
	Device string
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
		Time:   time.Now().UTC(),
		Device: dev.Info().Model,
	}
}
