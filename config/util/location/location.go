package location

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/astro/coord"
	strings2 "github.com/peter-mount/piweather.center/util/strings"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

// Location defines a location on the Earth
type Location struct {
	Pos       lexer.Position
	Name      string         `parser:"'location' '(' @String"`          // Name of location
	Latitude  string         `parser:"@String"`                         // Latitude, North positive, South negative
	Longitude string         `parser:"@String"`                         // Longitude, East positive, West negative
	Altitude  float64        `parser:"(@Number)?"`                      // Altitude in meters. Optional will default to 0
	Notes     string         `parser:"(('note'|'notes') @String)? ')'"` // Optional note
	latLong   *coord.LatLong // Parsed location details
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
	var err error

	s.Name = strings.ToLower(s.Name)

	s.latLong = &coord.LatLong{
		Altitude: s.Altitude,
		Name:     s.Name,
		Notes:    s.Notes,
	}

	s.latLong.Latitude, err = strings2.ParseAngle(s.Latitude)

	if err == nil {
		s.latLong.Longitude, err = strings2.ParseAngle(s.Longitude)
	}

	return errors.Error(s.Pos, err)
}
