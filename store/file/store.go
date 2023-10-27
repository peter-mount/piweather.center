package file

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	cron2 "gopkg.in/robfig/cron.v2"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func init() {
	kernel.RegisterAPI((*Store)(nil), &store{})
}

type Store interface {
	// Append a record to a metric
	Append(metric string, rec record.Record) error
	AppendBulk(metric string, rec record.Record) error
	Sync(metric string) error
	// GetRecord returns the numbered record for a metric on a specific date
	GetRecord(metric string, date time.Time, num int) (record.Record, error)
	// NumRecords returns the number of metrics for a metric on a specific date
	NumRecords(metric string, date time.Time) (int, error)
	// Query returns a builder to build a query against a metric
	Query(metric string) QueryBuilder
}

// Store manages all open and existing File's stored on disk.
// This is the main entry point for accessing them as it manages
// which ones are open
type store struct {
	Cron       *cron.CronService `kernel:"inject"`                                       // Cron to run periodic jobs
	Latest     memory.Latest     `kernel:"inject"`                                       // Used to store most recent metric
	BaseDir    *string           `kernel:"flag,metric-db,Directory for storing metrics"` // Base directory of database
	FileExpiry *int              `kernel:"flag,metric-expiry,Expiry time in minutes,2"`  // Expiry time for open files in minutes

	mutex     sync.Mutex       // Mutex for internal structures like openFiles
	openFiles map[string]*File // Map of open files. Use addFile, getFile, getFileKeys & removeFile only to access this
	expiryId  cron2.EntryID    // Expiry cron job ID
}

func (s *store) PostInit() error {
	// Ensure BaseDir is valid. If not then use the default. After this BaseDir will be the absolute path to the db
	if *s.BaseDir == "" {
		*s.BaseDir = filepath.Join(filepath.Dir(os.Args[0]), "../db/metrics")
	}
	if d, err := filepath.Abs(*s.BaseDir); err != nil {
		return err
	} else {
		*s.BaseDir = d
	}

	// Ensure open file expiry is set to a minimum of 1 minute
	if *s.FileExpiry < 1 {
		return fmt.Errorf("invalid metric-expiry, must be >= 1 minutes, got %d", s.FileExpiry)
	}

	return nil
}

func (s *store) Start() error {
	if err := os.MkdirAll(*s.BaseDir, 0755); err != nil {
		return err
	}

	s.openFiles = make(map[string]*File)

	// Expiry daemon
	if id, err := s.Cron.AddTask("* * * * ?", func(_ context.Context) error {
		s.close(false)
		return nil
	}); err != nil {
		return err
	} else {
		s.expiryId = id
	}

	return s.initLatest()
}

func (s *store) Stop() {
	s.Cron.Remove(s.expiryId)
	s.close(true)
}

func (s *store) Append(metric string, rec record.Record) error {
	return s.append(metric, rec, false)
}

func (s *store) AppendBulk(metric string, rec record.Record) error {
	return s.append(metric, rec, true)
}

func (s *store) append(metric string, rec record.Record, bulk bool) error {
	file, err := s.openOrCreateFile(metric, rec.Time)
	if err == nil {
		if bulk {
			err = file.AppendBulk(rec)
		} else {
			err = file.Append(rec)
		}
	}
	if err == nil {
		s.Latest.Append(metric, rec)
	}
	return err
}

func (s *store) Sync(metric string) error {
	file, err := s.openOrCreateFile(metric, time.Now())
	if err == nil {
		err = file.Sync()
	}
	return err
}

func (s *store) GetRecord(metric string, date time.Time, num int) (record.Record, error) {
	var rec record.Record
	file, err := s.openFile(metric, date)
	if err == nil {
		rec, err = file.GetRecord(num)
	}
	return rec, err
}

func (s *store) NumRecords(metric string, date time.Time) (int, error) {
	file, err := s.openFile(metric, date)
	if err == nil {
		return file.EntryCount()
	}
	return -1, err
}

func (s *store) GetLatestRecord(metric string, date time.Time) (record.Record, error) {
	var rec record.Record
	file, err := s.openFile(metric, date)
	if err == nil {
		rec, err = file.GetLatestRecord()
	}
	return rec, err
}
