package ql

import (
	"github.com/peter-mount/piweather.center/config/util"
)

func NewParser() util.Parser[Query] {
	return newParser[Query](scriptInit)
}

func NewExpressionParser() util.Parser[Expression] {
	return newParser[Expression](expressionInit)
}

func newParser[G any](init util.ParserInit[G]) util.Parser[G] {
	return util.NewParser[G](nil, nil, init)
}
