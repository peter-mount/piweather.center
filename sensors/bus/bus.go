package bus

type BusType uint8

func (b BusType) Label() string {
	return busLabels[b]
}

const (
	// Used to ensure the entry is set in DeviceInfo
	busUndefined BusType = iota
	// BusI2C represents I2C devices
	BusI2C
	// BusSPI represents SPI devices
	BusSPI
	// BusSerial represents devices used over serial, either TTL, RS232, RS423 etc
	BusSerial
)

type PollMode uint8

func (b PollMode) Label() string {
	return pollLabels[b]
}

const (
	// Used to ensure the entry is set in DeviceInfo
	pollUndefined PollMode = iota
	// PollReading indicates device should be called at regular intervals
	PollReading
	// PushReading indicates the device will submit readings itself in real time
	PushReading
)

var (
	busLabels  = []string{"", "I2C", "SPI", "Serial"}
	pollLabels = []string{"", "Poll", "Push"}
)
