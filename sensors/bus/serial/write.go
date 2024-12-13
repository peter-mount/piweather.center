package serial

import "io"

func (s *serialDevice) Write(data []byte) (int, error) {
	return s.port.Write(data)
}

// WriteByte sends a single byte to the remote i2c device.
// The interpretation of the message is implementation dependant.
func (s *serialDevice) WriteByte(b byte) error {
	var buf [1]byte
	buf[0] = b
	n, err := s.port.Write(buf[:])
	if err == nil && n == 0 {
		err = io.ErrShortWrite
	}
	return err
}
