package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

// Current returns the current value of the calculation being performed
type Current struct {
	Pos     lexer.Position
	Current bool `parser:"@'current'"`
}

func (c *visitor[T]) Current(b *Current) error {
	var err error
	if b != nil {
		if c.current != nil {
			err = c.current(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) Current(f func(Visitor[T], *Current) error) Builder[T] {
	b.current = f
	return b
}
