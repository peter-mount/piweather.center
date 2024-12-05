package parser

import (
	"github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/config/util"
)

func New() util.Parser[ql.Query] {
	return newParser[ql.Query](scriptInit)
}

func NewExpressionParser() util.Parser[ql.Expression] {
	return newParser[ql.Expression](expressionInit)
}

func newParser[G any](init util.ParserInit[G]) util.Parser[G] {
	return util.NewParser[G](nil, nil, init)
}
