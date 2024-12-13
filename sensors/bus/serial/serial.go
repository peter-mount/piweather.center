package serial

import "go.bug.st/serial"

type Serial interface {
	// Read a []byte from the Serial device
	Read([]byte) (int, error)

	// Write a []byte to the Serial device
	Write([]byte) (int, error)

	// WriteByte writes a single byte to the Serial device
	WriteByte(byte) error
}

type serialDevice struct {
	// Connected serial port
	port serial.Port
	// Copy of the mode used on the port
	mode *serial.Mode
}

func New(portName string, mode *serial.Mode) (Serial, error) {
	port, err := serial.Open(portName, mode)
	if err != nil {
		return nil, err
	}

	return &serialDevice{
		port: port,
		mode: mode,
	}, nil

}
