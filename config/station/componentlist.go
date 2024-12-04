package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type ComponentList struct {
	Pos     lexer.Position
	Entries []*ComponentListEntry `parser:"@@*"`
}

func (c *visitor[T]) ComponentList(d *ComponentList) error {
	var err error
	if d != nil {
		if c.componentList != nil {
			err = c.componentList(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Entries {
				err = c.ComponentListEntry(e)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) ComponentList(f func(Visitor[T], *ComponentList) error) Builder[T] {
	b.componentList = f
	return b
}
