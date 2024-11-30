package state

import (
	"crypto/rand"
	"crypto/sha1"
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

func compress(b []byte) []byte {
	// Now compress by xor each 2 bytes into 1. do this multiple times to create a short
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
