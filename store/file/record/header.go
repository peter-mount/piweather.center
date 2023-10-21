package record

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/peter-mount/piweather.center/util"
	"io"
)

const (
	currentHeaderVersion     = 1          // Current version of header record
	currentFileVersion       = 1          // Current version of record format
	headerMagic              = "PIWTHRDB" // File magic, must be 8 bytes!
	headerMagicOffset        = 0
	headerSizeOffset         = headerMagicOffset + 8        // 2 bytes
	headerVersionOffset      = headerSizeOffset + 2         // 2 bytes
	headerFileVersionOffset  = headerVersionOffset + 2      // 2 bytes
	headerRecordLengthOffset = headerFileVersionOffset + 2  // 2 bytes
	headerMetricLengthOffset = headerRecordLengthOffset + 2 // 2 bytes
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
		Name:          metric,
		Size:          size,
		HeaderVersion: currentHeaderVersion,
		RecordVersion: currentFileVersion,
		RecordLength:  currentHandler.Size(),
	}
}

func CurrentHandler() RecordHandler {
	return currentHandler
}

type FileHeader struct {
	Size          int    // Header size
	HeaderVersion int    // Header version
	RecordVersion int    // File format version
	RecordLength  int    // Record length
	Name          string // Name of metric this file contains
}

func (h *FileHeader) GetRecordHandler() (RecordHandler, error) {
	if h.RecordVersion == 0 {
		h.RecordVersion = currentFileVersion
	}
	switch h.RecordVersion {
	case 1:
		return version1Handler, nil
	default:
		return nil, fmt.Errorf("unsupported file version %d", h.RecordVersion)
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
	b = binary.LittleEndian.AppendUint16(b, uint16(h.HeaderVersion))
	b = binary.LittleEndian.AppendUint16(b, uint16(h.RecordVersion))
	b = binary.LittleEndian.AppendUint16(b, uint16(h.RecordLength))
	b = binary.LittleEndian.AppendUint16(b, uint16(nameLen))
	b = append(b, h.Name...)

	// pad to fit expected size as we keep records working on a 16 byte boundary
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

	// Ensure this is an actual DB file by checking the headerMagic
	//log.Printf("%X", b[headerMagicOffset:headerMagicOffset+8])
	if string(b[headerMagicOffset:headerMagicOffset+8]) != headerMagic {
		return errors.New("not a db file")
	}

	// Read in the default data from the header
	h.Size = int(binary.LittleEndian.Uint16(b[headerSizeOffset : headerSizeOffset+2]))
	h.HeaderVersion = int(binary.LittleEndian.Uint16(b[headerVersionOffset : headerVersionOffset+2]))
	h.RecordVersion = int(binary.LittleEndian.Uint16(b[headerFileVersionOffset : headerFileVersionOffset+2]))
	h.RecordLength = int(binary.LittleEndian.Uint16(b[headerRecordLengthOffset : headerRecordLengthOffset+2]))
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
