package lang

import (
	"encoding/json"
	"encoding/xml"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/go-script/script"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"gopkg.in/yaml.v3"
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
	Pos       lexer.Position
	Name      string `parser:"'amqp' '(' 'name' @String"`
	Url       string `parser:"'url' @String"`
	Exchange  string `parser:"('exchange' @String)? ')'"`
	MQ        *amqp.MQ
	Publisher *amqp.Publisher
}

type Action struct {
	Pos    lexer.Position
	Metric *Metric `parser:"( @@ )"`
}

// Metric on receipt
type Metric struct {
	Pos       lexer.Position
	Metrics   []string           `parser:"'metric' (@String | 'in' '(' @String (',' @String)* ')' )"`
	Statement *script.Statements `parser:"(@@)?"`
	Publish   []*Publish         `parser:"'publish' (@@)+"`
}

type Publish struct {
	Pos     lexer.Position
	Amqp    string       `parser:"( 'amqp' @String"`
	Console bool         `parser:"| @'console' )"`
	As      *PublishType `parser:"(@@)?"`
}

type PublishType struct {
	Pos    lexer.Position
	Bytes  bool `parser:"'as' ( @'bytes'"`
	Json   bool `parser:"| @'json'"`
	String bool `parser:"| @('string' | 'text')"`
	Xml    bool `parser:"| @'xml'"`
	Yaml   bool `parser:"| @'yaml' )"`
}

type Marshaller func(v any) ([]byte, error)

func (p *PublishType) Marshaller() Marshaller {
	switch {
	// p==nil must be first here as it's valid for PublishType to be nil
	case p == nil, p.Bytes, p.String:
		return bytes
	case p.Json:
		return json.Marshal
	case p.Xml:
		return xml.Marshal
	case p.Yaml:
		return yaml.Marshal
	default:
		return func(_ any) ([]byte, error) {
			return nil, participle.Errorf(p.Pos, "invalid type")
		}
	}
}

func bytes(v any) ([]byte, error) {
	v1 := calculator.GetValue(v)
	if b, ok := v1.([]byte); ok {
		return b, nil
	}
	s, err := calculator.GetString(v1)
	if err != nil {
		return nil, err
	}
	return []byte(s), err
}
