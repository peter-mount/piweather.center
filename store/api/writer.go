package api

import (
	"encoding/binary"
	"errors"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"math"
	"time"
)

var (
	order      = binary.LittleEndian
	shortWrite = errors.New("short write")
)

func newWriter(wr io.Writer) *writer {
	return &writer{w: wr, b: make([]byte, 16)}
}

type writer struct {
	w io.Writer // Writer to write to
	b []byte    // scratch byte slice
}

func (w *writer) writeX(data any) error {
	return binary.Write(w.w, order, data)
}

func (w *writer) write(b []byte) error {
	l := len(b)
	n, err := w.w.Write(b)
	if err == nil && n != l {
		err = shortWrite
	}
	return err
}

func (w *writer) writeB(l int) error {
	return w.write(w.b[:l])
}

func (w *writer) bool(b bool) error {
	if b {
		return w.uint8(1)
	}
	return w.uint8(0)
}

func (w *writer) int8(i int8) error {
	return w.uint8(uint8(i))
}

func (w *writer) int16(i int16) error {
	return w.uint16(uint16(i))
}

func (w *writer) int32(i int32) error {
	return w.uint32(uint32(i))
}

func (w *writer) int64(i int64) error {
	return w.uint64(uint64(i))
}

func (w *writer) uint8(i uint8) error {
	w.b[0] = i
	return w.writeB(1)
}

func (w *writer) uint16(i uint16) error {
	order.PutUint16(w.b, i)
	return w.writeB(2)
}

func (w *writer) uint32(i uint32) error {
	order.PutUint32(w.b, i)
	return w.writeB(4)
}

func (w *writer) uint64(i uint64) error {
	order.PutUint64(w.b, i)
	return w.writeB(8)
}

func (w *writer) string(s string) error {
	b := []byte(s)
	l := uint16(len(b))
	err := w.uint16(l)

	if err == nil && l > 0 {
		err = w.write(b)
	}

	return err
}

func (w *writer) float64(f float64) error {
	return w.uint64(math.Float64bits(f))
}

func (w *writer) value(v value.Value) error {

	valid := v.IsValid()
	err := w.bool(valid)
	if err != nil || !valid {
		return err
	}

	b := util.AppendValue([]byte{}, v)
	return w.write(b)
}

func (w *writer) time(t time.Time) error {
	// Is t valid (e.g. not zero)
	valid := !t.IsZero()
	err := w.bool(valid)

	if err == nil && valid {
		err = w.int64(t.UTC().Unix())
	}

	return err
}

func (w *writer) duration(d time.Duration) error {
	return w.int64(int64(d))
}
