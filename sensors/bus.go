package sensors

import "errors"

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
	busLabels            = []string{"I2C", "SPI", "Serial"}
	deviceNotFound       = errors.New("device not found")
	deviceNotImplemented = errors.New("device not implemented")
)

type PollMode uint8

func (b PollMode) Label() string {
	return pollLabels[b]
}

const (
	// PollReading indicates device should be called at regular intervals
	PollReading PollMode = iota
	// PushReading indicates the device will submit readings itself in real time
	PushReading
)

var pollLabels = []string{"Poll", "Push"}
