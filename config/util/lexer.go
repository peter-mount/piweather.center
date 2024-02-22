package util

import (
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
		{"whitespace", `\s+`},
		//{"Ident", `([a-zA-Z_][a-zA-Z0-9_]*)`},
		{"Ident", `\b([a-zA-Z_][a-zA-Z0-9_]*)\b`},
		//{"Ident", `\b(([a-zA-Z_][a-zA-Z0-9_]*)(\.([a-zA-Z_][a-zA-Z0-9_]*))*)\b`},
		{"Punct", `[-,()*/+%{};&!=:<>\|]|\[|\]|\^`},
		{"Number", `[-+]?(\d+\.\d+)`},
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

type ParserInit[G any] func(q *G, err error) (*G, error)

type defaultParser[G any] struct {
	lexer  *lexer.StatefulDefinition
	parser *participle.Parser[G]
	init   ParserInit[G]
}

func NewParser[G any](rules []lexer.SimpleRule, opts []participle.Option, init ParserInit[G]) Parser[G] {
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
	}
	o = append(o, opts...)
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

func defaultInit[G any](q *G, err error) (*G, error) {
	return q, err
}

func (p *defaultParser[G]) Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*G, error) {
	return p.init(p.parser.Parse(fileName, r, opts...))
}

func (p *defaultParser[G]) ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*G, error) {
	return p.init(p.parser.ParseBytes(fileName, b, opts...))
}

func (p *defaultParser[G]) ParseString(fileName, src string, opts ...participle.ParseOption) (*G, error) {
	return p.init(p.parser.ParseString(fileName, src, opts...))
}

func (p *defaultParser[G]) ParseFile(fileName string, opts ...participle.ParseOption) (*G, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return p.init(p.ParseBytes(fileName, b, opts...))
}

type Merge[G any] interface {
	Merge(*G) (*G, error)
}

func (p *defaultParser[G]) ParseFiles(fileNames ...string) (*G, error) {
	var r *G
	for _, n := range fileNames {
		script, err := p.ParseFile(n)
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
	return r, nil
}
