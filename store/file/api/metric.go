package api

import "time"

type Metric struct {
	Metric string    `json:"metric" xml:"metric,attr"`
	Time   time.Time `json:"time" xml:"time,attr"`
	Unit   string    `json:"unit" xml:"unit,attr"`
	Value  float64   `json:"value" xml:",chardata"`
}

type MetricValue struct {
	Time  time.Time `json:"time" xml:"time,attr"`
	Unit  string    `json:"unit" xml:"unit,attr"`
	Value float64   `json:"value" xml:",chardata"`
}
