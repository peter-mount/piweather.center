package util

import (
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// DataSource represents a collection of values to be plotted
type DataSource interface {
	// Size of the DataSource
	Size() int
	// Get a specific entry in the DataSource
	Get(int) (time.Time, value.Value)
	// Period returns the Period of the entries within the DataSource
	Period() time2.Period
	// GetYRange returns the Range of values in the DataSource
	GetYRange() *value.Range
	// GetUnit returns the Unit of the values in the DataSource
	GetUnit() *value.Unit
	// ForEach calls a function for each entry in the DataSource
	ForEach(func(int, time.Time, value.Value))
}
