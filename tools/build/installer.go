package build

import (
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
	"path/filepath"
)

type Installer struct {
	Encoder *core.Encoder `kernel:"inject"`
	Build   *core.Build   `kernel:"inject"`
}

func (s *Installer) Start() error {
	s.Build.AddExtension(s.extension)
	return nil
}

func (s *Installer) extension(arch arch.Arch, target target.Builder, meta *meta.Meta) {
	for _, srcDir := range []string{"demos" /*"include",*/} {
		destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), application.FileName(application.STATIC, filepath.Base(srcDir)))

		target.Target(destDir).
			MkDir(destDir).
			Echo("INSTALL", destDir).
			BuildTool("-copydir", srcDir, "-d", destDir)
	}
}
