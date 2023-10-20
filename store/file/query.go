package file

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"time"
)

type Query interface {
	HasNext() bool
	Next() Record
}

type queryScanner struct {
	metric     string    // Metric being iterated
	start      time.Time // start time inclusive
	end        time.Time // end time exclusive
	date       time.Time // Current date in search
	store      *store    // link to Store
	file       *File     // current file or nil if non
	recordNum  int       // entry in current file
	entryCount int       // number of entries in file at the point it was requested
	record     Record    // current record
}

func (q *queryScanner) Next() Record {
	return q.record
}

func (q *queryScanner) HasNext() bool {
	// Stop once we hit the end
	if q.date.After(q.end) {
		return false
	}

	// We have a file but at the end then disassociate the file
	if q.file != nil && q.recordNum >= q.entryCount {
		q.file = nil
	}

	// no file then look up the next one
	if q.file == nil {
		q.nextFile()
	}

	// Still no file then stop
	if q.file == nil {
		return false
	}

	r, err := q.file.GetRecord(q.recordNum)
	if err != nil {
		log.Println(err)
		return false
	}

	q.record = r
	q.recordNum++
	return true
}

func (q *queryScanner) nextFile() {
	q.file = nil
	for q.file == nil && q.date.Before(q.end) {
		f, err := q.store.openFile(q.metric, q.date)
		if err == nil && f != nil {
			entryCount, err := f.EntryCount()
			if err == nil {
				q.file = f
				q.recordNum = 0
				q.entryCount = entryCount
			}
		}

		q.date = q.date.Add(24 * time.Hour)
	}
}

// filteredQuery wraps a Query with a Filter
type filteredQuery struct {
	query  Query
	filter Filter
	record Record
}

func (q *filteredQuery) HasNext() bool {
	found := false
	for !found && q.query.HasNext() {
		q.record = q.query.Next()
		found = q.filter(q.record)
	}
	return found
}

func (q *filteredQuery) Next() Record {
	return q.record
}
