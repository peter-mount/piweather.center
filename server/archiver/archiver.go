package archiver

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/weather/value"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Archiver struct {
	Store       *store.Store `kernel:"inject"`
	storeDir    *string      `kernel:"flag,archive-dir,Archive directory"`
	logMessages *bool        `kernel:"flag,archive-log,Dump messages to stdout"`
	worker      task.Queue   `kernel:"worker"`
	mutex       sync.Mutex
}

func FromContext(ctx context.Context) *Archiver {
	return ctx.Value("local.archiver").(*Archiver)
}

func (s *Archiver) AddContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "local.archiver", s)
}

func (s *Archiver) Start() error {
	if *s.storeDir == "" {
		log.Println("Store:persistence of readings disabled unless -store-dir is set")
	} else {
		log.Printf("Store:persisting to %s", *s.storeDir)
	}

	return nil
}
func (s *Archiver) archiveFileName(name string, t time.Time) string {
	p := filepath.Join(strings.Split(name, ".")...)
	return filepath.Join(*s.storeDir, p, t.UTC().Format("2006/01/02")+".log")
}

func (s *Archiver) Archive(ctx context.Context) error {
	rec := payload.GetPayload(ctx)
	if *s.storeDir != "" {
		s.worker.AddTask(task.Of(s.archiveReadingDisk).WithValue("record", rec))
	} else {
		log.Printf("Store:ignore %v", rec)
	}
	return nil
}

func (s *Archiver) archiveReadingDisk(ctx context.Context) error {
	rec := ctx.Value("record").(*payload.Payload)

	if *s.logMessages {
		log.Printf("Archive:%v", string(rec.Msg()))
	}

	fileName := s.archiveFileName(rec.Id(), rec.Time())
	_ = s.appendReading(fileName, rec)
	return nil
}

func (s *Archiver) appendReading(fileName string, rec *payload.Payload) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(rec.ToLog())

	return nil
}

func (s *Archiver) Preload(ctx context.Context) error {

	// Load yesterday
	if err := s.preload(ctx, time.Now().Add(-24*time.Hour)); err != nil {
		return err
	}

	// Load today
	return s.preload(ctx, time.Now())
}

func (s *Archiver) preload(ctx context.Context, t time.Time) error {
	sensors := station.SensorsFromContext(ctx)
	fileName := s.archiveFileName(sensors.ID, t)

	// Visitor to process the reading into memory cache
	visitor := station.NewVisitor().
		Sensors(value.ResetMap).
		Reading(s.Store.ProcessReading).
		CalculatedValue(s.Store.Calculate)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	lc := 0
	err := io.NewReader().
		ForEachLine(func(line string) error {
			p, err := payload.FromLog(line)
			if p != nil && err == nil {
				lc++
				_ = visitor.
					WithContext(p.AddContext(ctx)).
					VisitSensors(sensors)
			}
			return nil
		}).
		Open(fileName)

	// Ignore if the file doesn't exist
	if os.IsNotExist(err) {
		return nil
	}

	return err
}
