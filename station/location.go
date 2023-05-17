package station

// Location represents a location on the Earth's surface.
type Location struct {
	// Name of the location, optional.
	Name string `json:"name" xml:"name,attr,omitempty" yaml:"name,omitempty"`
	// Latitude of the Location.
	// This can be in decimal degrees, or in dd:mm.mm or dd:mm:ss formats.
	// North is positive.
	Latitude string `json:"latitude" xml:"latitude,attr" yaml:"latitude"`
	// Longitude of the Location.
	// This can be in decimal degrees, or in dd:mm.mm or dd:mm:ss formats.
	// East is positive, West negative.
	Longitude string `json:"longitude" xml:"longitude,attr" yaml:"longitude"`
	// Altitude of the Location in meters.
	Altitude float64 `json:"altitude" xml:"altitude,attr,omitempty" yaml:"altitude,omitempty"`
	// Notes of the location, optional.
	Notes string `json:"notes" xml:"notes,omitempty" yaml:"notes,omitempty"`
}
