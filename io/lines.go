package io

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const (
	stx    = 0x02 // stx Start of Text used to mark the beginning of a record
	etx    = 0x03 // etx End of Text used to mark the end of a record
	lf     = "\n" // lf Line feed used to mark the end of a line
	lfByte = '\n' // same as lf but as a single byte
	cr     = "\r" // cr Carriage Return used with lf for Microsoft/DOS text files
	crlf   = "\r\n"
)

// ReaderHandler accepts a line or record from a Reader
type ReaderHandler func(string) error

// WriterHandler is a function called by WriteLines when writing Lines or WriteRecords when writing Records
type WriterHandler func(w LineWriter) error

type LineWriter func(s string) error

type lineWriter struct {
	w      io.Writer
	prefix byte
	suffix byte
}

// WriteLines will call a WriterHandler and each line written to the handler will be terminated with a Line Feed.
func (a Writer) WriteLines(w WriterHandler) Writer {
	return a.lineWriter(lineWriter{suffix: lfByte}, w)
}

// WriteRecords will call a WriterHandler and each line written to the handler will be wrapped with
// a stx/etx pair.
func (a Writer) WriteRecords(w WriterHandler) Writer {
	return a.lineWriter(lineWriter{prefix: stx, suffix: etx}, w)
}

func (a Writer) lineWriter(lw lineWriter, w WriterHandler) Writer {
	return func(writer io.Writer) error {
		lw.w = writer
		return w(lw.Write)
	}
}

func stripSuffix(s, suffix string) string {
	for strings.HasSuffix(s, suffix) {
		s = strings.TrimSuffix(s, suffix)
	}
	return s
}

func (lw *lineWriter) Write(s string) error {
	s = stripSuffix(s, crlf)
	s = stripSuffix(s, cr)
	s = stripSuffix(s, lf)
	var b []byte
	if lw.prefix != 0 {
		b = append(b, lw.prefix)
	}
	b = append(b, s...)
	if lw.suffix != 0 {
		b = append(b, lw.suffix)
	}
	n, err := lw.w.Write(b)
	if err == nil && n != len(b) {
		err = io.ErrShortWrite
	}
	return err
}

func (a Reader) ForEachLine(handler ReaderHandler) Reader {
	return a.forEach(bufio.ScanLines, handler)
}

func (a Reader) ForEachRecord(handler ReaderHandler) Reader {
	return a.forEach(ScanStxEtxRecord, handler)
}

func (a Reader) forEach(splitFunc bufio.SplitFunc, handler ReaderHandler) Reader {
	return func(r io.Reader) error {
		scanner := bufio.NewScanner(r)
		scanner.Split(splitFunc)
		for scanner.Scan() {
			if err := handler(scanner.Text()); err != nil {
				return err
			}
		}
		return nil
	}
}

// ScanStxEtxRecord is a scanner with looks for a record which is
// within an STX (0x02) and ETX (0x03) characters
func ScanStxEtxRecord(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if s := bytes.IndexByte(data, stx); s >= 0 {
		if e := bytes.IndexByte(data, etx); e > s {
			// We have a full stx-etx record.
			return e + 1, data[s+1 : e], nil
		}
	}
	// Request more data.
	return 0, nil, nil
}
