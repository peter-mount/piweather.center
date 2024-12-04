package statistics

import (
	"errors"
	"flag"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	record2 "github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/tools/weatherutil"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
)

// Stats is a utility to
type Stats struct {
	Stat  *string            `kernel:"flag,refresh-statistic"`
	Store *weatherutil.Store `kernel:"inject"`
}

var (
	supportedFunctions = []string{"min", "max"}
)

func (r *Stats) Run() error {
	if *r.Stat != "" {
		var f value.Calculation
		switch *r.Stat {
		case "min":
			f = math.Min
		case "max":
			f = math.Max
		default:
			return fmt.Errorf("-refresh-statistic %q is not one of %q", *r.Stat, strings.Join(supportedFunctions, ","))
		}

		args := flag.Args()
		if len(args) == 0 {
			return errors.New("syntax: -refresh-statistic function metric")
		}
		for _, arg := range args {
			if err := r.refresh(arg, arg+"."+*r.Stat, f); err != nil {
				return err
			}
		}
		log.Printf("Stats completed")
	}
	return nil
}

func (r *Stats) refresh(from, to string, f value.Calculation) error {
	files, err := r.Store.GetFiles(from)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("no data for metric %q", from)
	}

	for _, file := range files {
		if err := r.refreshMetric(from, to, file, f); err != nil {
			return err
		}
	}
	return nil
}

func (r *Stats) refreshMetric(from, to, file string, f value.Calculation) error {
	t := r.Store.SplitFilename(file)
	if !r.Store.IsDateValid(t) {
		return nil
	}

	if err := r.Store.RemoveFile(to, t); err != nil {
		return err
	}

	rec, err := r.Store.GetRecords(from, t)
	if err != nil || len(rec) == 0 {
		return err
	}

	var current value.Value
	first := true
	for _, record := range rec {
		if first {
			current = record.Value
			first = !record.IsValid()
		} else {
			v, err := current.Calculate(record.Value, f)
			if err == nil && v.IsValid() {
				current = v
			}
		}

		newRec := record2.Record{
			Time:  record.Time,
			Value: current,
		}
		if newRec.IsValid() {
			err = r.Store.AppendBulk(to, newRec)
			if err != nil {
				return err
			}
		}
	}

	return r.Store.Sync(to)
}
