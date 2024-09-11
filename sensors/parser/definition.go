package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type SensorDefinition struct {
	Pos     lexer.Position
	Sensors []Sensor `parser:"@@+"`
}

type Sensor struct {
	Pos      lexer.Position
	ID       string    `parser:"'sensor' '(' @String"`
	I2C      *I2C      `parser:"( @@"`
	Serial   *Serial   `parser:"| @@ )"`
	Device   string    `parser:"'device' '(' @String ')'"`
	Cron     string    `parser:"'cron' '(' @String ')'"`
	Readings []Reading `parser:"@@+ ')'"`
}

type I2C struct {
	Pos    lexer.Position
	Bus    int `parser:"'i2c' '(' @Number ','"`
	Device int `parser:"@Number ')'"`
}

type Serial struct {
	Pos    lexer.Position
	Device string `parser:"'serial' '(' @String ')'"`
}

type Reading struct {
	Pos    lexer.Position
	Source string `parser:"'reading' '(' @String ','"`
	Unit   string `parser:"@String ')'"`
	As     string `parser:"( 'as' @String ')' )?"`
}
