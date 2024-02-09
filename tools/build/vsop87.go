package build

import (
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
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
}

func (s *Vsop87Encoder) Start() error {
	s.Build.AddExtension(s.extension)
	s.Build.Makefile(0, s.makefile)
	return nil
}

func (s *Vsop87Encoder) extension(arch arch.Arch, target target.Builder, meta *meta.Meta) {
	destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), "lib/vsop87b")

	destDirTarget := target.Target(destDir)

	for _, planet := range []string{"mer", "ven", "ear", "mar", "jup", "sat", "ura", "nep"} {
		src := filepath.Join("data", "vsop87b."+planet+".gz")
		dest := filepath.Join(destDir, "VSOP87B."+planet)
		destDirTarget.
			Target(dest, src).
			MkDir(destDir).
			Echo("VSOP87", dest).
			BuildTool("-gunzip", src, "-d", dest)
	}
}

func (s *Vsop87Encoder) makefile(root makefile.Builder, _ target.Builder, _ *meta.Meta) {
	// This adds a dependency to test for vsop87.
	//
	// With gnu-make, this works because it merges this test rule with the normal one
	// so when test is run, vsop87 is performed first so that the tests have access
	// to that data
	root.Rule("test",
		filepath.Join(s.Build.BuildArch().BaseDir(*s.Encoder.Dest), "lib", "vsop87b"),
	)
}
