package i2c

import "errors"

var (
	errSmbusBlockSize = errors.New("invalid SMBus block size")
)

func IsSmbusBlockSize(err error) bool {
	return err == errSmbusBlockSize
}
