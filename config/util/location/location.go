package location

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

// Location defines a location on the Earth
type Location struct {
	Pos       lexer.Position
	Name      string         `parser:"'LOCATION' @String"` // Name of location
	Latitude  string         `parser:"@String"`            // Latitude, North positive, South negative
	Longitude string         `parser:"@String"`            // Longitude, East positive, West negative
	Altitude  float64        `parser:"(@Number)?"`         // Altitude in meters. Optional will default to 0
	latLong   *coord.LatLong // Parsed location details
	time      value.Time     // Time based on latLong
}

func (s *Location) LatLong() *coord.LatLong {
	return s.latLong
}

// Time returns a value.Time for this location.
// If the Location is nil then this returns a value.PlainTime
func (s *Location) Time() value.Time {
	if s == nil {
		return value.PlainTime(time.Time{})
	}

	return value.BasicTime(time.Time{}, s.latLong.Coord(), s.Altitude)
}

func (s *Location) Init() error {
	s.Name = strings.ToLower(s.Name)

	lat, err := util.ParseAngle(s.Latitude)
	if err != nil {
		return err
	}

	lon, err := util.ParseAngle(s.Longitude)
	if err != nil {
		return err
	}

	s.latLong = &coord.LatLong{
		Longitude: lon,
		Latitude:  lat,
		Altitude:  s.Altitude,
		Name:      s.Name,
	}

	return nil
}
