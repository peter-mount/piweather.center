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
	Target    *Metric       `parser:"'sensor' '(' @@"`
	I2C       *I2C          `parser:"( @@"`
	Serial    *Serial       `parser:"| @@ )"`
	Poll      *time.CronTab `parser:"('poll' '(' @@ ')')?"`
	Publisher []*Publisher  `parser:"'publish' '(' @@+ ')' ')'"`
}

type Publisher struct {
	Pos lexer.Position
	Log bool `parser:"( @'log'"`
	DB  bool `parser:"| @'db' )"`
}

type I2C struct {
	Pos lexer.Position
	// smbus is a subset of i2c so it's an alias here
	Driver string `parser:"('i2c'|'smbus') '(' @String"` // device driver id
	Bus    int    `parser:"    @Number"`                 // i2c bus id in the OS kernel
	Device int    `parser:"':' @Number ')'"`             // i2c address on the specific bus
}

type Serial struct {
	Pos      lexer.Position
	Driver   string `parser:"'serial' '(' @String"` // device driver id
	Port     string `parser:" @String"`             // serial port
	BaudRate int    `parser:" @Number ')'"`         // Baud rate
	//DataBits int    `parser:"('data' @('5'|'6'|'7'|'8'))?"`
	//Parity   string `parser:"('parity' @('no'|'none'|'odd'|'even'))?"`
	//StopBits string `parser:"('stop' @('1'|'1.5'|'2'))?"`
}
