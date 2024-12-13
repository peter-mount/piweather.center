package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

// SourceWithin allows you to embed SourceParameter(s) with a common path
type SourceWithin struct {
	Pos              lexer.Position
	Prefix           *SourcePath          `parser:"'within' '(' @@"` // Prefix for source parameters
	SourceParameters *SourceParameterList `parser:"@@ ')'"`          // Parameters to read from source
}

func (c *visitor[T]) SourceWithin(d *SourceWithin) error {
	var err error

	if d != nil {
		if c.sourceWithin != nil {
			err = c.sourceWithin(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.SourcePath(d.Prefix)
		}

		if err == nil {
			err = c.SourceParameterList(d.SourceParameters)
		}

		return errors.Error(d.Pos, err)
	}

	return err
}

func initSourceWithin(v Visitor[*initState], d *SourceWithin) error {
	s := v.Get()

	err := v.SourcePath(d.Prefix)

	if err == nil {
		// Append the within path, restoring it when we complete
		// This allows paths within this to get it prefixed
		old := s.sourcePath
		s.sourcePath = []string{}
		s.sourcePath = append(s.sourcePath, old...)
		s.sourcePath = append(s.sourcePath, d.Prefix.Path...)
		defer func() {
			s.sourcePath = old
		}()
	}

	if err == nil {
		err = v.SourceParameterList(d.SourceParameters)
	}

	if err != nil {
		return errors.Error(d.Pos, err)
	}

	return util.VisitorStop
}

func (b *builder[T]) SourceWithin(f func(Visitor[T], *SourceWithin) error) Builder[T] {
	b.sourceWithin = f
	return b
}
