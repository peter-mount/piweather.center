package weathercron

import (
	"github.com/peter-mount/go-kernel/v2/cron"
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
