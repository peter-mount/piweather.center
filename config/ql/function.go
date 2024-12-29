package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

// Function handles function calls
type Function struct {
	Pos lexer.Position

	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

func (v *visitor[T]) Function(b *Function) error {
	var err error
	if b != nil {
		if v.function != nil {
			err = v.function(v, b)
		}
		if errors.IsVisitorStop(err) {
			return nil
		}
		if err == nil {
			for _, ex := range b.Expressions {
				err = v.Expression(ex)
				if err != nil {
					break
				}
			}
		}
	}
	return err
}

func (b *builder[T]) Function(f func(Visitor[T], *Function) error) Builder[T] {
	b.common.function = f
	return b
}
