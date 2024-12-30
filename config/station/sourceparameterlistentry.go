package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type SourceParameterListEntry struct {
	Pos       lexer.Position
	Within    *SourceWithin    `parser:"( @@"`
	Parameter *SourceParameter `parser:"| @@ )"`
}

func (c *visitor[T]) SourceParameterListEntry(d *SourceParameterListEntry) error {
	var err error

	if d != nil {
		if c.sourceParameterListEntry != nil {
			err = c.sourceParameterListEntry(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			switch {
			case d.Within != nil:
				err = c.SourceWithin(d.Within)
			case d.Parameter != nil:
				err = c.SourceParameter(d.Parameter)
			}
		}

		return errors.Error(d.Pos, err)
	}

	return err
}

func (b *builder[T]) SourceParameterListEntry(f func(Visitor[T], *SourceParameterListEntry) error) Builder[T] {
	b.sourceParameterListEntry = f
	return b
}
