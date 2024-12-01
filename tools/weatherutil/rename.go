package weatherutil

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/file"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Rename struct {
	FromMetric *string        `kernel:"flag,from-metric,Metric to rename to"`
	ToMetric   *string        `kernel:"flag,to-metric,Metric to rename to"`
	Store      file.Store     `kernel:"inject"`
	Daemon     *kernel.Daemon `kernel:"inject"`
}

func (r *Rename) Start() error {
	r.Daemon.ClearDaemon()

	return nil
}

func (r *Rename) Run() error {
	if *r.FromMetric != "" && *r.ToMetric != "" {
		return r.rename()
	}
	return nil
}

func (r *Rename) rename() error {
	// Test to check the destination metric does not exist
	if existing, _ := r.Store.GetFiles(*r.ToMetric); len(existing) > 0 {
		return fmt.Errorf("existing entries found for %q", *r.ToMetric)
	}

	filenames, err := r.Store.GetFiles(*r.FromMetric)
	if err != nil {
		return err
	}
	if len(filenames) == 0 {
		return fmt.Errorf("no entries found for %q", *r.FromMetric)
	}

	log.Printf("Renaming %s to %s", *r.FromMetric, *r.ToMetric)

	for _, filename := range filenames {
		if date := splitFilename(filename); !date.IsZero() {
			err = r.renameMetric(*r.FromMetric, *r.ToMetric, date)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Rename) renameMetric(from, to string, date time.Time) error {
	rec, err := r.Store.GetRecords(from, date)
	if err != nil {
		return err
	}

	defer r.Store.Sync(to)

	for _, record := range rec {
		err = r.Store.AppendBulk(to, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func splitFilename(filename string) time.Time {
	s := strings.SplitN(filename, ".", 2)
	s = strings.SplitN(s[0], string(filepath.Separator), 3)
	if len(s) == 3 {
		year, err1 := strconv.Atoi(s[0])
		month, err2 := strconv.Atoi(s[1])
		day, err3 := strconv.Atoi(s[2])
		if err1 == nil && err2 == nil && err3 == nil {
			return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		}
	}

	// Zero time if failed to parse
	return time.Time{}
}
