package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type UsingDefinition struct {
	Pos      lexer.Position
	Name     string                `parser:"@String 'as'"`
	Modifier []*ExpressionModifier `parser:"(@@)+"`
}

func initUsingDefinition(v Visitor[*parserState], u *UsingDefinition) error {
	p := v.Get()

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
