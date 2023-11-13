package geiger

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	mq "github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/store/broker"
	"go.bug.st/serial"
)

type Geiger struct {
	Cron             *cron.CronService     `kernel:"inject"`
	DatabaseBroker   broker.DatabaseBroker `kernel:"inject"`
	Amqp             mq.Pool               `kernel:"inject"`
	Id               *string               `kernel:"flag,geiger-id,Geiger ID"`
	Port             *string               `kernel:"flag,geiger-port,Geiger Serial Port"`
	BaudRate         *int                  `kernel:"flag,geiger-baud,Geiger Baud Rate,115200"`
	Realtime         *int                  `kernel:"flag,geiger-realtime,Geiger in realtime mode report duration in seconds"`
	Debug            *bool                 `kernel:"flag,geiger-debug,Debug protocol"`
	port             serial.Port           // Connected serial port
	realtimeReadings []CpmReading          // Readings in realtime mode
}

func (m *Geiger) PostInit() error {
	if *m.Id == "" {
		return errors.New("-geiger-id required")
	}

	if *m.Port == "" {
		return errors.New("-geiger-port required")
	}

	switch *m.BaudRate {
	case 1200, 2400, 4800, 9600, 14400, 19200, 28800, 38400, 57600, 115200:
	default:
		return fmt.Errorf("-geiger-baud %d invalid", *m.BaudRate)
	}

	return nil
}

func (m *Geiger) Start() error {
	// This sets the baud rate of the connection.
	// This is the default for the GMC-300 Plus Geiger counter.
	// You might have to change this for other models.
	// See GQ-RFC1201.txt for details.
	mode := &serial.Mode{
		BaudRate: *m.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: 0,
	}

	port, err := serial.Open(*m.Port, mode)
	if err != nil {
		return err
	}
	m.port = port

	// Log application version
	log.Println(version.Version)

	if *m.Realtime > 0 {
		return m.runRealtime()
	}
	return m.run10s()
}

// Poll the Geiger counter every 10 seconds
func (m *Geiger) run10s() error {
	log.Println("Polling every 10s")
	err := m.heartBeat(false)
	if err == nil {
		_, err = m.Cron.AddTask("0/10 * * * * *", m.getStats)
	}
	return err
}

func (m *Geiger) runRealtime() error {
	log.Println("Running in realtime mode")
	err := m.heartBeat(true)
	if err == nil {
		err = m.realtime()
	}
	return err
}
