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

// ForEachLine will call a ReaderHandler function for each lf or crlf terminated line
// from the file.
//
// This is shorthand for Reader.ForEach(bufio.ScanLines, handler)
func (a Reader) ForEachLine(handler ReaderHandler) Reader {
	return a.ForEach(bufio.ScanLines, handler)
}

// ForEachRecord will call a ReaderHandler function for each STX-ETX delimited record
// from the file.
//
// This is shorthand for Reader.ForEach(ScanStxEtxRecord, handler)
func (a Reader) ForEachRecord(handler ReaderHandler) Reader {
	return a.ForEach(ScanStxEtxRecord, handler)
}

// ForEach will call a ReaderHandler function for each token returned by a bufio.Scanner
// running over the file. The scanner will use the supplied bufio.SplitFunc to determine
// the token's passed to the ReaderHandler.
func (a Reader) ForEach(splitFunc bufio.SplitFunc, handler ReaderHandler) Reader {
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

// ScanStxEtxRecord is a bufio.SplitFunc with looks for a record which is
// within an STX (0x02) and ETX (0x03) characters
func ScanStxEtxRecord(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	s := bytes.IndexByte(data, stx)
	e := bytes.IndexByte(data, etx)
	if s >= 0 {
		if e > s {
			// We have a full stx-etx record.
			return e + 1, data[s+1 : e], nil
		}
	}
	if atEOF {
		// at eof, we have a stx but not etx then fail
		return len(data), nil, nil
	}
	// Request more data.
	return 0, nil, nil
}

// ScanStxEtxCombiRecord is a bufio.SplitFunc which can handle a mix of plain lines and STX-ETX records.
// This is used instead of ScanStxEtcRecord or bufio.ScanLines if the content can be mixed.
//
// e.g. a log file was originally plain lines then changed to STX-ETX record format.
//
// The tokens
func ScanStxEtxCombiRecord(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	n := bytes.IndexByte(data, lfByte)
	s := bytes.IndexByte(data, stx)
	e := bytes.IndexByte(data, etx)

	if atEOF {
		// No STX in data then return the final non-terminated line
		if s == -1 {
			return len(data), dropCR(data), nil
		}

		// we have a stx then fail
		return len(data), nil, nil
	}

	switch {

	// stx at start but no etx then request more data
	case s == 0 && e == -1:
		return 0, nil, nil

	// stx at start and we have a full record
	case s == 0 && e > s:
		return e + 1, data[s+1 : e], nil

	// We have a line feed and no stx or stx but after it then we have a plain line
	case n >= 0 && (s == -1 || n < s):
		return n + 1, dropCR(data[0:n]), nil

	// We have a stx but either no new line or it's after the stx then plain line up to the stx
	case s > 0 && (n == -1 || n > s):
		return s, data[:s], nil

		// Request more data.
	default:
		return 0, nil, nil
	}
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
