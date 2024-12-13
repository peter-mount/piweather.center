package build

import (
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
	"path/filepath"
)

type ConfigInstaller struct {
	Encoder *core.Encoder `kernel:"inject"`
	Build   *core.Build   `kernel:"inject"`
}

func (s *ConfigInstaller) Start() error {
	s.Build.AddExtension(s.extension)
	return nil
}

func (s *ConfigInstaller) extension(arch arch.Arch, target target.Builder, meta *meta.Meta) {
	//destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), "etc")
	destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), application.FileName(application.CONFIG))

	target.
		Target(destDir).
		MkDir(destDir).
		Echo("CONFIG", destDir).
		BuildTool("-copydir", "etc", "-d", destDir)
}
