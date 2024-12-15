package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/command"
)

type TaskCondition struct {
	Pos        lexer.Position
	Expression *Expression     `parser:"'case' @@"`
	Execute    command.Command `parser:"':' @@"`
}

func (c *visitor[T]) TaskCondition(d *TaskCondition) error {
	var err error
	if d != nil {
		if c.taskCondition != nil {
			err = c.taskCondition(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
			if errors.IsBreak(err) {
				return err
			}
		}

		if err == nil {
			err = c.Expression(d.Expression)
		}

		if err == nil {
			err = c.Command(d.Execute)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) TaskCondition(f func(Visitor[T], *TaskCondition) error) Builder[T] {
	b.taskCondition = f
	return b
}
