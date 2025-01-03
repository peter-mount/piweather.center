package i2c

import "sync"

var (
	mutex   sync.Mutex
	devices map[address]*device
)

func init() {
	devices = make(map[address]*device)
}

type address struct {
	bus  int   // The i2c bus for this device
	addr uint8 // The i2c address of the device on the bus
}

type device struct {
	address
	mutex sync.Mutex
}

func (d *device) execute(task Task) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return task.execute(d.bus, d.addr)
}

func getDevice(bus int, addr uint8) *device {
	mutex.Lock()
	defer mutex.Unlock()

	key := address{bus: bus, addr: addr}

	dev, ok := devices[key]
	if !ok {
		dev = &device{address: key}
		devices[key] = dev
	}

	return dev
}

// UseI2CConcurrent will call a Task against a specific I2C device, ensuring that only one Task is running against
// that specific device at any one time.
//
// If task is nil this does nothing.
//
// Note: The Task is specific for a single device.
// It MUST NOT attempt to access any other devices as that will almost certainly cause a deadlock.
func UseI2CConcurrent(bus int, addr uint8, task Task) error {
	if task == nil {
		return nil
	}
	return getDevice(bus, addr).execute(task)
}
