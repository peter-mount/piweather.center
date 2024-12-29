package keogram

import "github.com/peter-mount/go-script/packages"

func init() {
	packages.RegisterPackage(&Package{})
}

type Package struct{}

func (*Package) Keogram() *Keogram {
	return New()
}
