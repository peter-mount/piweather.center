package serial

func (s *serialDevice) Read(p []byte) (int, error) {
	return s.port.Read(p)
}
