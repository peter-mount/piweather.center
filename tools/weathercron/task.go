package weathercron

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/command"
	cron2 "gopkg.in/robfig/cron.v2"
	"os"
	"os/exec"
	"strings"
)

// Task holds all the details of an individual job
type Task struct {
	entryId cron2.EntryID
	job     *station.Task
	metrics map[string]interface{}
}

func newTask(d *station.Task) *Task {
	return &Task{
		job:     d,
		metrics: make(map[string]interface{}),
	}
}

func (j *Task) addMetric(name string) {
	name = strings.ToLower(strings.TrimSpace(name))
	j.metrics[name] = true
}

var (
	jobRunner = station.NewBuilder[*Task]().
		Command(runCommand).
		Task(runTask).
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

func runTask(v station.Visitor[*Task], d *station.Task) error {
	var err error

	switch {
	case len(d.Conditions) > 0:
		err = runTaskConditions(v, d)

	case d.Execute != nil:
		err = v.Command(d.Execute)
	}

	if err == nil {
		err = util.VisitorStop
	}

	return err
}

func runTaskConditions(v station.Visitor[*Task], d *station.Task) error {
	// Run through each condition, stop on error or the first one to claim the entry
	for _, cond := range d.Conditions {
		claimed, err := runTaskCondition(v, cond)
		if claimed || err != nil {
			return err
		}
	}

	// No error, not claimed & we have a default then call it
	if d.Default != nil {
		return v.Command(d.Default)
	}

	return nil
}

func runTaskCondition(v station.Visitor[*Task], d *station.TaskCondition) (bool, error) {

	return false, nil
}

func runCommand(_ station.Visitor[*Task], d command.Command) error {
	cmd := exec.Command(d.Command(), d.Args()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
