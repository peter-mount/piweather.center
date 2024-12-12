package file

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/walk"
	strings2 "github.com/peter-mount/piweather.center/util/strings"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// GenKey returns the file name for a metric at a specific time.Time
// Note the file name returned will be that in UTC not the timezone for the passed time
func GenKey(metric string, t time.Time) string {
	// Always store in UTC
	t = t.UTC()

	// Path for home.test.temp for 2023 Oct 20 12:13:14 UTC becomes home/test/temp/2023/10/20.mdb
	name := GetMetricBasePath(metric)
	return filepath.Join(name, strings2.Itoa(t.Year(), 4), strings2.Itoa(int(t.Month()), 2), strings2.Itoa(t.Day(), 2)+".mdb")
}

// GetMetricBasePath returns the metrics base bath within the database.
func GetMetricBasePath(metric string) string {
	return filepath.Join(strings.Split(metric, ".")...)
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

func (s *store) RemoveFile(metric string, t time.Time) error {
	key := GenKey(metric, t)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check we have it already open
	if f := s.getFileInternal(key, false); f != nil {
		s.closeFile(f)
	}

	fileName := filepath.Join(*s.BaseDir, key)
	err := os.Remove(fileName)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	log.Printf("deleted %s", key)
	return nil
}

func (s *store) addFile(key string, f *File) {
	if key != "" && f != nil {
		s.openFiles[key] = f
		f.touch()
	}

	if len(s.openFiles) > *s.MaxOpenFiles {
		go s.closeOldestFile()
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
				log.Printf("closed %s", f.name)
			}
		}()
		_ = f.Close()
	}
}

func (s *store) closeOldestFile() {
	var files []*File
	for _, k := range s.getFileKeys() {
		f := s.getFileImpl(k, false)
		if f != nil {
			files = append(files, f)
		}
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].lastAccess.Before(files[j].lastAccess)
	})

	for len(files) > 0 {
		if files[0].IsOpen() {
			s.closeFile(files[0])
			return
		}
		files = files[1:]
	}
}

func (s *store) GetFiles(metric string) ([]string, error) {
	var r []string

	basePath := filepath.Join(*s.BaseDir, GetMetricBasePath(metric))

	err := walk.NewPathWalker().
		Then(func(path string, info os.FileInfo) error {
			p, err := filepath.Rel(basePath, path)
			if err == nil && len(p) > 0 {
				ps := strings.SplitN(p, string(filepath.Separator), 2)
				// Only accept integer here - e.g. should be a year
				year, err1 := strconv.Atoi(ps[0])
				if err1 == nil && year >= 100 {
					r = append(r, p)
				}
			}
			return err
		}).
		PathHasSuffix(".mdb").
		IsFile().
		Walk(basePath)

	return r, err
}
