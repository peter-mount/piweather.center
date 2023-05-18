package archiver

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/station/payload"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Archiver struct {
	storeDir    *string    `kernel:"flag,archive-dir,Archive directory"`
	logMessages *bool      `kernel:"flag,archive-log,Dump messages to stdout"`
	worker      task.Queue `kernel:"worker"`
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

func (s *Archiver) Archive(rec *payload.Payload) {
	if *s.storeDir != "" {
		s.worker.AddTask(task.Of(s.archiveReadingDisk).WithValue("record", rec))
	} else {
		log.Printf("Store:ignore %v", rec)
	}
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

func (s *Archiver) Preload(ctx context.Context, id string, t task.Task) error {
	log.Printf("Archiver:Preloading %s", id)

	fileName := s.archiveFileName(id, time.Now().UTC())

	s.mutex.Lock()
	defer s.mutex.Unlock()

	lc := 0
	err := io.NewReader().
		ForEachLine(func(line string) error {
			p, err := payload.FromLog(line)
			if p != nil && err == nil {
				lc++
				_ = t.Do(p.AddContext(ctx))
			}
			return nil
		}).
		Open(fileName)

	// Ignore if the file doesn't exist
	if os.IsNotExist(err) {
		return nil
	}

	log.Printf("Archiver:Preloaded %d from %s", lc, id)
	return err
}
