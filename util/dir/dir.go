package dir

import (
	"github.com/peter-mount/go-kernel/v2"
	"path"
)

func init() {
	// -d is optional if the tool supports it
	kernel.RegisterAPI((*Directory)(nil), &directory{})
}

// Directory service implements the -d command line flag common to all tools
type Directory interface {
	Resolve(fileName string) string
	Split(fileName string) (string, string)
}

// directory is the implementation
type directory struct {
	Dir *string `kernel:"flag,d,Directory to use,."`
}

func (d *directory) Start() error {
	// Ensure we default to current directory if someone used: -d ""
	if *d.Dir == "" {
		*d.Dir = "."
	}
	*d.Dir = path.Clean(*d.Dir)
	return nil
}

func (d *directory) Resolve(fileName string) string {
	dir, file := d.Split(fileName)
	return path.Join(dir, file)
}

func (d *directory) Split(fileName string) (string, string) {
	dir, file := path.Split(fileName)
	dir = path.Clean(dir)

	if dir == "." {
		dir = *d.Dir
	}

	return dir, file
}
