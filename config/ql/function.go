package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

// Function handles function calls
type Function struct {
	Pos lexer.Position

	Name        string        `parser:"@Ident"`
	Expressions []*Expression `parser:"'(' (@@ (',' @@)*)? ')'"`
}

func (v *visitor) Function(b *Function) error {
	var err error
	if b != nil {
		if v.function != nil {
			err = v.function(v, b)
		}
		if util.IsVisitorStop(err) {
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

func (b *builder) Function(f func(Visitor, *Function) error) Builder {
	b.common.function = f
	return b
}
