package file

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/file/record"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type File struct {
	mutex      sync.Mutex           // Mutex used to keep reads/writes atomic
	name       string               // Name of file
	header     record.FileHeader    // File header
	file       *os.File             // Underlying file
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
	return f.close()
}

func (f *File) close() error {

	if f.file != nil {
		defer func() {
			f.file = nil
		}()
		return f.file.Close()
	}
	return nil
}

func (f *File) IsOpen() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.isOpen()
}

func (f *File) isOpen() bool {
	return f.file != nil
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
	return f.entryCountImpl(f.Size())
}

func (f *File) entryCount() (int, error) {
	return f.entryCountImpl(f.size())
}

func (f *File) entryCountImpl(s int, err error) (int, error) {
	if err == nil && s > f.header.Size {
		return (s - f.header.Size) / f.header.RecordLength, nil
	}
	return 0, err
}

func (f *File) Append(rec record.Record) error {
	return f.append(rec, true)
}

func (f *File) AppendBulk(rec record.Record) error {
	return f.append(rec, false)
}

func (f *File) append(rec record.Record, sync bool) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if err := f.assertOpen(); err != nil {
		return err
	}

	// Touch the file as we are modifying it
	f.touch()

	// Get the latest record in this file if we don't already have it cached
	latestRecord, err := f.getLatestRecord()

	// Check the file length is at an expected record start location.
	// If it is not, for now just remove the last partial record
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
		switch {
		// Write the record if it's due to go to the end of the file
		case !latestRecord.IsValid(), latestRecord.Time.Before(rec.Time):
			b := f.handler.Append(nil, rec)

			n, err1 := f.file.Write(b)
			switch {
			case err1 != nil:
				err = err1
			case n != len(b):
				err = fmt.Errorf("%s: wrote %d/%d bytes", f.header.Name, n, len(b))
			default:
				f.latest = rec
			}

		case latestRecord.Time.Equal(rec.Time):
			// Do nothing, duplicate entry - always keep first one
			return nil

		// new record is before last one so rebuild the file, inserting the record in the correct place
		default:
			return f.insertRecord(rec)
		}
	}

	// Force the data to disk
	if sync && err == nil && f.isOpen() {
		err = f.file.Sync()
	}

	return err
}

func (f *File) Sync() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.isOpen() {
		return f.file.Sync()
	}
	return nil
}

func (f *File) GetLatestRecord() (record.Record, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.getLatestRecord()
}

func (f *File) getLatestRecord() (record.Record, error) {
	if f.latest.IsValid() {
		return f.latest, nil
	}

	// Get current position
	pos, err := f.file.Seek(0, io.SeekCurrent)
	if err == nil {
		// Seek to start of most recent record
		offset, err1 := f.file.Seek(-int64(f.header.RecordLength), io.SeekEnd)
		if err1 != nil {
			err = err1
		} else if offset >= int64(f.header.Size) {
			// We have a valid entry
			f.latest, err = f.readRecord()
		}
	}

	// Restore previous position otherwise we could corrupt the database
	_, err1 := f.file.Seek(pos, io.SeekStart)
	if err1 != nil && err == nil {
		err = err1
	}

	if err != nil {
		// Replace with an invalid entry
		f.latest = record.Record{}
	}

	return f.latest, err
}

func (f *File) GetRecord(i int) (record.Record, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.getRecord(i)
}

func (f *File) getRecord(i int) (record.Record, error) {

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
	f, err := os.OpenFile(name, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	file := &File{
		name:       name,
		file:       f,
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

// insertRecord inserts a record into the File.
// This is usually called when the record being inserted is older than the
// latest entry, so we have to rebuild the file
func (f *File) insertRecord(rec record.Record) error {
	r, err := f.readAllRecords()

	if err == nil {
		r = f.insertRecords(r, rec)
	}

	if err == nil {
		err = f.writeAllRecords(r)
	}

	// Invalidate the latest record cache
	f.latest = record.Record{}

	return err
}

// readAllRecords will real all records in the File into a slice.
func (f *File) readAllRecords() ([]record.Record, error) {
	var r []record.Record

	entryCount, err := f.entryCount()
	if err != nil {
		return nil, err
	}

	for i := 0; i < entryCount; i++ {
		rec, err := f.getRecord(i)
		if err != nil {
			return nil, err
		}
		r = append(r, rec)
	}

	// Whilst here, sort into time order
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Time.Before(r[j].Time)
	})

	return r, nil
}

// insertRecords inserts a Record into a slice at the correct position in the slice
func (f *File) insertRecords(r []record.Record, rec record.Record) []record.Record {
	// No records so create a single entry
	if len(r) == 0 {
		return []record.Record{rec}
	}

	l := len(r)
	t0 := rec.Time
	for i := 0; i < l; i++ {
		if r[i].Time.After(t0) {
			switch {
			case i == 0:
				r = append([]record.Record{rec}, r...)
			case i == (l - 1):
				r = append(r, rec)
			default:
				r = append(append(r[:i], rec), r[i:]...)
			}
		}
	}

	// Simple append
	return append(r, rec)
}

// writeAllRecords writes all records into the File.
//
// This works by creating a temporary file containing the new data then
// renames the new file over the original one.
//
// As the File instance is locked, this should be atomic as far as the
// database engine is concerned.
func (f *File) writeAllRecords(recs []record.Record) error {
	tmpName := f.name + ".tmp"
	defer os.Remove(tmpName)

	if err := f.writeTmpFile(tmpName, recs); err != nil {
		return err
	}

	// Close the db file, but kept under lock.
	// We will reopen as needed after the lock is released
	if err := f.close(); err != nil {
		return err
	}

	if err := os.Rename(tmpName, f.name); err != nil {
		return err
	}

	return nil
}

// writeTmpFile writes a sloce of records into a new temporary file
func (f *File) writeTmpFile(tmpName string, recs []record.Record) error {
	nf, err := os.Create(tmpName)
	if err != nil {
		return err
	}
	defer nf.Close()

	err = f.header.Write(nf)
	if err != nil {
		return err
	}

	for _, r := range recs {
		b := f.handler.Append(nil, r)
		n, err := nf.Write(b)
		if err != nil {
			return err
		}
		if n != len(b) {
			return fmt.Errorf("%s: wrote %d/%d bytes", f.header.Name, n, len(b))
		}
	}

	return nil
}
