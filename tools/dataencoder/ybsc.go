package dataencoder

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/tools/dataencoder"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/io"
	"github.com/soniakeys/unit"
	"path/filepath"
	"strconv"
	"strings"
)

// YbscEncoder takes the raw Yale Bright Star Catalogue and created the bsc5.bin binary.
//
// You can download bsc5.dat.gz from Harvard University
// http://tdc-www.harvard.edu/catalogs/bsc5.html
type YbscEncoder struct {
	Encoder *dataencoder.Encoder `kernel:"inject"`
	Build   *dataencoder.Build   `kernel:"inject"`
	Source  *string              `kernel:"flag,bsc5,Encode bsc5.dat"`
}

func (s *YbscEncoder) Start() error {
	s.Build.AddLibProvider(s.includeYbsc)

	if *s.Source != "" {
		return s.encode()
	}
	return nil
}

func (s *YbscEncoder) includeYbsc(dest string) (string, []string) {
	return filepath.Join(dest, "lib/ybsc"), []string{"-bsc5", "data/bsc5.dat.gz"}
}

func (s *YbscEncoder) encode() error {
	var bsc catalogue.Catalog
	if err := io.NewReader().
		ForEachLine(func(l string) error {
			e, err := s.parseEntry(l)
			if err != nil {
				return err
			}

			if e.IsValid() {
				bsc.Append(e)
			}
			return nil
		}).
		DecompressIf(strings.HasSuffix(*s.Source, ".gz")).
		Open(*s.Source); err != nil {
		return err
	}

	dstFile := filepath.Join(*s.Encoder.Dest, "bsc5.bin")

	if err := io.NewWriter(bsc.Write).
		Compress().
		CreateFile(dstFile); err != nil {
		return err
	}
	log.Printf("Written %d stars", bsc.Size())

	// Verify we can read the catalog
	readBsc := &catalogue.Catalog{}
	if err := io.NewReader(readBsc.Read).
		Decompress().
		Open(dstFile); err != nil {
		return err
	}

	mc := 0
	for i := 0; i < bsc.Size(); i++ {
		a1 := bsc.Get(i)
		a2 := readBsc.Get(i)
		if a1.String() != a2.String() {
			log.Printf("%s -> %s", a1, a2)
			mc++
		}
	}
	log.Printf("%d errors", mc)

	if !bsc.Equals(readBsc) {
		return fmt.Errorf("YBSC catalog mismatch wrote %d read %d", bsc.Size(), readBsc.Size())
	}
	log.Printf("Read %d stars", readBsc.Size())

	return nil
}

func (s *YbscEncoder) parseEntry(l string) (catalogue.Entry, error) {
	ang, err := util.ParseAngle(l[75:77] + ":" + l[77:79] + ":" + l[79:83])
	if err != nil {
		return catalogue.Entry{}, err
	}
	ra := unit.RAFromDeg(ang.Deg() * 15.0)

	ang, err = util.ParseAngle(l[83:84] + l[84:86] + ":" + l[86:88] + ":" + l[88:90])
	if err != nil {
		return catalogue.Entry{}, err
	}
	dec := ang

	ms := strings.TrimSpace(l[102:107])
	mag := -99.0
	if ms != "" {
		mag, err = strconv.ParseFloat(strings.TrimSpace(l[102:107]), 64)
		if err != nil {
			return catalogue.Entry{}, err
		}
	}

	return catalogue.NewEntry(ra, dec, mag), nil
}
