package lang

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"io"
	"os"
)

var (
	scriptLexer = lexer.MustSimple([]lexer.SimpleRule{
		//{"Keyword", `(?i)\b(AMQP|CONSOLE|EXCHANGE|FORMAT|IN|METRIC|PUBLISH|UNIXTIME|URL|VALUE)\b`},
		{"hashComment", `#.*`},
		{"sheBang", `#\!.*`},
		{"comment", `//.*|/\*.*?\*/`},
		{"whitespace", `\s+`},
		{"Ident", `\b([a-zA-Z_][a-zA-Z0-9_]*)\b`},
		{"Punct", `[-,()*/+%{};&!=:<>\|]|\[|\]|\^`},
		{"Number", `[-+]?(\d+\.\d+)`},
		{"Int", `[-+]?\d+`},
		{"String", `"(\\"|[^"])*"`},
		{"Period", `(\.)`},
		{"NewLine", `[\n\r]+`},
		{"Comma", `,`},
		{"Query", `\?`},
	})

	scriptParser = participle.MustBuild[Script](
		participle.Lexer(scriptLexer),
		participle.UseLookahead(2),
		participle.Unquote("String"),
		//participle.CaseInsensitive("Keyword"),
	)
)

// Parser allows a full Query to be parsed with it fully initialised
type Parser interface {
	Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*Script, error)
	ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*Script, error)
	ParseString(fileName, src string, opts ...participle.ParseOption) (*Script, error)
	ParseFile(fileName string, opts ...participle.ParseOption) (*Script, error)
	ParseFiles(fileNames ...string) (*Script, error)
}

type defaultParser struct {
	lexer       *lexer.StatefulDefinition
	parser      *participle.Parser[Script]
	includePath []string
}

func NewParser() Parser {
	return &defaultParser{
		lexer:  scriptLexer,
		parser: scriptParser,
	}
}

func (p *defaultParser) Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*Script, error) {
	return p.init(p.parser.Parse(fileName, r, opts...))
}

func (p *defaultParser) ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*Script, error) {
	return p.init(p.parser.ParseBytes(fileName, b, opts...))
}

func (p *defaultParser) ParseString(fileName, src string, opts ...participle.ParseOption) (*Script, error) {
	return p.init(p.parser.ParseString(fileName, src, opts...))
}

func (p *defaultParser) ParseFile(fileName string, opts ...participle.ParseOption) (*Script, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return p.init(p.ParseBytes(fileName, b, opts...))
}

func (p *defaultParser) ParseFiles(fileNames ...string) (*Script, error) {
	var r *Script
	for _, n := range fileNames {
		script, err := p.ParseFile(n)
		if err == nil {
			r, err = r.merge(script)
		}
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
