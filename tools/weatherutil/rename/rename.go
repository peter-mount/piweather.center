package rename

import (
	"errors"
	"flag"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/tools/weatherutil"
	"time"
)

// Rename is a utility to rename a metric
type Rename struct {
	Rename *bool              `kernel:"flag,rename-metric,rename a metric,false"`
	Store  *weatherutil.Store `kernel:"inject"`
}

func (r *Rename) Run() error {
	if *r.Rename {
		args := flag.Args()
		if len(args) != 2 {
			return errors.New("syntax: -rename-metric old.metric new.metric")
		}
		return r.rename(args[0], args[1])
	}
	return nil
}

func (r *Rename) rename(fromMetric, toMetric string) error {
	// Test to check the destination metric does not exist
	if existing, _ := r.Store.GetFiles(toMetric); len(existing) > 0 {
		return fmt.Errorf("existing entries found for %q", toMetric)
	}

	filenames, err := r.Store.GetFiles(fromMetric)
	if err != nil {
		return err
	}
	if len(filenames) == 0 {
		return fmt.Errorf("no entries found for %q", fromMetric)
	}

	log.Printf("Renaming %s to %s", fromMetric, toMetric)

	for _, filename := range filenames {
		if date := r.Store.SplitFilename(filename); !date.IsZero() {
			err = r.renameMetric(fromMetric, toMetric, date)
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
