package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type AliasedGroup struct {
	Pos         lexer.Position
	Name        string            `parser:"'group' '(' @String"`
	Expressions *SelectExpression `parser:"@@ ( ',' @@ )* ')'"`
}

func (v *visitor[T]) AliasedGroup(b *AliasedGroup) error {
	var err error
	if b != nil {
		if v.aliasedGroup != nil {
			err = v.aliasedGroup(v, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = v.SelectExpression(b.Expressions)
		}
	}
	return err
}

func (b *builder[T]) AliasedGroup(f func(Visitor[T], *AliasedGroup) error) Builder[T] {
	b.common.aliasedGroup = f
	return b
}
