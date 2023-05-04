package ybsc

import (
	"compress/gzip"
	"encoding/gob"
	"github.com/soniakeys/unit"
	"io"
)

// Catalog contains just the Entries of the Bright Star catalog
type Catalog []Entry

// Entry represents an entry within the binary BSC5.bin file
type Entry struct {
	RA  unit.RA    // RA J2000
	Dec unit.Angle // Dec J2000
	Mag int16      // Visual Magnitude, unit=0.01 mag
}

func (e Entry) IsValid() bool {
	return !(e.Mag == 0 && e.RA == 0.0 && e.Dec == 0.0)
}

func ReadCatalog(r io.Reader) (Catalog, error) {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	var cat Catalog
	dec := gob.NewDecoder(gr)
	err = dec.Decode(&cat)
	return cat, err
}

func (c Catalog) Write(w io.Writer) error {
	return gob.NewEncoder(w).Encode(c)
}
