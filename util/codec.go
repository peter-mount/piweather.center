package util

import (
	"encoding/binary"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"time"
)

func AssertLength(s int, b []byte) error {
	bl := len(b)
	if bl < s {
		return fmt.Errorf("invalid record length, got %d expect %d", bl, s)
	}
	return nil
}

func PadToLength(b []byte, l int) ([]byte, error) {
	if len(b) < l {
		b = append(b, make([]byte, l-len(b))...)
	}
	if len(b) > l {
		return b, fmt.Errorf("invalid block length %d require %d", len(b), l)
	}
	return b, nil
}

func ReadTime(b []byte) (time.Time, error) {
	if err := AssertLength(8, b); err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(binary.LittleEndian.Uint64(b[0:8])), 0), nil
}

func AppendTime(b []byte, t time.Time) []byte {
	return binary.LittleEndian.AppendUint64(b, uint64(t.Unix()))
}

func ReadValue(b []byte) (value.Value, error) {
	if err := AssertLength(16, b); err != nil {
		return value.Value{}, err
	}
	fb := binary.LittleEndian.Uint64(b[0:8])
	ub := binary.LittleEndian.Uint64(b[8:16])

	u, ok := value.GetUnitByHash(ub)
	if !ok {
		return value.Value{}, fmt.Errorf("invalid Value, unknown unit %d", ub)
	}

	return u.Value(math.Float64frombits(fb)), nil
}

func AppendValue(b []byte, v value.Value) []byte {
	b = binary.LittleEndian.AppendUint64(b, math.Float64bits(v.Float()))
	b = binary.LittleEndian.AppendUint64(b, v.Unit().Hash())

	if v.Unit().Hash() == 7599665368900986221 {
		log.Printf("*** Hash 7599665368900986221 %q %s", v.Unit().ID(), v.String())
	}
	return b
}

func AppendString(b []byte, s string) []byte {
	b = binary.LittleEndian.AppendUint16(b, uint16(len(s)))
	return append(b, s...)
}

func ReadString(b []byte) (string, []byte) {
	l := int(binary.LittleEndian.Uint16(b[0:2]))
	s := string(b[2 : 2+l])
	return s, b[2+l:]
}
