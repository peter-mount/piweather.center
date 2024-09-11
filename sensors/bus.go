package sensors

type BusType uint8

func (b BusType) Label() string {
	return busLabels[b]
}

const (
	// BusI2C represents I2C devices
	BusI2C BusType = iota
	// BusSPI represents SPI devices
	BusSPI
	// BusSerial represents devices used over serial, either TTL, RS232, RS423 etc
	BusSerial
)

var (
	busLabels = []string{"I2C", "SPI", "Serial"}
)
