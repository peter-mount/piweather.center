package state

import (
	"math"
	"strconv"
	"time"
)

type RoundedFloat float64

func (rf RoundedFloat) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(rf), 'f', 3, 64)), nil
}

// Station represents the state of a weather station, including general
// statistics, trends etc.
//
// It's exposed by the server and available for external tools to use.
// e.g. the Mastodon bot uses this.
type Station struct {
	// The ID of the station, taken from station.Station
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	// The Measurement's from the station
	Measurements []*Measurement `json:"measurements"`
}

func (s *Station) Clone() *Station {
	return &Station{
		ID:   s.ID,
		Meta: s.Meta,
	}
}

func (s *Station) AddMeasurement(m *Measurement) {
	if m != nil {
		s.Measurements = append(s.Measurements, m)
		if m.Time.After(s.Meta.Time) {
			s.Meta.Time = m.Time
		}
	}
}

type Meta struct {
	// The name of the station, taken from station.Station
	Name string `json:"name"`
	// Time this data was compiled
	Time time.Time `json:"time"`
	// Start of Minute10 Value's
	Minute10 time.Time `json:"minute_10"`
	// Start of Hour Value's
	Hour time.Time `json:"hour"`
	// Start of Today Value's
	Today time.Time `json:"today"`
	// Start of Hour24 Value's
	Hour24 time.Time `json:"hour_24"`
	// The units used
	Units map[string]Unit `json:"units"`
}

type Unit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

// Measurement representing a reading or calculated value from the station
type Measurement struct {
	// The ID name
	ID string `json:"id"`
	// The time of this entry
	Time time.Time `json:"time"`
	// The unit of this Measurement
	Unit string `json:"unit"`
	// The last 10 minutes
	Minute10 Value `json:"minute_10"`
	// The last hour
	Hour Value `json:"hour"`
	// The current day since 00:00 local time
	Today Value `json:"today"`
	// The Last 24 hours
	Hour24 Value `json:"hour_24"`
	// Trend of this ID
	Trend Trend `json:"trend"`
}

// Value represents the state for a given unit of time
type Value struct {
	// Latest value
	Latest RoundedFloat `json:"latest"`
	// Minimum for this period
	Min RoundedFloat `json:"min"`
	// Maximum for this period
	Max RoundedFloat `json:"max"`
	// Mean for this period
	Mean RoundedFloat `json:"mean"`
	// Total for this period
	Total RoundedFloat `json:"total"`
	// Count the number of entries in this period, used for Total and Mean
	Count int `json:"count"`
}

// Include includes a new value
func (v Value) Include(val float64) Value {
	if v.Count == 0 {
		return Value{
			Latest: RoundedFloat(val),
			Min:    RoundedFloat(val),
			Max:    RoundedFloat(val),
			Mean:   RoundedFloat(val),
			Total:  RoundedFloat(val),
			Count:  1,
		}
	}

	nv := Value{
		Latest: RoundedFloat(val),
		Min:    RoundedFloat(math.Min(float64(v.Min), val)),
		Max:    RoundedFloat(math.Max(float64(v.Max), val)),
		Total:  RoundedFloat(float64(v.Total) + val),
		Count:  v.Count + 1,
	}
	nv.Mean = RoundedFloat(float64(nv.Total) / float64(nv.Count))
	return nv
}

type Trend struct {
	// Trend of a measurement
	Trend string `json:"trend"`
	// Trend as a unicode character
	TrendUnicode string `json:"trend_unicode"`
}
