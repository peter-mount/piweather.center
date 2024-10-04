package api

import (
	"errors"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"time"
)

var (
	shortRead = errors.New("short read")
)

type Readable interface {
	Read(io.Reader) error
}

func (r *Result) Read(src io.Reader) error {
	rdr := &reader{r: src, b: make([]byte, 16)}
	return r.read(rdr)
}

type reader struct {
	r io.Reader
	b []byte // scratch byte slice
}

func (r *reader) readB(l int) error {
	if l > len(r.b) {
		r.b = make([]byte, l)
	}
	b1 := r.b[:l]
	n, err := r.r.Read(b1)
	if err == nil && n < l {
		err = shortRead
	}
	if n == l && err == io.EOF {
		err = nil
	}

	return err
}

func (r *reader) uint8() (uint8, error) {
	err := r.readB(1)
	return r.b[0], err
}

func (r *reader) uint16() (uint16, error) {
	err := r.readB(2)
	if err != nil {
		return 0, err
	}

	return order.Uint16(r.b[:2]), nil
}

func (r *reader) uint32() (uint32, error) {
	err := r.readB(4)
	if err != nil {
		return 0, err
	}

	return order.Uint32(r.b[:4]), nil
}

func (r *reader) uint64() (uint64, error) {
	err := r.readB(8)
	if err != nil {
		return 0, err
	}

	return order.Uint64(r.b[:8]), nil
}

func (r *reader) bool() (bool, error) {
	v, err := r.uint8()
	return v != 0, err
}

func (r *reader) int8() (int8, error) {
	v, err := r.uint8()
	return int8(v), err
}

func (r *reader) int16() (int16, error) {
	v, err := r.uint16()
	return int16(v), err
}

func (r *reader) int32() (int32, error) {
	v, err := r.uint32()
	return int32(v), err
}

func (r *reader) int64() (int64, error) {
	v, err := r.uint64()
	return int64(v), err
}

func (r *reader) string() (string, error) {
	l, err := r.uint16()
	if err != nil || l == 0 {
		return "", err
	}

	err = r.readB(int(l))
	if err != nil {
		return "", err
	}

	return string(r.b[:l]), nil
}

func (r *reader) value() (value.Value, error) {
	valid, err := r.bool()
	if err != nil || !valid {
		return value.Value{}, err
	}

	err = r.readB(16)
	if err != nil {
		return value.Value{}, err
	}
	return util.ReadValue(r.b[:16])
}

func (r *reader) time() (time.Time, error) {
	// Valid time in stream?
	if b, err := r.bool(); err != nil || !b {
		// No or error then a zero time returned
		return time.Time{}, err
	}

	t, err := r.int64()
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(t, 0).UTC(), nil
}

func (r *reader) duration() (time.Duration, error) {
	t, err := r.int64()
	if err != nil {
		return 0, err
	}
	return time.Duration(t), nil
}
