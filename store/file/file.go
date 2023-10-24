package file

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/file/record"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type File struct {
	mutex      sync.Mutex           // Mutex used to keep reads/writes atomic
	name       string               // Name of file
	header     record.FileHeader    // File header
	file       io.ReadWriteSeeker   // File access
	closer     io.Closer            // To close underlying file
	handler    record.RecordHandler // Handler for version
	lastAccess time.Time            // Time of last access
	latest     record.Record        // The most recent record
}

func (f *File) touch() {
	f.lastAccess = time.Now()
}

func (f *File) Expired(t time.Time) bool {
	return f.lastAccess.Before(t)
}

func (f *File) Close() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.closer != nil && f.file != nil {
		defer func() {
			f.closer = nil
			f.file = nil
		}()
		return f.closer.Close()
	}
	return nil
}

func (f *File) IsOpen() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.isOpen()
}

func (f *File) isOpen() bool {
	return f.closer != nil && f.file != nil
}

func (f *File) assertOpen() error {
	if !f.isOpen() {
		return fmt.Errorf("%s: already closed", f.header.Name)
	}
	return nil
}

// Size returns the total file size in bytes
func (f *File) Size() (int, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.size()
}

// size must be called from within a lock
func (f *File) size() (int, error) {
	if err := f.assertOpen(); err != nil {
		return 0, err
	}

	l, err := f.file.Seek(0, io.SeekEnd)
	return int(l), err
}

// EntryCount returns the number of metrics in the file
func (f *File) EntryCount() (int, error) {
	s, err := f.Size()
	if err == nil && s > f.header.Size {
		return (s - f.header.Size) / f.header.RecordLength, nil
	}
	return 0, err
}

func (f *File) Append(rec record.Record) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if err := f.assertOpen(); err != nil {
		return err
	}

	// Touch the file as we are modifying it
	f.touch()

	b := f.handler.Append(nil, rec)

	// Seek to the end of the file
	offset, err := f.file.Seek(0, io.SeekEnd)

	if err == nil {
		// Test the file end is at the end of a record.
		lastRecSize := (int(offset) - f.header.Size) % f.header.RecordLength
		if lastRecSize != 0 {
			// If not, log a warning and try to fix by removing the extra bytes on the end
			_, _ = fmt.Fprintf(os.Stderr, "Warning: last record size is %d bytes not %d, trying to fix %s", lastRecSize, f.header.RecordLength, f.name)

			// Not ideal but loose the last record
			offset, err = f.file.Seek(offset-int64(lastRecSize), io.SeekStart)
			if err == nil {
				lastRecSize = (int(offset) - f.header.Size) % f.header.RecordLength
				if lastRecSize != 0 {
					_, _ = fmt.Fprintf(os.Stderr, "Warning: file potentially corrupt %q", f.name)
				}
			}
		}
	}

	if err == nil {
		n, err1 := f.file.Write(b)
		switch {
		case err1 != nil:
			err = err1
		case n != len(b):
			err = fmt.Errorf("%s: wrote %d/%d bytes", f.header.Name, n, len(b))
		default:
			f.latest = rec
		}
	}

	return err
}

func (f *File) GetLatestRecord() (record.Record, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.latest.IsValid() {
		return f.latest, nil
	}

	// Seek to start of most recent record
	offset, err := f.file.Seek(-int64(f.header.RecordLength), io.SeekEnd)
	if err == nil && offset >= int64(f.header.Size) {
		// We have a valid entry
		f.latest, err = f.readRecord()
	} else {
		// Replace with an invalid entry
		f.latest = record.Record{}
	}

	return f.latest, err
}

func (f *File) GetRecord(i int) (record.Record, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	var rec record.Record
	recordSize := f.header.RecordLength

	err := f.assertOpen()

	if err == nil {
		// Touch the file as we are reading from it.
		// Do this here so if we are not open we don't touch and it should then expire
		f.touch()

		var size int
		size, err = f.size()
		if err == nil {
			offset := f.header.Size + (i * recordSize)
			if offset >= size {
				err = io.EOF
			} else {
				_, err = f.file.Seek(int64(offset), io.SeekStart)
			}
		}
	}

	if err == nil {
		rec, err = f.readRecord()
	}

	return rec, err
}

func (f *File) readRecord() (rec record.Record, err error) {
	buf := make([]byte, f.header.RecordLength)
	n, err1 := f.file.Read(buf)
	switch {
	case err1 != nil:
		err = err1
	case n != f.header.RecordLength:
		err = fmt.Errorf("expected %d bytes got %d file %s", f.header.RecordLength, n, f.name)
	default:
		rec, err = f.handler.Read(buf)
	}
	return
}

// openFile opens the named file.
// Warning, the file will be open when this returns, so it's up to the caller to close it.
// returns nil,nil if the file does not exist.
func openFile(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	file := &File{
		name:       name,
		file:       f,
		closer:     f,
		lastAccess: time.Now(),
	}

	// Read the header and configure the RecordHandler for this files version
	err = file.header.Read(f)
	if err == nil {
		file.handler, err = file.header.GetRecordHandler()
	}

	// If we have an error then close the underlying file before returning the error
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	if log.IsVerbose() {
		log.Printf("opened %s", name)
	}

	return file, nil
}

// createFile creates a new file, erasing any existing one
// Warning, the file will be open when this returns, so it's up to the caller to close it.
// returns nil,nil if the file does not exist.
func createFile(name, metric string) (*File, error) {
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return nil, err
	}

	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	file := &File{
		name:       name,
		file:       f,
		closer:     f,
		lastAccess: time.Now(),
		header:     record.NewFileHeader(metric),
	}

	file.handler, err = file.header.GetRecordHandler()
	if err != nil {
		return nil, err
	}

	err = file.header.Write(f)
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	if log.IsVerbose() {
		log.Printf("created %s", name)
	}

	return file, nil
}
