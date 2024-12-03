package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type CalculationList struct {
	Pos          lexer.Position
	Calculations []*Calculation `parser:"(@@)*"`
}

// Calculation defines a metric to calculate
type Calculation struct {
	Pos        lexer.Position
	Target     string        `parser:"'calculate' '(' @String"` // Name of metric to calculate
	Every      *time.CronTab `parser:"('every' @@)?"`           // Calculate at specified intervals
	ResetEvery *time.CronTab `parser:"('reset' 'every' @@)?"`   // Crontab to reset the value
	Load       *Load         `parser:"(@@)?"`                   // Load from the DB on startup
	UseFirst   *UseFirst     `parser:"(@@)?"`                   // If set and no value use this expression
	Expression *Expression   `parser:"('as' @@) ')'"`           // Expression to perform calculation
}

type Load struct {
	Pos  lexer.Position
	When string `parser:"'load' @String"` // When to load from
	With string `parser:"'with' @String"` // Query to perform
}

type Expression struct {
	Pos      lexer.Position
	Current  *Current            `parser:"( @@"`   // Get the current value of calculation
	Function *Function           `parser:"| @@"`   // Generic Function Call
	Location *LocationExpression `parser:"| @@"`   // Return values from the stations location
	Metric   *Metric             `parser:"| @@ )"` // Metric reference
	Using    *units.Unit         `parser:"(@@)?"`  // Optional target Unit
}

type LocationExpression struct {
	Pos       lexer.Position
	Latitude  bool `parser:"( @'latitude'"`
	Longitude bool `parser:"| @'longitude'"`
	Altitude  bool `parser:"| @'altitude' )"`
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

type UseFirst struct {
	Pos    lexer.Position
	Metric *Metric `parser:"'usefirst' @@"`
}
