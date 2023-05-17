package station

// Config is the root and consists of one or more Station's
type Config struct {
	// Stations one or more Weather Stations supported by this instance
	Stations map[string]*Station `json:"stations" xml:"stations" yaml:"stations"`
}

// Station defines a Weather Station at a specific location.
// It consists of one or more Sensor's
type Station struct {
	// Name of the station
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Location of the station
	Location Location `json:"location" xml:"location,omitempty" yaml:"location,omitempty"`
	// One or more Sensors collection
	Sensors map[string]*Sensors `json:"sensors" xml:"sensors" yaml:"sensors"`
}

// Sensors define a Sensor collection within the Station.
// A Sensor collection is
type Sensors struct {
	// Name of the Sensors collection
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Source of data for this collection
	Source Source `json:"source" xml:"source" yaml:"source"`
	// Format of the message, default is json
	Format string
	// Timestamp Path to timestamp, "" for none
	Timestamp string
	// Sensor's provided by this collection
	Sensors map[string]Sensor `json:"sensors" xml:"sensors" yaml:"sensors"`
}
