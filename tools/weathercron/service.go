package weathercron

import (
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/config"
	"path/filepath"
	"sync"
)

type Service struct {
	Daemon   *kernel.Daemon    `kernel:"inject"`
	Cron     *cron.CronService `kernel:"inject"`
	Config   config.Manager    `kernel:"inject"`
	Stations *station.Stations `kernel:"inject"`
	// internal from here
	dashDir string
	mutex   sync.Mutex
	jobs    *Tasks // The loaded jobs
}

const (
	dashDir    = "stations"
	fileSuffix = ".cron"
)

func (s *Service) Start() error {
	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)

	// Load existing tasks
	stations, err := s.Stations.LoadDirectory(s.dashDir, fileSuffix, station.JobOption)
	if err != nil {
		return err
	}

	// Configure the tasks
	err = s.loadJobs(stations)
	if err != nil {
		return err
	}

	log.Println(version.Version)
	return nil
}
