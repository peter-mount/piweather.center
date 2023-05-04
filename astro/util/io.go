package util

import (
	"compress/gzip"
	"github.com/peter-mount/go-kernel/v2/log"
	"io"
	"os"
)

type Writer func(io.Writer) error

func Create(filename string, w Writer) error {
	log.Printf("Creating %s", filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return w(f)
}

func Compress(w Writer) Writer {
	return func(w1 io.Writer) error {
		gw := gzip.NewWriter(w1)
		if err := w(gw); err != nil {
			return err
		}
		return gw.Close()
	}
}
