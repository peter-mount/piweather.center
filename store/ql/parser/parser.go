package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/util/unit"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"os"
	"strings"
	"time"
)

var (
	scriptLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Keyword", `(?i)\b(ADD|AND|AS|AT|BETWEEN|DECLARE|EVERY|FOR|FROM|LIMIT|OFFSET|SELECT|TRUNCATE|UNIT|USING)\b`},
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

func (p *defaultParser) init(q *lang.Query, err error) (*lang.Query, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = q.Accept(lang.NewBuilder().
			Query(queryInit).
			QueryRange(queryRangeInit).
			UsingDefinition(parserState.usingDefinitionInit).
			Select(selectInit).
			AliasedExpression(aliasedExpressionInit).
			Expression(parserState.expressionInit).
			ExpressionModifier(parserState.expressionModifierInit).
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
		return errors.Errorf(p, "invalid LIMIT %d", l)
	}
	return nil
}

func queryInit(_ lang.Visitor, s *lang.Query) error {
	return assertLimit(s.Pos, s.Limit)
}

func queryRangeInit(v lang.Visitor, q *lang.QueryRange) error {
	// If no Every statement then set it to 1 minute
	if q.Every == nil {
		q.Every = &lang.Duration{Pos: q.Pos, Def: "1m"}
	}

	if err := v.Duration(q.Every); err != nil {
		return err
	}

	// Negative duration for Every is invalid
	if q.Every.Duration(0) < time.Second {
		return fmt.Errorf("invalid step size %v", q.Every.Duration(0))
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

func aliasedExpressionInit(_ lang.Visitor, s *lang.AliasedExpression) error {
	if s.Unit != "" {
		u, exists := value.GetUnit(s.Unit)
		if !exists {
			return errors.Errorf(s.Pos, "unsupported unit %q", s.Unit)
		}
		s.SetUnit(u)
	}
	return nil
}

func functionInit(_ lang.Visitor, b *lang.Function) error {
	if functions.HasFunction(b.Name) {
		return nil
	}
	return errors.Errorf(b.Pos, "unknown function %q", b.Name)
}

func metricInit(_ lang.Visitor, b *lang.Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

func timeInit(v lang.Visitor, t *lang.Time) error {
	if t == nil {
		return nil
	}

	if err := t.SetTime(unit.ParseTime(t.Def), 0, v); err != nil {
		return err
	}

	return lang.VisitorStop
}

func durationInit(_ lang.Visitor, d *lang.Duration) error {
	if d.Def != "" && !d.IsEvery() {
		v, err := time.ParseDuration(d.Def)
		if err != nil {
			return errors.Error(d.Pos, err)
		}
		d.Set(v)
	}

	return nil
}

type parserState struct {
	usingNames util.StringSet
}

func (p *parserState) usingDefinitionInit(v lang.Visitor, u *lang.UsingDefinition) error {
	if !p.usingNames.Add(u.Name) {
		return errors.Errorf(u.Pos, "alias %q already defined", u.Name)
	}
	for _, e := range u.Modifier {
		if err := v.ExpressionModifier(e); err != nil {
			return err
		}
	}
	return nil
}

func (p *parserState) expressionInit(_ lang.Visitor, s *lang.Expression) error {
	if s.Using != "" && !p.usingNames.Contains(s.Using) {
		return errors.Errorf(s.Pos, "%q undefined", s.Using)
	}
	return nil
}

func (p *parserState) expressionModifierInit(v lang.Visitor, s *lang.ExpressionModifier) error {
	err := v.QueryRange(s.Range)
	if err == nil {
		err = v.Duration(s.Offset)
	}
	return err
}
