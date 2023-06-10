package bot

import (
	"github.com/peter-mount/piweather.center/weather/value"
)

// Post represents the weatherbot config for a specific weather.Station
type Post struct {
	Name      string    `yaml:"name"`      // Name of the post
	StationId string    `yaml:"stationId"` // StationId for this post
	Threads   []*Thread `yaml:"thread"`    // Threads to generate
}

// Thread defines the format of an actual post.
// You can have multiple Thread's in a Post, they will appear as a
// thread in Mastodon.
type Thread struct {
	Prefix string `yaml:"prefix"` // Text to go at the start of the post
	Suffix string `yaml:"suffix"` // Text to go at the end of the post
	Table  []*Row `yaml:"table"`  // Data table
}

// Row of data within a Thread
type Row struct {
	Format string  `yaml:"format"`         // Format for this row
	Values []Value `yaml:"values"`         // Values to pass to the format
	When   []When  `yaml:"when,omitempty"` // Define conditions to show this Row
}

type When struct {
	Value            Value  `yaml:"value"`
	LessThan         *Value `yaml:"lessThan"`
	LessThanEqual    *Value `yaml:"lessThanEqual"`
	Equal            *Value `yaml:"equal"`
	NotEqual         *Value `yaml:"notEqual"`
	GreaterThanEqual *Value `yaml:"greaterThanEqual"`
	GreaterThan      *Value `yaml:"greaterThan"`
}

// Value in a Row that will provide data for the Row formatter
type Value struct {
	Sensor string   `yaml:"sensor"`          // Sensor to inject
	Type   string   `yaml:"type"`            // Type of result expected
	Factor float64  `yaml:"factor"`          // Factor to apply to value
	Unit   Unit     `yaml:"unit"`            // Units to use
	Range  string   `yaml:"range,omitempty"` // The state.Value range to use, default current10
	Value  *float64 `yaml:"value"`
}

func (v Value) GetValue(src value.Value) (value.Value, error) {
	if dest, destOk := value.GetUnit(v.Unit.Unit); destOk {
		v1, err := src.As(dest)
		if err == nil {
			// If we have an alternate now convert to it
			altUnit := v.GetUnit(v1.Float())
			if altUnit != dest.Unit() {
				if alt, altOk := value.GetUnit(altUnit); altOk {
					return v1.As(alt)
				}
			}
		}
		return v1, err
	}
	return src, nil
}

// GetUnit returns the appropriate unit for f.
// This allows us to use different units based on f.
// e.g. "lux" is the default, but use "kiloLux" if >= 1000 lux
func (v Value) GetUnit(f float64) string {
	found := v.Unit.Unit

	for _, u := range v.Unit.Alternate {
		minSet, maxSet := !value.IsZero(u.Min), !value.IsZero(u.Max)

		switch {

		// min & max set and min < max
		case minSet && maxSet && value.LessThan(u.Min, u.Max):
			if value.Within(f, u.Min, u.Max) {
				found = u.Unit
			}

		// min & max set and min > max
		case minSet && maxSet && value.GreaterThan(u.Min, u.Max):
			if value.Without(f, u.Min, u.Max) {
				found = u.Unit
			}

		case minSet:
			if value.GreaterThanEqual(f, u.Min) {
				found = u.Unit
			}

		case maxSet:
			if value.LessThanEqual(f, u.Max) {
				found = u.Unit
			}

		}
	}

	return found
}

// Unit conversion
type Unit struct {
	Unit      string    `yaml:"unit"`                // Unit ID to use
	Alternate []SubUnit `yaml:"alternate,omitempty"` // Alternates
}
type SubUnit struct {
	Unit string  `yaml:"unit"`          // Unit ID to use
	Min  float64 `yaml:"min,omitempty"` // Minimum value to use. If min==max==0 then this is the default
	Max  float64 `yaml:"max,omitempty"` // Maximum value to use
}

// Value.Type values
const (
	ValueLatest      = "latest"      // Default, latest value
	ValuePrevious    = "previous"    // latest value in previous10
	ValueTrend       = "trend"       // Trend between first and last value in the range
	ValueTime        = "time"        // Time of the latest value. When no sensor the current time
	ValueStationName = "stationName" // Station name
	ValueMin         = "min"         // Min value
	ValueMax         = "max"         // Max value
	ValueMean        = "mean"        // Mean of all values in the range
	ValueTotal       = "total"       // The total of all values in the range
	ValueCount       = "count"       // The number of entries in the range
)

// Measurement range selections
const (
	RangeCurrent  = "current"  // default, use current10
	RangePrevious = "previous" // use previous10
	RangeHour     = "hour"     // use hour
	RangeHour24   = "hour24"   // use hour24
	RangeToday    = "today"    // use today
)
