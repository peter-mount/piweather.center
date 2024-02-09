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

func (s *Script) Accept(v Visitor) error {
	return v.Script(s)
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

func (s *Amqp) Accept(v Visitor) error {
	return v.Amqp(s)
}

type Action struct {
	Pos    lexer.Position
	Metric *Metric `parser:"( @@ )"`
}

func (s *Action) Accept(v Visitor) error {
	return v.Action(s)
}

// Metric on receipt
type Metric struct {
	Pos     lexer.Position
	Metrics []string `parser:"'METRIC' (@String | 'IN' '(' @String (',' @String)* ')' )"`
	Format  *Format  `parser:"(@@)"`
	Amqp    string   `parser:"'PUBLISH' ( 'AMQP' @String )"`
}

func (s *Metric) Accept(v Visitor) error {
	return v.Metric(s)
}

type Format struct {
	Pos         lexer.Position
	Format      string              `parser:"'FORMAT' '(' @String ','"`
	Expressions []*FormatExpression `parser:"( @@ (',' @@)* )? ')'"`
}

type FormatExpression struct {
	Pos      lexer.Position
	Metric   bool   `parser:"( @'METRIC'"`
	Value    bool   `parser:"| @'VALUE'"`
	UnixTime bool   `parser:"| @'UNIXTIME'"`
	String   string `parser:"| @String )"`
}
