package build

import (
	"compress/gzip"
	"fmt"
	"github.com/peter-mount/go-build/core"
	"io"
	"os"
	"path/filepath"
)

// Vsop87Encoder takes the compressed VSOP87 files and installs them uncompressed
// into the build.
//
// The data are the VSOP87B.* files (there's 8) from Vizier
// http://cdsarc.u-strasbg.fr/viz-bin/cat/VI/81#/browse
type Vsop87Encoder struct {
	Encoder *core.Encoder `kernel:"inject"`
	Build   *core.Build   `kernel:"inject"`
	Source  *string       `kernel:"flag,vsop87,install vsop87 data"`
}

func (s *Vsop87Encoder) Start() error {
	s.Build.AddLibProvider(s.includeVsop87)

	if *s.Source != "" {
		return s.encode()
	}
	return nil
}

func (s *Vsop87Encoder) includeVsop87(dest string) (string, []string) {
	return filepath.Join(dest, "lib/vsop87b"), []string{"-vsop87", "data"}
}

func (s *Vsop87Encoder) encode() error {
	if err := os.MkdirAll(*s.Source, 0755); err != nil {
		return err
	}

	for _, planet := range []string{"mer", "ven", "ear", "mar", "jup", "sat", "ura", "nep"} {
		if err := s.installPlanet(planet); err != nil {
			return err
		}
	}

	return nil
}

func (s *Vsop87Encoder) installPlanet(planet string) error {

	srcFile := filepath.Join(*s.Source, fmt.Sprintf("vsop87b.%s.gz", planet))
	dstFile := filepath.Join(*s.Encoder.Dest, fmt.Sprintf("VSOP87B.%s", planet))

	srcF, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer srcF.Close()

	gr, err := gzip.NewReader(srcF)
	if err != nil {
		return err
	}
	defer gr.Close()

	destF, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer destF.Close()

	_, err = io.Copy(destF, gr)

	return err
}
