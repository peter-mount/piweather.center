package station

// Location represents a location on the Earth's surface.
type Location struct {
	// Name of the location, optional.
	Name string `yaml:"name,omitempty"`
	// Latitude of the Location.
	// This can be in decimal degrees, or in dd:mm.mm or dd:mm:ss formats.
	// North is positive.
	Latitude string `yaml:"latitude"`
	// Longitude of the Location.
	// This can be in decimal degrees, or in dd:mm.mm or dd:mm:ss formats.
	// East is positive, West negative.
	Longitude string `yaml:"longitude"`
	// Altitude of the Location in meters.
	Altitude float64 `yaml:"altitude,omitempty"`
	// Notes of the location, optional.
	Notes string `yaml:"notes,omitempty"`
}
