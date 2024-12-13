package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/time"
)

// Job represents a scheduled task
type Job struct {
	Pos     lexer.Position
	Name    string        `parser:"'job' '(' @String"`
	CronTab *time.CronTab `parser:"@@"`
	Task    *JobTask      `parser:"@@ ')'"`
}

func (c *visitor[T]) Job(d *Job) error {
	var err error
	if d != nil {
		if c.job != nil {
			err = c.job(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		if err == nil {
			err = c.CronTab(d.CronTab)
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (b *builder[T]) Job(f func(Visitor[T], *Job) error) Builder[T] {
	b.job = f
	return b
}
