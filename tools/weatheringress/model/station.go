package model

import (
	"context"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/tools/weatheringress/source"
)

// Stations  Map of all defined Station's
type Stations map[string]*Station

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
	ID string `yaml:"-"`
	// Name of the station
	Name string `yaml:"name"`
	// Location of the station
	Location Location `yaml:"location,omitempty"`
	// One or more Sensors collection
	Sensors map[string]*Sensors `yaml:"sensors"`
	latLong *coord.LatLong
}

func (s *Station) LatLong() *coord.LatLong {
	return s.latLong
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
	ID string `yaml:"-"`
	// Name of the Readings collection
	Name string `yaml:"name"`
	// Source of data for this collection
	Source source.Source `yaml:"source"`
	// Format of the message, default is json
	Format string
	// Timestamp Path to timestamp, "" for none
	Timestamp string
	// Reading's provided by this collection
	Readings map[string]*Reading `yaml:"readings"`
	// The station containing this sensor
	station *Station
}

func SensorsFromContext(ctx context.Context) *Sensors {
	return ctx.Value("Sensors").(*Sensors)
}

func (s *Sensors) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "Sensors", s), nil
}

func (s *Sensors) Station() *Station {
	return s.station
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
