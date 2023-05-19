package station

import (
	"context"
)

// Stations  Map of all defined Station's
type Stations map[string]Station

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

func (s *Station) Accept(v Visitor) error {
	return v.VisitStation(s)
}

func InitStation(ctx context.Context) error {
	s := StationFromContext(ctx)
	s.ID = ctx.Value("StationId").(string)
	return nil
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
}

func (s *Sensors) Accept(v Visitor) error {
	return v.VisitSensors(s)
}

func InitSensors(ctx context.Context) error {
	s := SensorsFromContext(ctx)

	// Set the VisitStation.ID
	stationConfig := StationFromContext(ctx)
	s.ID = stationConfig.ID + "." + ctx.Value("SensorId").(string)

	return nil
}
