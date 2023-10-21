package record

import (
	"github.com/peter-mount/piweather.center/util"
)

// RecV1 implements record format version 1
type RecV1 struct {
}

const (
	revV1Size = 24
)

func (r RecV1) Size() int {
	return revV1Size
}

func (r RecV1) Read(b []byte) (Record, error) {
	var rec Record

	err := AssertRecordLength(r, b)

	if err == nil {
		rec.Time, err = util.ReadTime(b[0:8])
	}

	if err == nil {
		rec.Value, err = util.ReadValue(b[8:24])
	}

	return rec, err
}

func (r RecV1) Append(b []byte, rec Record) []byte {
	b = util.AppendTime(b, rec.Time)
	b = util.AppendValue(b, rec.Value)
	return b
}
