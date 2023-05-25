package station

import (
	"context"
)

// Stations  Map of all defined Station's
type Stations map[string]Station

func StationsFromContext(ctx context.Context) *Stations {
	return ctx.Value("Stations").(*Stations)
}

func (s *Stations) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "Stations", s)
}

func (s *Stations) Accept(v Visitor) error {
	return v.VisitStations(s)
}

// Station defines a Weather Station at a specific location.
// It consists of one or more Reading's
type Station struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Name of the station
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Location of the station
	Location Location `json:"location" xml:"location,omitempty" yaml:"location,omitempty"`
	// One or more Sensors collection
	Sensors map[string]*Sensors `json:"sensors" xml:"sensors" yaml:"sensors"`
}

func StationFromContext(ctx context.Context) *Station {
	return ctx.Value("Station").(*Station)
}

func (s *Station) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "Station", s)
}

func (s *Station) Accept(v Visitor) error {
	return v.VisitStation(s)
}

// Sensors define a Reading collection within the Station.
// A Reading collection is
type Sensors struct {
	ID string `json:"-" xml:"-" yaml:"-"`
	// Name of the Readings collection
	Name string `json:"name" xml:"name,attr" yaml:"name"`
	// Source of data for this collection
	Source Source `json:"source" xml:"source" yaml:"source"`
	// Format of the message, default is json
	Format string
	// Timestamp Path to timestamp, "" for none
	Timestamp string
	// Reading's provided by this collection
	Readings map[string]*Reading `json:"readings" xml:"readings" yaml:"readings"`
	// CalculatedValue's to calculate with this calculation
	Calculations map[string]*CalculatedValue `json:"calculations,omitempty" xml:"calculations,omitempty" yaml:"calculations,omitempty"`
}

func SensorsFromContext(ctx context.Context) *Sensors {
	return ctx.Value("Sensors").(*Sensors)
}

func (s *Sensors) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "Sensors", s), nil
}

func (s *Sensors) Accept(v Visitor) error {
	return v.VisitSensors(s)
}

// ReadingsKeys returns a slice containing the keys for each Reading
func (s *Sensors) ReadingsKeys() []string {
	var keys []string
	for k, _ := range s.Readings {
		keys = append(keys, k)
	}
	return keys
}

// CalculationsKeys returns a slice containing the keys for each Calculations entry
func (s *Sensors) CalculationsKeys() []string {
	var keys []string
	for k, _ := range s.Calculations {
		keys = append(keys, k)
	}
	return keys
}

// GetGraph returns the Graph slice for a specified key.
// It first looks at Readings and if not found then Calculations.
// Returns nil if the key is not found
func (s *Sensors) GetGraph(k string) []*Graph {
	r := s.Readings[k]
	if r != nil {
		return r.Graph
	}

	c := s.Calculations[k]
	if c != nil {
		return c.Graph
	}

	return nil
}
