package main

import (
	"fmt"
	_ "github.com/peter-mount/go-anim/script"
	"github.com/peter-mount/go-anim/tools/goanim"
	"github.com/peter-mount/go-kernel/v2"
	_ "github.com/peter-mount/go-script/stdlib"
	_ "github.com/peter-mount/go-script/stdlib/fmt"
	_ "github.com/peter-mount/go-script/stdlib/io"
	_ "github.com/peter-mount/go-script/stdlib/math"
	_ "github.com/peter-mount/go-script/stdlib/time"
	"os"
)

func main() {
	if err := kernel.Launch(&goanim.Anim{}); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
