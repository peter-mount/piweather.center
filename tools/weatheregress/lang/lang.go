package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Script struct {
	Pos lexer.Position
	//Locations []*lang.Location `parser:"(@@)*"` // Share the same from calculator
	Amqp    []*Amqp   `parser:"(@@)?"`
	Actions []*Action `parser:"(@@)+"`
	state   *State
}

func (s *Script) State() *State {
	return s.state
}

func (s *Script) merge(b *Script) (*Script, error) {
	if s == nil {
		return b, nil
	}

	s.state.merge(b.state)

	return s, nil
}

// Amqp broker definition
type Amqp struct {
	Pos      lexer.Position
	Name     string `parser:"'AMQP' '(' 'NAME' @String"`
	Url      string `parser:"'URL' @String"`
	Exchange string `parser:"('EXCHANGE' @String)? ')'"`
}

type Action struct {
	Pos    lexer.Position
	Metric *Metric `parser:"( @@ )"`
}

// Metric on receipt
type Metric struct {
	Pos     lexer.Position
	Metrics []string   `parser:"'METRIC' (@String | 'IN' '(' @String (',' @String)* ')' )"`
	Format  *Format    `parser:"(@@)"`
	Publish []*Publish `parser:"'PUBLISH' (@@)+"`
}

type Publish struct {
	Pos     lexer.Position
	Amqp    string `parser:"( 'AMQP' @String"`
	Console bool   `parser:"| @'CONSOLE' )"`
}

type Format struct {
	Pos         lexer.Position
	Format      string              `parser:"'FORMAT' '(' @String"`
	Expressions []*FormatExpression `parser:"(',' @@)* ')'"`
}

type FormatExpression struct {
	Pos   lexer.Position
	Left  *FormatAtom       `parser:"@@"`
	Op    string            `parser:"( (@'+' | @'-')"`
	Right *FormatExpression `parser:"@@ )?"`
}

type FormatAtom struct {
	Pos      lexer.Position
	Metric   bool    `parser:"( @'METRIC'"`
	Value    bool    `parser:"| @'VALUE'"`
	UnixTime bool    `parser:"| @'UNIXTIME'"`
	String   *string `parser:"| @String )"`
}
