package smbus

import "sync"

type Manager struct {
	mutex sync.Mutex
	bus   map[int]*i2cBus
}

func (m *Manager) Start() error {
	m.bus = make(map[int]*i2cBus)
	return nil
}

func (m *Manager) getDevice(bus int, i2cAddr uint8) (*i2cBus, *smBus) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	handler, ok := m.bus[bus]
	if !ok {
		handler = &i2cBus{
			bus:     bus,
			devices: make(map[uint8]*smBus),
		}
		m.bus[bus] = handler
	}

	device, ok := handler.devices[i2cAddr]
	if !ok {
		device = &smBus{
			bus:     handler,
			i2cAddr: i2cAddr,
		}
	}

	return handler, device
}

func (m *Manager) UseDevice(bus int, i2cAddr uint8, task Task) error {
	handler, device := m.getDevice(bus, i2cAddr)
	return handler.handle(device, task)
}
