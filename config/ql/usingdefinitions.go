package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type UsingDefinitions struct {
	Pos  lexer.Position
	Defs []*UsingDefinition `parser:" 'declare' @@ (',' @@)* "`
}

func (b *builder[T]) UsingDefinitions(f func(Visitor[T], *UsingDefinitions) error) Builder[T] {
	b.common.usingDefinitions = f
	return b
}

func (v *visitor[T]) UsingDefinitions(b *UsingDefinitions) error {
	var err error

	if b != nil {
		if v.usingDefinitions != nil {
			err = v.usingDefinitions(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			for _, e := range b.Defs {
				err = v.UsingDefinition(e)
			}
		}
	}

	return err
}
