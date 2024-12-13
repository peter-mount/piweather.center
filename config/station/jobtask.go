package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
)

type JobTask struct {
	Pos        lexer.Position
	Executable string `parser:"'execute' '(' @String ')'"`
}

func (c *visitor[T]) JobTask(d *JobTask) error {
	var err error
	if d != nil {
		if c.jobTask != nil {
			err = c.jobTask(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) JobTask(f func(Visitor[T], *JobTask) error) Builder[T] {
	b.jobTask = f
	return b
}
