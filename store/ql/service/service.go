package service

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
	"github.com/peter-mount/piweather.center/store/ql/exec"
	"net/http"
)

func init() {
	kernel.RegisterAPI((*Service)(nil), &service{})
}

// Service provides a Kernel interface to the query language.
// This supports throttling of queries so that the system is not overloaded by too many concurrent
// requests being made.
type Service interface {
	Query(fileName string, query []byte, opts ...exec.QueryOption) *api.Result
}

type service struct {
	// PoolSize is the number of concurrent requests that can be performed at once.
	PoolSize *int       `kernel:"flag,query-pool,Query Pool Size,5"`
	Store    file.Store `kernel:"inject"`
	ch       chan *Request
}

func (s *service) Query(fileName string, query []byte, opts ...exec.QueryOption) *api.Result {
	r := &Request{
		FileName: fileName,
		Query:    query,
		Opts:     opts,
		ch:       make(chan *api.Result),
	}

	s.ch <- r
	return <-r.ch
}

func (s *service) Start() error {
	if *s.PoolSize < 1 {
		*s.PoolSize = 1
	}

	s.ch = make(chan *Request, *s.PoolSize)

	for i := 0; i < *s.PoolSize; i++ {
		go s.run()
	}

	return nil
}

func (s *service) run() {
	for {
		req := <-s.ch

		resp := s.query(req)
		go func() {
			req.ch <- resp
		}()
	}
}

func (s *service) query(req *Request) (result *api.Result) {
	defer func() {
		if err1 := recover(); err1 != nil {
			result = &api.Result{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("%v", err1),
			}
		}
	}()

	result, _ = exec.Query(s.Store, req.FileName, req.Query, req.Opts...)
	return
}
