package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"github.com/peter-mount/piweather.center/util"
)

var (
	expressionParser = participle.MustBuild[lang.Expression](
		participle.Lexer(scriptLexer),
		participle.UseLookahead(2),
		participle.Unquote("String"),
		participle.CaseInsensitive("Keyword"),
	)
)

// ExpressionParser allows just an Expression to be parsed.
// This is useful to verify that an expression is valid in a client before sending a query to the server.
type ExpressionParser interface {
	Parse(s string) (*lang.Expression, error)
}

type defaultExpressionParser struct {
	lexer  *lexer.StatefulDefinition
	parser *participle.Parser[lang.Expression]
}

func NewExpressionParser() ExpressionParser {
	return &defaultExpressionParser{
		lexer:  scriptLexer,
		parser: expressionParser,
	}
}

func (p *defaultExpressionParser) Parse(s string) (*lang.Expression, error) {
	return p.init(p.parser.ParseString("", s))
}

func (p *defaultExpressionParser) init(q *lang.Expression, err error) (*lang.Expression, error) {
	if err == nil {
		parserState := &parserState{usingNames: util.NewStringSet()}
		err = q.Accept(lang.NewBuilder().
			ExpressionModifier(parserState.expressionModifierInit).
			Function(functionInit).
			Metric(metricInit).
			Time(timeInit).
			Duration(durationInit).
			Build())
	}
	return q, err
}
