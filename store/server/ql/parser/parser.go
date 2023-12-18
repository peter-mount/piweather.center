package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/server/ql/functions"
	"github.com/peter-mount/piweather.center/store/server/ql/lang"
	"github.com/peter-mount/piweather.center/util/unit"
	"io"
	"os"
	"strings"
	"time"
)

var (
	scriptLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Keyword", `(?i)\b(ADD|AND|AS|AT|BETWEEN|EVERY|FOR|FROM|LIMIT|OFFSET|SELECT|TRUNCATE)\b`},
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

func (p *defaultParser) init(q *lang.Query, err error) (*lang.Query, error) {
	if err == nil {
		err = q.Accept(lang.NewBuilder().
			Query(queryInit).
			QueryRange(queryRangeInit).
			Select(selectInit).
			Function(functionInit).
			Metric(metricInit).
			Time(timeInit).
			Duration(durationInit).
			Build())
	}
	return q, err
}

func assertLimit(p lexer.Position, l int) error {
	if l < 0 {
		return participle.Errorf(p, "invalid LIMIT %d", l)
	}
	return nil
}

func queryInit(_ lang.Visitor, s *lang.Query) error {
	return assertLimit(s.Pos, s.Limit)
}

func queryRangeInit(v lang.Visitor, q *lang.QueryRange) error {
	// If no Every statement then set it to 1 minute
	if q.Every == nil {
		q.Every = &lang.Duration{
			Pos:      q.Pos,
			Duration: time.Minute,
			Def:      "1m",
		}
	}

	if err := v.Duration(q.Every); err != nil {
		return err
	}

	// Negative duration for Every is invalid
	if q.Every.Duration < time.Second {
		return fmt.Errorf("invalid step size %v", q.Every.Duration)
	}

	if err := v.Time(q.At); err != nil {
		return err
	}

	if err := v.Time(q.From); err != nil {
		return err
	}
	if err := v.Duration(q.For); err != nil {
		return err
	}

	if err := v.Time(q.Start); err != nil {
		return err
	}
	if err := v.Time(q.End); err != nil {
		return err
	}

	return lang.VisitorStop
}

func selectInit(_ lang.Visitor, s *lang.Select) error {
	return assertLimit(s.Pos, s.Limit)
}

func functionInit(_ lang.Visitor, b *lang.Function) error {
	if functions.HasFunction(b.Name) {
		return nil
	}
	return participle.Errorf(b.Pos, "unknown function %q", b.Name)
}

func metricInit(_ lang.Visitor, b *lang.Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

func timeInit(v lang.Visitor, t *lang.Time) error {
	if t == nil {
		return nil
	}

	t.Time = unit.ParseTime(t.Def)
	if t.Time.IsZero() {
		return participle.Errorf(t.Pos, "invalid datetime")
	}

	for _, e := range t.Expression {
		switch {
		case e.Add != nil:
			if err := v.Duration(e.Add); err != nil {
				return err
			}
			t.Time = t.Time.Add(e.Add.Duration)

		case e.Truncate != nil:
			if err := v.Duration(e.Truncate); err != nil {
				return err
			}
			t.Time = t.Time.Truncate(e.Truncate.Duration)
		}
	}

	return lang.VisitorStop
}

func durationInit(_ lang.Visitor, d *lang.Duration) error {
	if d.Def != "" {
		v, err := time.ParseDuration(d.Def)
		if err != nil {
			return err
		}
		d.Set(v)
	}

	return nil
}
