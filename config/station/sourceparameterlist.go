package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type SourceParameterList struct {
	Pos        lexer.Position
	Parameters []*SourceParameterListEntry `parser:"@@*"`
}

func (c *visitor[T]) SourceParameterList(d *SourceParameterList) error {
	var err error

	if d != nil {
		if c.sourceParameterList != nil {
			err = c.sourceParameterList(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, e := range d.Parameters {
				err = c.SourceParameterListEntry(e)
				if err != nil {
					break
				}
			}
		}

		return errors.Error(d.Pos, err)
	}

	return err
}

func (b *builder[T]) SourceParameterList(f func(Visitor[T], *SourceParameterList) error) Builder[T] {
	b.sourceParameterList = f
	return b
}
