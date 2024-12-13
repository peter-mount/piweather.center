package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type UseFirst struct {
	Pos    lexer.Position
	Metric *Metric `parser:"'usefirst' @@"`
}

func (c *visitor[T]) UseFirst(b *UseFirst) error {
	var err error
	if b != nil {
		if c.useFirst != nil {
			err = c.useFirst(c, b)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.Metric(b.Metric)
		}

		err = errors.Error(b.Pos, err)
	}
	return err
}

func (b *builder[T]) UseFirst(f func(Visitor[T], *UseFirst) error) Builder[T] {
	b.useFirst = f
	return b
}
