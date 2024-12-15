package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type Task struct {
	Pos        lexer.Position
	CronTab    time.CronTab     `parser:"'schedule' @@"`           // primary cron schedule
	Conditions []*TaskCondition `parser:"( @@+"`                   // Condition list
	Default    command.Command  `parser:"  ( 'default' ':' @@ )?"` // Optional command if no condition is met
	Execute    command.Command  `parser:"| @@ )"`                  // Command when no conditions present
}

func (c *visitor[T]) Task(d *Task) error {
	var err error
	if d != nil {
		if c.task != nil {
			err = c.task(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.CronTab(d.CronTab)
		}

		if err == nil {
			switch {

			case len(d.Conditions) > 0:

			case d.Execute != nil:
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Task(f func(Visitor[T], *Task) error) Builder[T] {
	b.task = f
	return b
}
