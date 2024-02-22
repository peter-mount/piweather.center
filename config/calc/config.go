package calc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Script struct {
	Pos          lexer.Position
	Locations    []*location.Location `parser:"(@@)*"`
	Calculations []*Calculation       `parser:"(@@)+"`
	State        *State
}

func (s *Script) merge(b *Script) (*Script, error) {
	if s == nil {
		return b, nil
	}

	if err := s.State.MergeLocations(b.State.MapContainer); err != nil {
		return nil, err
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

// Calculation defines a metric to calculate
type Calculation struct {
	Pos        lexer.Position
	Target     string        `parser:"'calculate' @String"`   // Name of metric to calculate
	At         string        `parser:"('at' @String)?"`       // If set the Location to use
	Every      *time.CronTab `parser:"('every' @@)?"`         // Calculate at specified intervals
	ResetEvery *time.CronTab `parser:"('reset' 'every' @@)?"` // Crontab to reset the value
	Load       *Load         `parser:"(@@)?"`                 // Load from the DB on startup
	UseFirst   *UseFirst     `parser:"(@@)?"`                 // If set and no value use this expression
	Expression *Expression   `parser:"('as' @@)?"`            // Expression to perform calculation
}

type Load struct {
	Pos  lexer.Position
	When string `parser:"'load' @String"` // When to load from
	With string `parser:"'with' @String"` // Query to perform
}

type Expression struct {
	Pos      lexer.Position
	Current  *Current    `parser:"( @@"`   // Get the current value of calculation
	Function *Function   `parser:"| @@"`   // Generic Function Call
	Metric   *Metric     `parser:"| @@ )"` // Metric reference
	Using    *units.Unit `parser:"(@@)?"`  // Optional target Unit
}

// Current returns the current value of the calculation being performed
type Current struct {
	Pos     lexer.Position
	Current bool `parser:"@'current'"`
}

// Function handles function calls
type Function struct {
	Pos         lexer.Position
	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

// Metric handles a metric reference
type Metric struct {
	Pos    lexer.Position
	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

type UseFirst struct {
	Pos    lexer.Position
	Metric *Metric `parser:"'usefirst' @@"`
}
