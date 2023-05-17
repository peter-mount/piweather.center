package menu

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/util/template"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func init() {
	kernel.Register(&Service{})
}

type Menu struct {
	Section  string `yaml:"section"`
	Label    string `yaml:"label"`
	Link     string `yaml:"link"`
	Entries  []Menu `yaml:"entries"`
	Disabled bool   `yaml:"disabled"`
}

type Service struct {
	Templates *template.Manager `kernel:"inject"`
	menu      *Menu
}

func (s *Service) PostInit() error {
	s.Templates.AddFunction("getMenu", s.getMenu)
	return nil
}

func (s *Service) Start() error {
	s.menu = &Menu{}

	rootDir := filepath.Dir(s.Templates.GetRootDir())
	menuFile := filepath.Join(rootDir, "menu.yaml")

	log.Printf("Loading Menu %s", menuFile)
	b, err := os.ReadFile(menuFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, s.menu)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) getMenu() interface{} {
	return s.menu.Entries
}
