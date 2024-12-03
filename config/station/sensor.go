package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type SensorList struct {
	Pos     lexer.Position
	Sensors []*Sensor `parser:"@@*"`
}

type Sensor struct {
	Pos       lexer.Position
	Target    string        `parser:"'sensor' '(' @String"`
	Device    string        `parser:"'driver' '(' @String ')'"`
	I2C       *I2C          `parser:"( @@"`
	Serial    *Serial       `parser:"| @@ )"`
	Poll      *time.CronTab `parser:"('poll' '(' @@ ')')?"`
	Publisher []*Publisher  `parser:"'publish' '(' @@+ ')'"`
}

type Publisher struct {
	Pos lexer.Position
	Log bool `parser:"( @'log'"`
	DB  bool `parser:"| @'db' )"`
}

type I2C struct {
	Pos lexer.Position
	// smbus is a subset of i2c so it's an alias here
	Bus    int `parser:"('i2c'|'smbus') '(' @Int"`
	Device int `parser:"':' @Int ')'"`
}

type Serial struct {
	Pos      lexer.Position
	Port     string `parser:"'serial' '(port)' @String"`
	BaudRate int    `parser:"@Int ')'"`
	//DataBits int    `parser:"('data' @('5'|'6'|'7'|'8'))?"`
	//Parity   string `parser:"('parity' @('no'|'none'|'odd'|'even'))?"`
	//StopBits string `parser:"('stop' @('1'|'1.5'|'2'))?"`
}
