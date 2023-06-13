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
	// EtcDir returns the path to the etc directory
	EtcDir() string
	// ReadYaml will read a yaml file from inside etc
	ReadYaml(n string, o any) error
	// ReadYamlOptional is the same as ReadYaml except that the file not existing will not
	// return an error
	ReadYamlOptional(n string, o any) error
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

func (m *manager) EtcDir() string {
	return *m.RootDirFlag
}

func (m *manager) ReadYaml(n string, o any) error {
	return io.NewReader().
		Yaml(o).
		Open(filepath.Join(m.EtcDir(), n))
}

func (m *manager) ReadYamlOptional(n string, o any) error {
	err := m.ReadYaml(n, o)
	// Ignore the file not existing
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
}
