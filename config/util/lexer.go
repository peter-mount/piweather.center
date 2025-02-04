package util

import (
	"crypto/sha1"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"io"
	"os"
)

var (
	commonLexerRules = []lexer.SimpleRule{
		{"hashComment", `#.*`},
		{"sheBang", `#\!.*`},
		{"comment", `//.*|/\*.*?\*/`},
		{"Whitespace", `\s+`},
		//{"Ident", `([a-zA-Z_][a-zA-Z0-9_]*)`},
		{"Ident", `\b([a-zA-Z_][a-zA-Z0-9_]*)\b`},
		//{"Ident", `\b(([a-zA-Z_][a-zA-Z0-9_]*)(\.([a-zA-Z_][a-zA-Z0-9_]*))*)\b`},
		//{"Punct", `[-,()*/+%{};&!=:<>\|]|\[|\]|\^`},
		{"Punct", `[-,()*/+%{};&!=:<>~\|\\]|\[|\]|\^`},
		{"Number", `[-+]?(\d*\.)?\d+`},
		//{"Number", `(-?\d+(\.\d+)?)`},
		//{"Number", `[-+]?(\d+\.\d+|\.\d+|\d+)`},
		//{"Number", `[-+]?((\d*)?\.\d+|\d+\.(\d*)?)`},
		{"Int", `[-+]?\d+`},
		{"String", `"(\\"|[^"])*"`},
		{"Period", `(\.)`},
		{"NewLine", `[\n\r]+`},
		{"Comma", `,`},
		{"Query", `\?`},
		//{"Select", `"SELECT"`},
	}
)

// Parser allows a full Query to be parsed with it fully initialised
type Parser[G any] interface {
	Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*G, error)
	ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*G, error)
	ParseString(fileName, src string, opts ...participle.ParseOption) (*G, error)
	ParseFile(fileName string, opts ...participle.ParseOption) (*G, error)
	ParseFiles(fileNames ...string) (*G, error)
}

type ParserInit[G any] func(p Parser[G], q *G, err error) (*G, error)

type defaultParser[G any] struct {
	lexer  *lexer.StatefulDefinition
	parser *participle.Parser[G]
	init   ParserInit[G]
}

func NewParser[G any](rules []lexer.SimpleRule, opts []participle.Option, init ParserInit[G]) Parser[G] {
	return NewParserExt[G](rules, func(_ lexer.Definition) []participle.Option {
		return opts
	}, init)
}

func NewParserExt[G any](rules []lexer.SimpleRule, optFactory func(l lexer.Definition) []participle.Option, init ParserInit[G]) Parser[G] {
	var r []lexer.SimpleRule
	if len(rules) == 0 {
		r = commonLexerRules
	} else {
		r = append(r, rules...)
		r = append(r, commonLexerRules...)
	}
	l := lexer.MustSimple(append(rules, commonLexerRules...))

	o := []participle.Option{
		participle.Lexer(l),
		participle.UseLookahead(2),
		participle.Unquote("String"),
		participle.Elide("Whitespace"),
	}
	if optFactory != nil {
		o = append(o, optFactory(l)...)
	}
	p := participle.MustBuild[G](o...)

	if init == nil {
		init = defaultInit[G]
	}

	return &defaultParser[G]{
		lexer:  l,
		parser: p,
		init:   init,
	}
}

func defaultInit[G any](_ Parser[G], q *G, err error) (*G, error) {
	return q, err
}

func (p *defaultParser[G]) Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*G, error) {
	v, err := p.parser.Parse(fileName, r, opts...)
	return p.init(p, v, err)
}

// CheckSummable is implemented by parsable entities to store a checksum of the parsed content
type CheckSummable interface {
	GetChecksum() [20]byte
	SetChecksum([20]byte)
}

// CheckSum implements CheckSummable and stores a checksum of a parsed entity
type CheckSum struct {
	checksum [20]byte
}

func (c *CheckSum) GetChecksum() [20]byte {
	return c.checksum
}

func (c *CheckSum) SetChecksum(checksum [20]byte) {
	c.checksum = checksum
}

func (p *defaultParser[G]) ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*G, error) {
	return p.init(p.parseBytes(fileName, b, opts...))
}

func (p *defaultParser[G]) parseBytes(fileName string, b []byte, opts ...participle.ParseOption) (Parser[G], *G, error) {
	v, err := p.parser.ParseBytes(fileName, b, opts...)
	if err == nil {
		if c, ok := any(v).(CheckSummable); ok {
			c.SetChecksum(sha1.Sum(b))
		}
	}
	return p, v, err
}

func (p *defaultParser[G]) ParseString(fileName, src string, opts ...participle.ParseOption) (*G, error) {
	return p.ParseBytes(fileName, []byte(src), opts...)
}

func (p *defaultParser[G]) ParseFile(fileName string, opts ...participle.ParseOption) (*G, error) {
	return p.init(p.parseFile(fileName, opts...))
}

func (p *defaultParser[G]) parseFile(fileName string, opts ...participle.ParseOption) (Parser[G], *G, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	return p.parseBytes(fileName, b, opts...)
}

type Merge[G any] interface {
	Merge(*G) (*G, error)
}

func (p *defaultParser[G]) ParseFiles(fileNames ...string) (*G, error) {
	var r *G
	for _, n := range fileNames {
		// parse but do not init here
		_, script, err := p.parseFile(n)
		if err == nil {
			if r == nil {
				// First entry then use it
				r = script
			} else {
				// Try to merge but fail if r does not implement Merge[G]
				m, ok := any(r).(Merge[G])
				if !ok {
					panic(fmt.Errorf("cannot merge %T as it does not support merging", r))
				}
				r, err = m.Merge(script)
			}
		}
		if err != nil {
			return nil, err
		}
	}

	// Run init against the final merged entity
	return p.init(p, r, nil)
}
