package cloud

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/util/cloud"
	"image"
)

func init() {
	packages.RegisterPackage(&Cloud{})
}

type Cloud struct{}

func (_ *Cloud) Filter(mask image.Image) cloud.Filter {
	return cloud.NewFilter(mask)
}

func (_ *Cloud) FilterNoMask() cloud.Filter {
	return cloud.NewFilter(nil)
}
