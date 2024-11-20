package sensors

import "github.com/alecthomas/participle/v2/lexer"

type I2C struct {
	Pos    lexer.Position
	Bus    int `parser:"'i2c' 'bus' @Int"`
	Device int `parser:"'address' @Int"`
}
