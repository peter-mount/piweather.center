package main

import (
	"fmt"
	"github.com/peter-mount/go-anim/tools/build/font"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/build"
	"os"
)

func main() {
	if err := kernel.Launch(
		&build.ConfigInstaller{},
		&font.FontDownloader{},
		&build.Installer{},
		&build.FeatureSet{},
		&build.Vsop87Encoder{},
		&build.YbscEncoder{},
		&build.WebEncoder{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
