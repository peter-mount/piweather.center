package service

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/ql/exec"
)

type Request struct {
	FileName string
	Query    []byte
	Opts     []exec.QueryOption
	ch       chan *api.Result
}
