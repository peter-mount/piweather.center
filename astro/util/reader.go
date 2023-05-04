package util

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/strings"
	"io"
	"os"
)

type Reader func(r io.Reader) error

func NewReader(r ...Reader) Reader {
	switch len(r) {
	case 0:
		return nil
	case 1:
		return r[0]
	default:
		var ret Reader
		for _, b := range r {
			ret = ret.Then(b)
		}
		return ret
	}
}

func (a Reader) Then(b Reader) Reader {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(r io.Reader) error {
		if err := a(r); err != nil {
			return err
		}
		return b(r)
	}
}

func (a Reader) Read(r io.Reader) error {
	if a == nil {
		return nil
	}
	return a(r)
}

// Open a file and pass to the Reader
func (a Reader) Open(filename string) error {
	if a == nil {
		return nil
	}

	log.Printf("Reading %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return a(f)
}

func (a Reader) Decompress() Reader {
	if a == nil {
		return nil
	}

	return func(r io.Reader) error {
		gr, err := gzip.NewReader(r)
		if err == nil {
			err = a(gr)
		}
		if err == nil {
			err = gr.Close()
		}
		return err
	}
}

func (a Reader) DecompressIf(p bool) Reader {
	if p {
		return a.Decompress()
	}
	return a
}

func (a Reader) ForEachLine(handler strings.StringHandler) Reader {
	return func(r io.Reader) error {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if err := handler(scanner.Text()); err != nil {
				return err
			}
		}
		return nil
	}
}

// Gob will read a struct/value from the reader using the
// encoding/gob package.
func (a Reader) Gob(e any) Reader {
	return a.Then(func(r io.Reader) error {
		return gob.NewDecoder(r).Decode(e)
	})
}
