package weathercron

import (
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	cron2 "gopkg.in/robfig/cron.v2"
	"os"
	"os/exec"
)

type Tasks struct {
	cron   *cron.CronService
	tasks  []*Task
	active []*Task
}

func (j *Tasks) addJob(job *Task) error {
	j.tasks = append(j.tasks, job)
	return nil
}

func (j *Tasks) Start() error {
	for _, job := range j.tasks {
		entryId, err := j.cron.AddFunc(job.job.CronTab.Definition(), job.run)
		if err != nil {
			j.Stop()
			return err
		}
		job.entryId = entryId
		j.active = append(j.active, job)
	}
	return nil
}

func (j *Tasks) Stop() {
	for _, job := range j.active {
		j.cron.Remove(job.entryId)
	}
}

// Task holds all the details of an individual job
type Task struct {
	entryId cron2.EntryID
	job     *station.Task
}

var (
	jobRunner = station.NewBuilder[*Task]().
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

func runTask(_ station.Visitor[*Task], d *station.Task) error {
	cmd := exec.Command(d.Execute.Command(), d.Execute.Args()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
