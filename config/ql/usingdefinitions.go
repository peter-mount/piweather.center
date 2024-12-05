package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util"
)

type UsingDefinitions struct {
	Pos  lexer.Position
	Defs []*UsingDefinition `parser:" 'declare' @@ (',' @@)* "`
}

func (b *builder) UsingDefinitions(f func(Visitor, *UsingDefinitions) error) Builder {
	b.common.usingDefinitions = f
	return b
}

func (v *visitor) UsingDefinitions(b *UsingDefinitions) error {
	var err error

	if b != nil {
		if v.usingDefinitions != nil {
			err = v.usingDefinitions(v, b)
		}
		if util.IsVisitorStop(err) {
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
