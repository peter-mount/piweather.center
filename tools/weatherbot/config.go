package bot

type Post struct {
	Name      string    `yaml:"name"`      // Name of the post
	StationId string    `yaml:"stationId"` // StationId for this post
	Threads   []*Thread `yaml:"thread"`    // Threads to generate
}

type Thread struct {
	Prefix string `yaml:"prefix"` // Text to go at the start of the post
	Suffix string `yaml:"suffix"` // Text to go at the end of the post
	Table  []*Row `yaml:"table"`  // Data table
}

type Row struct {
	Format string  `yaml:"format"` // Format for this row
	Values []Value `yaml:"values"` // Values to pass to the format
}

type Value struct {
	Sensor string  `yaml:"sensor"` // Sensor to inject
	Type   string  `yaml:"type"`   // Type of result expected
	Factor float64 `yaml:"factor"` // Factor to apply to value
	Unit   string  `yaml:"unit"`   // Unit ID to use
}

// Value.Type values
const (
	ValueLatest      = "latest"      // Default, latest value
	ValueTrend       = "trend"       // Trend between first and last value in the range
	ValueTime        = "time"        // Time of the latest value. When no sensor the current time
	ValueStationName = "stationName" // Station name
	ValueMin         = "min"         // Min value
	ValueMax         = "max"         // Max value
	ValueMean        = "mean"        // Mean of all values in the range
)
