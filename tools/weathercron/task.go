package weathercron

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/station/expression"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/weather/value"
	cron2 "gopkg.in/robfig/cron.v2"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// Task holds all the details of an individual job
type Task struct {
	mutex    sync.Mutex
	entryId  cron2.EntryID       // Cron entryId
	station  *station.Station    // Station containing the Task
	job      *station.Task       // Task definition
	dbServer string              // Database server url
	latest   memory.Latest       // Latest metric service
	time     value.Time          // time of this run
	executor expression.Executor // Expression Executor for TaskCondition's
}

func newTask(dbServer string, s *station.Station, d *station.Task) *Task {
	return &Task{
		station:  s,
		job:      d,
		dbServer: dbServer,
		latest:   memory.NewLatest(),
	}
}

func (j *Task) addMetric(name string) {
	name = strings.ToLower(strings.TrimSpace(name))
	j.latest.Set(name, record.Record{})
}

func (j *Task) getExecutor() expression.Executor {
	if j.executor == nil {
		j.executor = expression.NewExecutor("", j.time, j.dbServer, j.latest)
	}

	return j.executor
}

var (
	jobRunner = station.NewBuilder[*Task]().
		Command(runCommand).
		Task(runTask).
		TaskCondition(runTaskCondition).
		Build()
)

func (j *Task) run() {
	log.Printf("Task %s triggered", j.job.Pos)

	err := jobRunner.Clone().
		Set(j).
		Task(j.job)

	if err != nil {
		log.Printf("Task %s failed: %v", j.job.Pos.String(), err)
	}
}

func (j *Task) loadMetrics() error {
	cl := client.Client{Url: j.dbServer}
	for _, metric := range j.latest.Metrics() {
		resp, err := cl.Metric(metric)
		if err != nil {
			return errors.Error(j.job.Pos, err)
		}
		if unit, ok := value.GetUnit(resp.Result.Unit); ok {
			j.latest.Set(metric, record.Record{
				Time:  resp.Result.Time,
				Value: unit.Value(resp.Result.Value),
			})
		}
	}
	return nil
}

func runTask(v station.Visitor[*Task], d *station.Task) error {
	var err error

	task := v.Get()

	switch {
	case len(d.Conditions) > 0:
		// Load the required metrics from the DB
		err = task.loadMetrics()

		// Run through each condition, stop on error or the first one to claim the entry via Break()
		if err == nil {
			for _, cond := range d.Conditions {
				err = v.TaskCondition(cond)

				// Stop on error or Break
				if err != nil {
					break
				}
			}
		}

		// Run default command if defined and no error (including Break)
		if err == nil {
			err = v.Command(d.Default)
		}

		// If break then clear the error. Must be after check for default to run
		if errors.IsBreak(err) {
			err = nil
			break
		}

	case d.Execute != nil:
		err = v.Command(d.Execute)
	}

	if err == nil {
		err = util.VisitorStop
	}

	return err
}

func runTaskCondition(v station.Visitor[*Task], d *station.TaskCondition) error {
	task := v.Get()
	executor := task.getExecutor()

	result, _, err := executor.Evaluate(d.Expression)

	if err == nil && result.IsValid() && !value.IsZero(result.Float()) {
		err = v.Command(d.Execute)
		if err == nil {
			// Break so we stop checking any more conditions within Task
			return errors.Break()
		}
	}

	if err == nil {
		err = util.VisitorStop
	}

	return err
}

func runCommand(_ station.Visitor[*Task], d command.Command) error {
	cmd := exec.Command(d.Command(), d.Args()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
