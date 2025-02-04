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
			err = visitTasks[T](c, d)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func visitTasks[T any](v Visitor[T], d *Tasks) error {
	var err error
	if d != nil {
		for _, task := range d.Tasks {
			err = v.Task(task)
			if err != nil {
				break
			}
		}
	}
	return err
}

func (b *builder[T]) Tasks(f func(Visitor[T], *Tasks) error) Builder[T] {
	b.tasks = f
	return b
}

func printTasks(v Visitor[*printState], d *Tasks) error {
	return v.Get().
		Start().
		AppendPos(d.Pos).
		AppendHead("tasks(").
		AppendFooter(")").
		EndError(d.Pos, visitTasks(v, d))
}
