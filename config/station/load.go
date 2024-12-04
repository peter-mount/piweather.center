package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Load struct {
	Pos  lexer.Position
	When string `parser:"'load' @String"` // When to load from
	With string `parser:"'with' @String"` // Query to perform
}

func (c *visitor[T]) Load(b *Load) error {
	var err error
	if b != nil {
		if c.load != nil {
			err = c.load(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) Load(f func(Visitor[T], *Load) error) Builder[T] {
	b.load = f
	return b
}
