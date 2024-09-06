package i2c

const (
	// ioctl commands
	i2cSlave      = 0x0703
	i2cSlaveForce = 0x0706
	i2cFuncs      = 0x0705
	i2cSMBus      = 0x0720

	i2cSMBusWrite uint8 = 0
	i2cSMBusRead  uint8 = 1

	i2cSMBusI2CBlockData uint32 = 8
	i2cSMBusBlockMax     uint32 = 32
)
