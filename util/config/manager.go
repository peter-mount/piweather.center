package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/util/io"
	"os"
	"path"
	"path/filepath"
	"sync"
)

func init() {
	kernel.RegisterAPI((*Manager)(nil), &manager{})
}

var (
	pathSep = string(os.PathSeparator)
)

type Manager interface {
	// EtcDir returns the path to the etc directory
	EtcDir() string
	// ReadJson will read a yaml file from inside etc
	ReadJson(n string, o any) error
	// ReadJsonOptional is the same as ReadJson except that the file not existing will not
	// return an error
	ReadJsonOptional(n string, o any) error
	// ReadYaml will read a yaml file from inside etc
	ReadYaml(n string, o any) error
	// ReadYamlOptional is the same as ReadYaml except that the file not existing will not
	// return an error
	ReadYamlOptional(n string, o any) error
	// FixPath ensures that, if the path is relative it will point to EtcDir().
	// This will do nothing is s==nil or *s==""
	FixPath(s *string)
	// ReadAndWatch read with a custom Unmarshaller but also if the file ever changes.
	ReadAndWatch(n string, f Factory, u Updater, um Unmarshaller) error
	// WatchDirectory watches a directory and passes all changes to a single Updater
	WatchDirectory(d string, f Factory, u Updater, um Unmarshaller)
	WatchDirectoryParser(d string, f Factory, u Updater, um Parser)
}

type manager struct {
	RootDirFlag *string `kernel:"flag,rootDir,Location of config files"`
	mutex       sync.Mutex
	watcher     *fsnotify.Watcher
	watching    map[string][]*updater
	loader      chan fsnotify.Event
}

func (m *manager) Start() error {
	// Path to script directory for data lookup
	if *m.RootDirFlag == "" {
		*m.RootDirFlag = path.Join(filepath.Dir(os.Args[0]), "../etc")
	}

	m.watching = make(map[string][]*updater)
	w, err := fsnotify.NewWatcher()
	if err == nil {
		m.watcher = w
		err = w.Add(m.EtcDir())
	}

	m.loader = make(chan fsnotify.Event, 100)
	// Ok then start updater
	if err == nil {
		go m.run()
	}

	return err
}

func (m *manager) EtcDir() string {
	return *m.RootDirFlag
}

func (m *manager) ReadJson(n string, o any) error {
	return io.NewReader().
		Json(o).
		Open(filepath.Join(m.EtcDir(), n))
}

func (m *manager) ReadJsonOptional(n string, o any) error {
	err := m.ReadJson(n, o)
	// Ignore the file not existing
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
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

func (m *manager) FixPath(s *string) {
	if s != nil && *s != "" && !filepath.IsAbs(*s) {
		*s = filepath.Join(m.EtcDir(), *s)
	}
}
