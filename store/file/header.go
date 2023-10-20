package file

import (
	"fmt"
	"io"
)

const (
	currentHeaderVersion     = 1          // Current version of header record
	currentFileVersion       = 1          // Current version of record format
	headerSize               = 256        // Size of header block
	headerMagic              = "PIWTHRDB" // File magic, must be 8 bytes!
	headerMagicOffset        = 0
	headerVersionOffset      = headerMagicOffset + 8
	headerFileVersionOffset  = headerVersionOffset + 1
	headerMetricLengthOffset = headerFileVersionOffset + 1
	headerMetricOffset       = 16
)

var (
	// handlers for supported file versions
	version1Handler = RecV1{}
	// default handler for new files, should always be the latest version.
	currentHandler = RecV1{}
)

type FileHeader struct {
	Version int    // File format version
	Name    string // Name of metric this file contains
}

func (h *FileHeader) write(w io.Writer) error {
	b := append([]byte{}, headerMagic...)     // 8 bytes
	b, _ = PadToLength(b, headerMetricOffset) // Name starts at byte 16
	b[headerVersionOffset] = currentHeaderVersion
	b[headerFileVersionOffset] = currentFileVersion
	b[headerMetricLengthOffset] = byte(len(h.Name))
	b = append(b, h.Name...)

	b, err := PadToLength(b, headerSize) // Pad so header is fully occupied
	if err != nil {
		// Can happen if metric name is too long
		return err
	}

	n, err := w.Write(b)
	if err == nil && n != len(b) {
		return fmt.Errorf("failed to write header for %q", h.Name)
	}
	return err
}

func (h *FileHeader) read(r io.Reader) error {
	b := make([]byte, headerSize)

	n, err := r.Read(b)
	if err != nil {
		return err
	}
	if n != headerSize {
		return fmt.Errorf("invalid header, expected %d bytes got %d", headerSize, n)
	}

	// we currently ignore headerVersionOffset
	h.Version = int(b[headerFileVersionOffset])
	nameLen := int(b[headerMetricLengthOffset])
	h.Name = string(b[headerMetricOffset : headerMetricOffset+nameLen])

	return nil
}
