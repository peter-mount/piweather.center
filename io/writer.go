package io

import (
	"compress/gzip"
	"encoding/gob"
	"io"
	"os"
)

// Writer writes data to an io.Writer
type Writer func(io.Writer) error

// NewWriter creates a new writer
func NewWriter(w ...Writer) Writer {
	switch len(w) {
	case 0:
		return nil
	case 1:
		return w[0]
	default:
		var ret Writer
		for _, b := range w {
			ret = ret.Then(b)
		}
		return ret
	}
}

// Then chains two Writers.
func (a Writer) Then(b Writer) Writer {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(w io.Writer) error {
		if err := a(w); err != nil {
			return err
		}
		return b(w)
	}
}

func (a Writer) Write(w io.Writer) error {
	if a == nil {
		return nil
	}
	return a(w)
}

// Gob will write a struct/value to the writer using the
// encoding/gob package.
func (a Writer) Gob(e any) Writer {
	return a.Then(func(w io.Writer) error {
		return gob.NewEncoder(w).Encode(e)
	})
}

// CreateFile creates a file using the writer
func (a Writer) CreateFile(filename string) error {
	if a == nil {
		return nil
	}

	//log.Printf("Creating %s", filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return a(f)
}

// Compress will compress using gzip
func (a Writer) Compress() Writer {
	if a == nil {
		return nil
	}

	return func(w io.Writer) error {
		gw := gzip.NewWriter(w)
		if err := a(gw); err != nil {
			return err
		}
		return gw.Close()
	}
}

func (a Writer) CompressIf(p bool) Writer {
	if p {
		return a.Compress()
	}
	return a
}
