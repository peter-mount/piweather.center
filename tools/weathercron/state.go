package weathercron

import (
	"fmt"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
)

type state struct {
	service  *Service
	station  *station.Station
	job      *station.Tasks
	jobEntry *Task // Task being created
	jobs     *Tasks
}

func (s *Service) loadJobs(stations *station.Stations) error {
	st := &state{
		service: s,
		jobs: &Tasks{
			cron: s.Cron,
		},
	}

	if err := station.NewBuilder[*state]().
		Metric(addMetric).
		Task(addTask).
		Build().
		Set(st).
		Stations(stations); err != nil {
		return err
	}

	// TODO replace existing Tasks with the new one
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.jobs != nil {
		oldJobs := s.jobs
		s.jobs = nil
		go oldJobs.Stop()
	}

	s.jobs = st.jobs
	return s.jobs.Start()
}

func addTask(v station.Visitor[*state], d *station.Task) error {
	st := v.Get()

	st.jobEntry = newTask(d)

	err := st.jobs.addJob(st.jobEntry)

	if err != nil {
		err = errors.Error(d.Pos, err)
	}

	return err
}

func addMetric(v station.Visitor[*state], d *station.Metric) error {
	fmt.Printf("metric %q\n", d.Name)
	st := v.Get()
	st.jobEntry.addMetric(d.Name)
	return util.VisitorStop
}
