package record

import (
	"encoding/binary"
	"fmt"
	"github.com/peter-mount/piweather.center/util"
	"io"
)

const (
	currentHeaderVersion     = 1          // Current version of header record
	currentFileVersion       = 1          // Current version of record format
	headerMagic              = "PIWTHRDB" // File magic, must be 8 bytes!
	headerMagicOffset        = 0
	headerSizeOffset         = headerMagicOffset + 8        // Offset of header size - 2 bytes
	headerVersionOffset      = headerSizeOffset + 2         // 1 byte
	headerFileVersionOffset  = headerVersionOffset + 1      // 1 byte
	headerMetricLengthOffset = headerFileVersionOffset + 1  // 2 byte
	headerMetricOffset       = headerMetricLengthOffset + 2 // Offset of metric
)

var (
	// handlers for supported file versions
	version1Handler = RecV1{}
	// default handler for new files, should always be the latest version.
	currentHandler = RecV1{}
)

func NewFileHeader(metric string) FileHeader {
	size, _ := headerSize(metric)
	return FileHeader{
		Name:    metric,
		Version: currentFileVersion,
		Size:    size,
	}
}

func CurrentHandler() RecordHandler {
	return currentHandler
}

type FileHeader struct {
	Size          int    // Header size
	HeaderVersion int    // Header version
	Version       int    // File format version
	Name          string // Name of metric this file contains
}

func (h *FileHeader) GetRecordHandler() (RecordHandler, error) {
	if h.Version == 0 {
		h.Version = currentFileVersion
	}
	switch h.Version {
	case 1:
		return version1Handler, nil
	default:
		return nil, fmt.Errorf("unsupported file version %d", h.Version)
	}
}

func headerSize(metric string) (int, int) {
	name := []byte(metric)
	nameLen := len(name)

	// Calculate header size, pad to nearest 16 byte boundary
	size := headerMetricOffset + nameLen
	return size + 16 - (size & 0x0f), nameLen
}

func (h *FileHeader) Write(w io.Writer) error {
	size, nameLen := headerSize(h.Name)

	b := append([]byte{}, headerMagic...) // 8 bytes
	b = binary.LittleEndian.AppendUint16(b, uint16(size))
	b = append(b, currentHeaderVersion, currentFileVersion)
	b = binary.LittleEndian.AppendUint16(b, uint16(nameLen))
	b = append(b, h.Name...)

	// Should not be needed but pad to fit expected size
	b, err := util.PadToLength(b, size)
	if err != nil {
		// Can happen if metric name is too long
		return err
	}
	// safety check: Check our header is of the correct size
	if len(b) != size {
		return fmt.Errorf("bug: Header size exceeding expected %d bytes, have %d", size, len(b))
	}

	n, err := w.Write(b)
	if err == nil && n != len(b) {
		return fmt.Errorf("failed to write header for %q", h.Name)
	}
	return err
}

func (h *FileHeader) Read(r io.Reader) error {
	// read the main part of the header minus the name
	b := make([]byte, headerMetricOffset)

	n, err := r.Read(b)
	if err != nil {
		return err
	}
	if n != headerMetricOffset {
		return fmt.Errorf("invalid header, expected %d bytes got %d", headerMetricOffset, n)
	}

	// Read in the default data from the header
	h.Size = int(binary.LittleEndian.Uint16(b[headerSizeOffset : headerSizeOffset+2]))
	h.HeaderVersion = int(b[headerVersionOffset])
	h.Version = int(b[headerFileVersionOffset])
	nameLen := int(binary.LittleEndian.Uint16(b[headerMetricLengthOffset : headerMetricLengthOffset+2]))

	// Now read in the metric name
	b = make([]byte, nameLen)
	n, err = r.Read(b)
	if err != nil {
		return err
	}
	if n != nameLen {
		return fmt.Errorf("invalid header, expected %d bytes got %d", nameLen, n)
	}
	h.Name = string(b)

	return nil
}
