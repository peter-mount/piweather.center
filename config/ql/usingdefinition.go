package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type UsingDefinition struct {
	Pos      lexer.Position
	Name     string                `parser:"@String 'as'"`
	Modifier []*ExpressionModifier `parser:"(@@)+"`
}

func (b *builder[T]) UsingDefinition(f func(Visitor[T], *UsingDefinition) error) Builder[T] {
	b.common.usingDefinition = f
	return b
}

func (v *visitor[T]) UsingDefinition(b *UsingDefinition) error {
	if b != nil && v.usingDefinition != nil {
		return v.usingDefinition(v, b)
	}
	return nil
}
