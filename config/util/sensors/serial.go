package sensors

import "github.com/alecthomas/participle/v2/lexer"

type Serial struct {
	Pos      lexer.Position
	Port     string `parser:"'serial' 'port' @String"`
	BaudRate int    `parser:"'baud' @Int"`
	DataBits int    `parser:"('data' @('5'|'6'|'7'|'8'))?"`
	Parity   string `parser:"('parity' @('no'|'none'|'odd'|'even'))?"`
	StopBits string `parser:"('stop' @('1'|'1.5'|'2'))?"`
}
