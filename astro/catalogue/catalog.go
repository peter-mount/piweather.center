package catalogue

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/chart"
	"io"
	"sort"
)

// Catalog contains just the Entries of the Bright Star catalog
type Catalog struct {
	size  int
	stars []Entry
}

// Equals checks that two Catalog instances contain the same data
func (c *Catalog) Equals(b *Catalog) bool {
	if c == nil {
		return b == nil
	}
	if c.size != b.size {
		return false
	}
	for i, v := range c.stars {
		if !v.Equals(b.stars[i]) {
			return false
		}
	}
	return true
}

// Write will write a Catalog to an io.Writer
func (c *Catalog) Write(w io.Writer) error {
	n, err := w.Write([]byte{byte((c.size >> 8) & 0xff), byte(c.size & 0xff)})
	if err != nil {
		return err
	}
	if n != 2 {
		return errors.New("failed to write length")
	}

	for i, e := range c.stars {
		dec := e.dec
		if dec < 0 {
			dec = -dec
			dec = dec | 0x800000
		}
		n, err := w.Write([]byte{
			byte((e.ra >> 16) & 0xff), byte((e.ra >> 8) & 0xff), byte(e.ra & 0xff),
			byte((dec >> 16) & 0xff), byte((dec >> 8) & 0xff), byte(dec & 0xff),
			byte((e.mag >> 8) & 0xff), byte(e.mag & 0xff),
		})
		if err != nil {
			return err
		}
		if n != 8 {
			return fmt.Errorf("failed to write star %d", i)
		}
	}

	return nil
}

// Read will read a Catalog from an io.Reader
func (c *Catalog) Read(r io.Reader) error {
	b := []byte{0, 0}
	n, err := r.Read(b)
	if err != nil {
		return err
	}
	if n != 2 {
		return errors.New("failed to read length")
	}
	size := (int(b[0]) << 8) | int(b[1])

	log.Printf("expecting %d stars", size)

	c.stars = nil
	b = make([]byte, 8)
	for i := 0; i < size; i++ {
		n, err = io.ReadAtLeast(r, b, 8)
		if err != nil {
			log.Printf("err: %d %d %d", c.size, len(b), n)
			return err
		}
		if n != 8 {
			return fmt.Errorf("failed to read star got %d", n)
		}

		ra := (int(b[0]) << 16) | (int(b[1]) << 8) | (int(b[2]))
		de := (int(b[3]&0x7f) << 16) | (int(b[4]) << 8) | (int(b[5]))
		if (b[3] & 0x80) == 0x80 {
			de = -de
		}
		c.Append(Entry{
			ra:  int32(ra),
			dec: int32(de),
			mag: int16((int(b[6]) << 8) | int(b[7])),
		})
	}

	if c.size != size {
		return fmt.Errorf("expected %d stars, read %d", size, c.size)
	}

	return nil
}

// Append an Entry to the Catalog
func (c *Catalog) Append(e Entry) {
	c.stars = append(c.stars, e)
	c.size = len(c.stars)
}

// Size of the Catalog
func (c *Catalog) Size() int {
	return c.size
}

// Get a specific Entry
func (c *Catalog) Get(i int) Entry {
	return c.stars[i]
}

// ForEach will cann Handler for each Entry in the Catalog
func (c *Catalog) ForEach(h Handler) error {
	for _, e := range c.stars {
		if err := h(e); err != nil {
			return err
		}
	}
	return nil
}

// Sort will sort the Catalog in order of Magnitude, faintest first.
// This ordering is preferable when plotting charts
func (c *Catalog) Sort() {
	sort.SliceStable(c.stars, func(i, j int) bool {
		return c.stars[i].mag > c.stars[j].mag
	})
}

func (c *Catalog) NewLayer(renderer StarRenderer, proj chart.Projection) CatalogLayer {
	return NewCatalogLayer(c, renderer, proj)
}
