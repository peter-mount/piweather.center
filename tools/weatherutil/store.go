package weatherutil

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/file/record"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Store struct {
	Store     file.Store     `kernel:"inject"`
	Daemon    *kernel.Daemon `kernel:"inject"`
	StartDate *string        `kernel:"flag,from,Start date"`
	EndDate   *string        `kernel:"flag,to,End date"`
	start     time.Time
	end       time.Time
	today     time.Time
}

func (r *Store) Start() error {
	r.Daemon.ClearDaemon()

	r.today = time2.LocalMidnight(time.Now().UTC())

	if *r.StartDate != "" {
		t := time2.ParseTime(*r.StartDate)
		if !t.IsZero() {
			r.start = time2.LocalMidnight(t)
		}
	}

	if *r.EndDate != "" {
		t := time2.ParseTime(*r.EndDate)
		if !t.IsZero() {
			r.end = time2.LocalMidnight(t)
		}
	}

	if !r.start.IsZero() && !r.end.IsZero() && r.end.Before(r.start) {
		r.start, r.end = r.end, r.start
	}

	return nil
}

func (r *Store) IsBeforeToday(t time.Time) bool {
	return t.Before(r.today)
}

func (r *Store) IsDateValid(t time.Time) bool {
	if r.IsBeforeToday(t) {
		s, e := r.start, r.end
		switch {
		case s.IsZero() && e.IsZero():
			return true
		case s.IsZero():
			return !t.After(e)
		case e.IsZero():
			return !t.Before(s)
		default:
			return !(t.Before(s) || t.After(e))
		}
	}
	return false
}

func (r *Store) GetFiles(metric string) ([]string, error) {
	files, err := r.Store.GetFiles(metric)
	if err == nil {
		sort.SliceStable(files, func(i, j int) bool {
			return strings.Compare(files[i], files[j]) < 0
		})
	}
	return files, err
}

func (r *Store) GetRecords(metric string, date time.Time) ([]record.Record, error) {
	rec, err := r.Store.GetRecords(metric, date)
	if err == nil {
		sort.SliceStable(rec, func(i, j int) bool {
			return rec[i].Time.Before(rec[j].Time)
		})
	}
	return rec, err
}

func (r *Store) AppendBulk(metric string, rec record.Record) error {
	return r.Store.AppendBulk(metric, rec)
}

func (r *Store) RemoveFile(metric string, t time.Time) error {
	return r.Store.RemoveFile(metric, t)
}

func (r *Store) Sync(metric string) error {
	return r.Store.Sync(metric)
}

func (r *Store) SplitFilename(filename string) time.Time {
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
