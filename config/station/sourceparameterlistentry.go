package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
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
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.SourceWithin(d.Within)
		}

		if err == nil {
			err = c.SourceParameter(d.Parameter)
		}

		return errors.Error(d.Pos, err)
	}

	return err
}

func (b *builder[T]) SourceParameterListEntry(f func(Visitor[T], *SourceParameterListEntry) error) Builder[T] {
	b.sourceParameterListEntry = f
	return b
}
