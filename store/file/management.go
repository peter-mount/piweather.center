package file

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/util"
	"path/filepath"
	"strings"
	"time"
)

// GenKey returns the file name for a metric at a specific time.Time
// Note the file name returned will be that in UTC not the timezone for the passed time
func GenKey(metric string, t time.Time) string {
	// Always store in UTC
	t = t.UTC()

	// Path for home.test.temp for 2023 Oct 20 12:13:14 UTC becomes home/test/temp/2023/10/20.mdb
	name := filepath.Join(strings.Split(metric, ".")...)
	return filepath.Join(name, util.Itoa(t.Year(), 4), util.Itoa(int(t.Month()), 2), util.Itoa(t.Day(), 2)+".mdb")
}

// removeFile removes an entry from openFiles
func (s *store) removeFile(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.removeFileInternal(key)
}

// removeFile removes an entry from openFiles - the mutex must be locked at this point
func (s *store) removeFileInternal(key string) {
	if f, exists := s.openFiles[key]; exists {
		if f.isOpen() {
			_ = f.Close()
		}
		delete(s.openFiles, key)
	}
}

// getFile returns the named entry in openFiles or nil if not present.
// This will touch the file to prevent it from expiry
func (s *store) getFile(k string) *File {
	return s.getFileImpl(k, true)
}

// getFileImpl returns the named entry in openFiles or nil if not present.
// This will touch the file if touch=true to prevent it from expiry.
// touch=false will not touch it so it could expire once this returns
func (s *store) getFileImpl(k string, touch bool) *File {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.getFileInternal(k, touch)
}

// getFileInternal returns the named entry - the mutex must be locked at this point
func (s *store) getFileInternal(key string, touch bool) *File {
	if f, exists := s.openFiles[key]; exists {
		if f.isOpen() {
			if touch {
				f.touch()
			}
			return f
		}

		// not open so remove it from the cache
		// do not use removeFile here as we are already locked!
		delete(s.openFiles, key)
	}
	return nil
}

// openOrCreateFile opens a file and adds it to openFiles.
// If the file does not exist it will be created.
func (s *store) openOrCreateFile(metric string, t time.Time) (*File, error) {
	return s.openOrCreateFileImpl(metric, t, true)
}

// openFile opens a file and adds it to openFiles.
// If the file does not exist it does nothing and returns nil
func (s *store) openFile(metric string, t time.Time) (*File, error) {
	return s.openOrCreateFileImpl(metric, t, false)
}

func (s *store) openOrCreateFileImpl(metric string, t time.Time, create bool) (*File, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Enforce UTC
	t = t.UTC()

	key := GenKey(metric, t)

	// Check we have it already open
	if f := s.getFileInternal(key, true); f != nil {
		return f, nil
	}

	// Try opening it. If it doesn't exist then create it
	fileName := filepath.Join(*s.BaseDir, key)
	f, err := openFile(fileName)
	if create && err == nil && f == nil {
		// Doesn't exist so create it if requested to do so
		f, err = createFile(fileName, metric)
	}

	if err == nil && f != nil {
		s.addFile(key, f)
	}

	return f, err
}

func (s *store) addFile(key string, f *File) {
	if key != "" && f != nil {
		s.openFiles[key] = f
		f.touch()
	}
}

// getFileKeys returns a slice of key names from openFiles
func (s *store) getFileKeys() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var keys []string
	for k, _ := range s.openFiles {
		keys = append(keys, k)
	}
	return keys
}

// close open files. If all is true then all files, false then only unused ones.
func (s *store) close(all bool) {
	// Add 10 seconds to allow for something accessing it at the moment we run
	expiry := (time.Duration(*s.FileExpiry) * time.Minute) + (10 * time.Second)
	now := time.Now().Add(-expiry)
	for _, key := range s.getFileKeys() {
		// Note we do not want to touch the file here!
		f := s.getFileImpl(key, false)
		if f != nil && (all || f.Expired(now)) {
			s.closeFile(f)
		}
	}
}

// closeFile closes an openFile
func (s *store) closeFile(f *File) {
	if f != nil {
		// TODO check concurrency, should we lock against f.Close() some how if another goroutine is using it?
		defer func() {
			s.removeFile(f.name)

			if log.IsVerbose() {
				log.Printf("Closed %q", f.name)
			}
		}()
		_ = f.Close()
	}
}
