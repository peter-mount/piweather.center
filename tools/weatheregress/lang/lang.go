package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/script"
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
	Name     string `parser:"'amqp' '(' 'name' @String"`
	Url      string `parser:"'url' @String"`
	Exchange string `parser:"('exchange' @String)? ')'"`
}

type Action struct {
	Pos    lexer.Position
	Metric *Metric `parser:"( @@ )"`
}

// Metric on receipt
type Metric struct {
	Pos        lexer.Position
	Metrics    []string    `parser:"'metric' (@String | 'in' '(' @String (',' @String)* ')' )"`
	Expression *Expression `parser:"('eval' '(' @@ ')')?"`
	Publish    []*Publish  `parser:"'publish' (@@)+"`
}

type Expression struct {
	Pos        lexer.Position
	Expression []*script.Expression `parser:"(@@)+"`
}

type Publish struct {
	Pos     lexer.Position
	Amqp    string `parser:"( 'amqp' @String"`
	Console bool   `parser:"| @'console' )"`
}
