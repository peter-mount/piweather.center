package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type Publisher struct {
	Pos lexer.Position
	Log bool `parser:"( @'log'"`
	DB  bool `parser:"| @'db' )"`
}

func (c *visitor[T]) Publisher(d *Publisher) error {
	var err error
	if d != nil {
		if c.publisher != nil {
			err = c.publisher(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Publisher(f func(Visitor[T], *Publisher) error) Builder[T] {
	b.publisher = f
	return b
}
