package model

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/peter-mount/piweather.center/store/api"
	"math"
)

// Unique ID for this instance, changes every time the application runs
var (
	uid  []byte
	uids string
)

func init() {
	// Generate a unique id for this instance
	uid = make([]byte, sha1.Size)
	if _, err := rand.Read(uid); err != nil {
		panic(err)
	}

	// Generate a compressed uid string
	uids = string(encode(compress(uid)))
}

func UID() string {
	return uids
}

// Dashboard is the top level Component representing an entire dashboard.
// Initially it's the same as Container, however it will have additional fields in the future.
type Dashboard struct {
	Container `yaml:",inline"`
	Live      bool   `yaml:"live,omitempty"` // If true then dashboard can have live updates
	Uid       string `yaml:"-"`              // Uuid of dashboard - generated
	idSeq     int16  // Used in initialising the ID's
}

func DashboardFactory() any {
	return &Dashboard{}
}

func UnmarshalDashboard(b []byte, o any) error {
	if err := Unmarshal(b, o); err != nil {
		return err
	}

	d, ok := o.(*Dashboard)
	if !ok {
		return fmt.Errorf("expected *Dashboard got %T", o)
	}

	// Now generate a unique id for the Dashboard.
	// This allows the frontend to refresh if the backend restarts but the dashboard
	// hasn't changed as the uuid sent to the front end when an update occurs will be different.

	// initial id is the instance uid appended with the sha1 sum of the dashboard's bytes
	var id []byte
	id = append(id, uid...)

	// Note: we cannot use id=append(id,sha1.Sum(b)...) here as it returns [20]byte instead of
	// []byte, and go does not allow that hence we manually append the bytes
	sum := sha1.Sum(b)
	for i := 0; i < len(sum); i++ {
		id = append(id, sum[i])
	}

	id = compress(id)

	d.Uid = string(encode(id))

	// Initialise the Dashboard's components, so they get their own unique ID's
	d.idSeq = 0
	d.Container.init(d)

	return nil
}

func compress(b []byte) []byte {
	// Now compress by xor each 2 bytes into 1. Do this multiple times to create a short
	// final uid that should be unique but not take up much space as this
	// will be sent on every metric update, so we don't want it to be too long.
	for len(b) > 8 {
		var c []byte
		l := len(b)
		for i := 0; i < l; i += 2 {
			if (i + 1) < l {
				c = append(c, b[i]^b[i+1])
			} else {
				c = append(c, b[i])
			}
		}
		b = c
	}
	return b
}

const (
	base    = 62
	charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func encode(b []byte) []byte {
	var o []byte

	l := len(b)
	for i := 0; i < l; i += 8 {
		var n uint64
		for j := 0; j < 8 && (i+j) < l; j++ {
			n <<= 8
			n |= uint64(b[i+j])
		}

		for n > 0 {
			r := int(math.Mod(float64(n), float64(base)))
			n /= base
			o = append(o, charSet[r])
		}
	}

	return o
}

func (c *Dashboard) NextId() string {
	c.idSeq++
	return string(encode([]byte{uid[0], uid[1], byte(c.idSeq & 0xcff), byte((c.idSeq >> 8) & 0xcff)}))
}

func (c *Dashboard) Process(m api.Metric, r *Response) {
	// Set the response Uuid to the Dashboard.
	// This allows the front end to detect a dashboard change.
	r.Uid = c.Uid
	c.Container.Process(m, r)
}
