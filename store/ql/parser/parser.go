package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/ql/lang"
)

const (
	keywords = `(?i)\b(ADD|AND|AS|AT|BETWEEN|DECLARE|EVERY|FOR|FROM|HISTOGRAM|LIMIT|OFFSET|SELECT|TRUNCATE|UNIT|USING|WINDROSE)\b`
)

func New() util.Parser[lang.Query] {
	return newParser[lang.Query](nil)
}

func NewExpressionParser() util.Parser[lang.Expression] {
	return newParser[lang.Expression](expressionInit)
}

func newParser[G any](init util.ParserInit[G]) util.Parser[G] {
	return util.NewParser[G](
		[]lexer.SimpleRule{
			lexer.SimpleRule{Name: "Keyword", Pattern: keywords},
		},
		[]participle.Option{
			participle.CaseInsensitive("Keyword"),
		},
		init)
}
