package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"strings"
)

// SourcePath is a series of keys to locate the required data
type SourcePath struct {
	Pos  lexer.Position
	Path []string `parser:"@String ('.' @String)*"`
}

func (c *visitor[T]) SourcePath(d *SourcePath) error {
	var err error
	if d != nil {
		if c.sourcePath != nil {
			err = c.sourcePath(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) SourcePath(f func(Visitor[T], *SourcePath) error) Builder[T] {
	b.sourcePath = f
	return b
}

func initSourcePath(v Visitor[*initState], d *SourcePath) error {
	for i, e := range d.Path {
		s := strings.TrimSpace(e)
		if e == "" {
			return participle.Errorf(d.Pos, "Path element %d is empty", i)
		}
		d.Path[i] = s
	}

	// Prefix any path by SourceWithin
	s := v.Get()
	if len(s.sourcePath) > 0 {
		var p []string
		p = append(p, s.sourcePath...)
		d.Path = append(p, d.Path...)
	}

	return nil
}

func (d *SourcePath) String() string {
	return strings.Join(d.Path, ".")
}
