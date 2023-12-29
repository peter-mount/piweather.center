package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"io"
	"os"
)

var (
	scriptLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Keyword", `(?i)\b(ADD|AND|AS|AT|BETWEEN|DECLARE|EVERY|FOR|FROM|HISTOGRAM|LIMIT|OFFSET|ROSE|SELECT|TABLE|TRUNCATE|UNIT|USING|WINDROSE)\b`},
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
	})

	scriptParser = participle.MustBuild[lang.Query](
		participle.Lexer(scriptLexer),
		participle.UseLookahead(2),
		participle.Unquote("String"),
		participle.CaseInsensitive("Keyword"),
	)
)

// Parser allows a full Query to be parsed with it fully initialised
type Parser interface {
	Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*lang.Query, error)
	ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*lang.Query, error)
	ParseString(fileName, src string, opts ...participle.ParseOption) (*lang.Query, error)
	ParseFile(fileName string, opts ...participle.ParseOption) (*lang.Query, error)
}

type defaultParser struct {
	lexer       *lexer.StatefulDefinition
	parser      *participle.Parser[lang.Query]
	includePath []string
}

func New() Parser {
	return &defaultParser{
		lexer:  scriptLexer,
		parser: scriptParser,
	}
}

func (p *defaultParser) Parse(fileName string, r io.Reader, opts ...participle.ParseOption) (*lang.Query, error) {
	return p.init(p.parser.Parse(fileName, r, opts...))
}

func (p *defaultParser) ParseBytes(fileName string, b []byte, opts ...participle.ParseOption) (*lang.Query, error) {
	return p.init(p.parser.ParseBytes(fileName, b, opts...))
}

func (p *defaultParser) ParseString(fileName, src string, opts ...participle.ParseOption) (*lang.Query, error) {
	return p.init(p.parser.ParseString(fileName, src, opts...))
}

func (p *defaultParser) ParseFile(fileName string, opts ...participle.ParseOption) (*lang.Query, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return p.init(p.ParseBytes(fileName, b, opts...))
}
