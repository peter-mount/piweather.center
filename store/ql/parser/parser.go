package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
)

const (
	keywords = `(?i)\b(ADD|AND|AS|AT|BETWEEN|DECLARE|EVERY|FOR|FROM|HISTOGRAM|LIMIT|OFFSET|SELECT|TRUNCATE|UNIT|USING|WINDROSE)\b`
)

func New() util.Parser[lang2.Query] {
	return newParser[lang2.Query](nil)
}

func NewExpressionParser() util.Parser[lang2.Expression] {
	return newParser[lang2.Expression](expressionInit)
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
