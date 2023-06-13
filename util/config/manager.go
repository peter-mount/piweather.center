package config

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/io"
	"os"
	"path"
	"path/filepath"
)

func init() {
	kernel.RegisterAPI((*Manager)(nil), &manager{})
}

type Manager interface {
	RootDir() string
	ReadYaml(n string, o any) error
}

type manager struct {
	RootDirFlag *string `kernel:"flag,rootDir,Location of config files"`
}

func (m *manager) Start() error {
	// Path to lib directory for data lookup
	if *m.RootDirFlag == "" {
		*m.RootDirFlag = path.Join(filepath.Dir(os.Args[0]), "../etc")
	}
	return nil
}

func (m *manager) RootDir() string {
	return *m.RootDirFlag
}

func (m *manager) ReadYaml(n string, o any) error {
	return io.NewReader().
		Yaml(o).
		Open(filepath.Join(m.RootDir(), n))
}
