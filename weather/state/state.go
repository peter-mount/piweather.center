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
	Minute10 time.Time `json:"minute10"`
	// Start of Previous10 Value's
	Previous10 time.Time `json:"previous10"`
	// Start of Hour Value's
	Hour time.Time `json:"hour"`
	// Start of Today Value's
	Today time.Time `json:"today"`
	// Start of Hour24 Value's
	Hour24 time.Time `json:"hour24"`
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
	// Current is the latest value in last 10 minutes represented by Current10
	Current Point `json:"current"`
	// The last 10 minutes
	Current10 Value `json:"current10"`
	// Previous is the latest value in the previous 10 minutes represented by Previous10
	Previous Point `json:"previous"`
	// The previous 10 minutes
	Previous10 Value `json:"previous10"`
	// The last hour
	Hour Value `json:"hour"`
	// The current day since 00:00 local time
	Today Value `json:"today"`
	// The Last 24 hours
	Hour24 Value `json:"hour24"`
	// Trend of this ID
	Trends Trends `json:"trend"`
}

type Point struct {
	Value RoundedFloat `json:"value"`
	Time  time.Time    `json:"time"`
}

// Value represents the state for a given unit of time
type Value struct {
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
			Min:   RoundedFloat(val),
			Max:   RoundedFloat(val),
			Mean:  RoundedFloat(val),
			Total: RoundedFloat(val),
			Count: 1,
		}
	}

	nv := Value{
		Min:   RoundedFloat(math.Min(float64(v.Min), val)),
		Max:   RoundedFloat(math.Max(float64(v.Max), val)),
		Total: RoundedFloat(float64(v.Total) + val),
		Count: v.Count + 1,
	}
	nv.Mean = RoundedFloat(float64(nv.Total) / float64(nv.Count))
	return nv
}

type Trends struct {
	// Time trend is from
	From time.Time `json:"from"`
	// Time trend is to
	To time.Time `json:"to"`
	// Current Trend
	Current Trend `json:"current"`
	// Minimum Trend
	Min Trend `json:"min"`
	// Maximum Trend
	Max Trend `json:"max"`
	// Mean Trend
	Mean Trend `json:"mean"`
}

type Trend struct {
	// Trend of a measurement
	Trend string `json:"trend"`
	// Char is Trend as a unicode character
	Char string `json:"char"`
}

const (
	NoData       = ""
	Falling      = "Falling"
	Rising       = "Rising"
	Stable       = "Stable"
	TrendNoData  = ""
	TrendFalling = "↓"
	TrendRising  = "↑"
	TrendStable  = "→"
)

// TrendFrom returns the Trend based on a previous and current value
func TrendFrom(previous, current float64) Trend {
	switch {
	case current < previous:
		return Trend{
			Trend: Falling,
			Char:  TrendFalling,
		}
	case current > previous:
		return Trend{
			Trend: Rising,
			Char:  TrendRising,
		}
	default:
		return Trend{
			Trend: Stable,
			Char:  TrendStable,
		}
	}
}
