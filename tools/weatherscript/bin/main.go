package main

import (
	"fmt"
	"github.com/peter-mount/go-anim/tools/goanim"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-script/script"
	"github.com/peter-mount/piweather.center/tools/weatherscript"
	"os"
)

func main() {
	if err := kernel.Launch(
		&goanim.Anim{},
		&script.Script{},
		&weatherscript.Script{},
	); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
