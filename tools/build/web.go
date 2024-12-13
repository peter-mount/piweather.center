package build

import (
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/core"
	"github.com/peter-mount/go-build/util/arch"
	"github.com/peter-mount/go-build/util/makefile/target"
	"github.com/peter-mount/go-build/util/meta"
	"path/filepath"
)

// WebEncoder simply installs the web content into the distribution
type WebEncoder struct {
	Encoder *core.Encoder `kernel:"inject"`
	Build   *core.Build   `kernel:"inject"`
}

func (s *WebEncoder) Start() error {
	s.Build.AddExtension(s.extension)
	return nil
}

func (s *WebEncoder) extension(arch arch.Arch, target target.Builder, meta *meta.Meta) {
	//destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), "web")
	destDir := filepath.Join(arch.BaseDir(*s.Encoder.Dest), application.FileName(application.STATIC, "web"))
	target.
		Target(destDir).
		MkDir(destDir).
		Echo("WEB", destDir).
		BuildTool("-copydir", "web", "-d", destDir)
}
