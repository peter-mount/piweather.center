package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

// Tasks represents a scheduled task
type Tasks struct {
	Pos   lexer.Position
	Tasks []*Task `parser:"'tasks' '(' @@* ')'"`
}

func (c *visitor[T]) Tasks(d *Tasks) error {
	var err error
	if d != nil {
		if c.tasks != nil {
			err = c.tasks(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			for _, task := range d.Tasks {
				err = c.Task(task)
				if err != nil {
					break
				}
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Tasks(f func(Visitor[T], *Tasks) error) Builder[T] {
	b.tasks = f
	return b
}
