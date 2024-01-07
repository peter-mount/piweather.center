package lang

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

type Script struct {
	Pos          lexer.Position
	Locations    []*Location    `parser:"(@@)*"`
	Calculations []*Calculation `parser:"(@@)+"`
	State        *State
}

func (s *Script) Accept(v Visitor) error {
	return v.Script(s)
}

func (s *Script) merge(b *Script) (*Script, error) {
	if s == nil {
		return b, nil
	}

	// Merge the state, dealing with id clashes
	for _, l := range b.State.GetLocations() {
		if e := s.State.GetLocation(l.Name); e != nil {
			return nil, participle.Errorf(l.Pos, "location %q already defined at %s", l.Name, e.Pos.String())
		}
		s.State.locations[l.Name] = l
	}

	for _, l := range b.State.GetCalculations() {
		if e := s.State.GetCalculation(l.Target); e != nil {
			return nil, participle.Errorf(l.Pos, "calculation %q already defined at %s", l.Target, e.Pos.String())
		}
		s.State.calculations[l.Target] = l
	}

	// Now merge the slices
	s.Locations = append(s.Locations, b.Locations...)
	s.Calculations = append(s.Calculations, b.Calculations...)

	return s, nil
}

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

func (s *Location) Accept(v Visitor) error {
	return v.Location(s)
}

// Calculation defines a metric to calculate
type Calculation struct {
	Pos        lexer.Position
	Target     string      `parser:"'CALCULATE' @String"`   // Name of metric to calculate
	At         string      `parser:"('AT' @String)?"`       // If set the Location to use
	Every      *CronTab    `parser:"('EVERY' @@)?"`         // Calculate at specified intervals
	ResetEvery *CronTab    `parser:"('RESET' 'EVERY' @@)?"` // Crontab to reset the value
	UseFirst   *UseFirst   `parser:"(@@)?"`                 // If set and no value use this expression
	Expression *Expression `parser:"'AS' @@"`               // Expression to perform calculation
}

func (s *Calculation) Accept(v Visitor) error {
	return v.Calculation(s)
}

type CronTab struct {
	Pos        lexer.Position
	Definition string `parser:"@String"` // CronTab definition
}

func (s *CronTab) Accept(v Visitor) error {
	return v.CronTab(s)
}

type Expression struct {
	Pos      lexer.Position
	Current  *Current  `parser:"( @@"`   // Get the current value of calculation
	Function *Function `parser:"| @@"`   // Generic Function Call
	Metric   *Metric   `parser:"| @@ )"` // Metric reference
	Using    *Unit     `parser:"(@@)?"`  // Optional target Unit
}

func (s *Expression) Accept(v Visitor) error {
	return v.Expression(s)
}

// Unit allows for Unit selection
type Unit struct {
	Pos   lexer.Position
	Using string `parser:"'USING' @String"`
	unit  *value.Unit
}

func (s *Unit) Accept(v Visitor) error {
	return v.Unit(s)
}

func (s *Unit) Unit() *value.Unit {
	return s.unit
}

// Current returns the current value of the calculation being performed
type Current struct {
	Pos     lexer.Position
	Current bool `parser:"@'CURRENT'"`
}

func (s *Current) Accept(v Visitor) error {
	return v.Current(s)
}

// Function handles function calls
type Function struct {
	Pos         lexer.Position
	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

func (s *Function) Accept(v Visitor) error {
	return v.Function(s)
}

// Metric handles a metric reference
type Metric struct {
	Pos    lexer.Position
	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

func (s *Metric) Accept(v Visitor) error {
	return v.Metric(s)
}

type UseFirst struct {
	Pos    lexer.Position
	Metric *Metric `parser:"'USEFIRST' @@"`
}

func (s *UseFirst) Accept(v Visitor) error {
	return v.UseFirst(s)
}
