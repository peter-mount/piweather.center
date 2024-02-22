package parser

import (
	"github.com/peter-mount/piweather.center/config/util"
	lang2 "github.com/peter-mount/piweather.center/config/util/ql"
)

func New() util.Parser[lang2.Query] {
	return newParser[lang2.Query](nil)
}

func NewExpressionParser() util.Parser[lang2.Expression] {
	return newParser[lang2.Expression](expressionInit)
}

func newParser[G any](init util.ParserInit[G]) util.Parser[G] {
	return util.NewParser[G](nil, nil, init)
}
