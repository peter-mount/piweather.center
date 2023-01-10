package util

import (
	"github.com/peter-mount/go-kernel/v2"
	"log"
)

func init() {
	kernel.RegisterAPI((*Debug)(nil), &debug{})
}

type Debug interface {
	IsVerbose() bool
	Println(...interface{})
	Printf(string, ...interface{})
}

type debug struct {
	Verbose *bool `kernel:"flag,v,Verbose"`
}

func (d *debug) IsVerbose() bool { return *d.Verbose }
func (d *debug) Println(v ...interface{}) {
	if *d.Verbose {
		log.Println(v...)
	}
}

func (d *debug) Printf(f string, v ...interface{}) {
	if *d.Verbose {
		log.Printf(f, v...)
	}
}
