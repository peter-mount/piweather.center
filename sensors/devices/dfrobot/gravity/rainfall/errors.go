package rainfall

import "errors"

var (
	hourRangeError = errors.New("hour range error, must be 1...24")
)

func IsHourRangeError(err error) bool {
	return err == hourRangeError
}
