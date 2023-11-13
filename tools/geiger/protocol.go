package geiger

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"math"
	"time"
)

const (
	getCpm         = "GETCPM"
	getTemp        = "GETTEMP"
	getVolt        = "GETVOLT"
	getGyro        = "GETGYRO"
	InvalidCPM     = -1
	InvalidReading = -math.MaxFloat64
)

func (m *Geiger) sendCmdf(len int, cmd string, args ...interface{}) ([]byte, error) {
	return m.sendCmd(fmt.Sprintf(cmd, args...), len)
}

// sendCmd sends a command to the Geiger Counter.
// cmd is the text content of the command without the prefix or suffix characters.
// len is the number of bytes expected from the command.
// len=0 if no response is expected.
func (m *Geiger) sendCmd(cmd string, len int) ([]byte, error) {
	c := fmt.Sprintf("<%s>>%c", cmd, 13)
	if *m.Debug {
		log.Printf("sending %q", c)
	}
	n, err := m.port.Write([]byte(c))
	if err != nil {
		return nil, err
	}
	if *m.Debug {
		log.Printf("Sent %d", n)
	}

	time.Sleep(time.Millisecond * 250)

	if len <= 0 {
		return nil, nil
	}

	if *m.Debug {
		log.Printf("Reading %d", len)
	}
	buf := make([]byte, len)

	n, err = m.port.Read(buf)
	if err != nil {
		return nil, err
	}
	if *m.Debug {
		log.Printf("Read %d/%d", n, len)
	}

	if n < len {
		return nil, fmt.Errorf("expected %d got %d (%v)", len, n, buf)
	}

	return buf, nil
}

// toInt16 converts a 2 byte value to an unsigned integer
func toInt16(b []byte) int {
	return (int(b[0]) * 256) + int(b[1])
}

// toFloat64 converts a 4 byte value to a signed float64.
// Byte 0 contains the integer portion.
// Byte 1 contains the fraction
// Byte 2 if not 0 indicates the value is negative.
// Byte 3 is not referenced here but is not used in the protocol
func toFloat64(b []byte) float64 {
	temp := float64(b[0]) + (float64(b[1]) / 100)
	if b[2] != 0 {
		temp = -temp
	}
	return temp
}

func (m *Geiger) getCpm() int {
	b, err := m.sendCmd(getCpm, 2)
	if err != nil {
		return InvalidCPM
	}
	return toInt16(b)
}

func (m *Geiger) getTemp() float64 {
	b, err := m.sendCmd(getTemp, 4)
	if err != nil {
		return InvalidReading
	}
	return toFloat64(b)
}

func (m *Geiger) getVolt() float64 {
	b, err := m.sendCmd(getVolt, 1)
	if err != nil {
		return InvalidReading
	}
	return float64(b[0]) / 10.0
}

func (m *Geiger) getGyro() (int, int, int) {
	b, err := m.sendCmd(getGyro, 7)
	if err != nil {
		return 0, 0, 0
	}
	return toInt16(b[0:2]), toInt16(b[2:4]), toInt16(b[4:6])
}

func (m *Geiger) setClock() {
	f := []string{
		"SETDATEYY", "SETDATEMM", "SETDATEDD",
		"SETTIMEHH", "SETTIMEMM", "SETTIMESS",
	}
	now := time.Now()
	v := []int{
		now.Year() % 100,
		int(now.Month()),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	}

	for i, e := range v {
		b, err := m.sendCmdf(1, f[i]+"%02x", e)
		if err != nil || len(b) != 1 || b[0] != 0xAA {
			log.Printf("Set clock failed %d %d", i, e)
		}
	}

	log.Println("Set Geiger Clock success")
}

func (m *Geiger) heartBeat(enabled bool) error {
	cmd := "HEARTBEAT0"
	if enabled {
		cmd = "HEARTBEAT1"
	}
	_, err := m.sendCmd(cmd, 0)
	if err == nil {
		// Clear buffer
		err = m.port.ResetInputBuffer()
	}
	return err
}
