package dataencoder

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"os"
)

type Encoder struct {
	Dest *string `kernel:"flag,d,Destination directory"`
}

func (e *Encoder) Start() error {
	if *e.Dest != "" {
		fi, err := os.Stat(*e.Dest)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Creating %q", *e.Dest)
				return os.MkdirAll(*e.Dest, 0755)
			}
			return err
		}

		if !fi.IsDir() {
			return fmt.Errorf("%q is not a directory", *e.Dest)
		}
	}
	return nil
}
