package exec

import (
	"fmt"
	"github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/file"
	"net/http"
	"strings"
)

type QueryOption int

const (
	OptQuery QueryOption = iota
	OptQueryPlan
)

func (o QueryOption) Present(opts []QueryOption) bool {
	for _, e := range opts {
		if o == e {
			return true
		}
	}
	return false
}

func (o QueryOption) AppendIf(opts []QueryOption, f bool) []QueryOption {
	if f {
		return append(opts, o)
	}
	return opts
}

func Query(s file.Store, fileName string, query []byte, opts ...QueryOption) (*api.Result, error) {
	result := &api.Result{Status: http.StatusOK}

	err := queryImpl(s, fileName, query, result, opts)
	if err != nil {
		result.Status = http.StatusBadRequest
		result.Message = err.Error()
		result.Table = nil

		// If we caught a panic in queryImpl then set status to 500
		if strings.HasPrefix(result.Message, "panic:") {
			result.Status = http.StatusInternalServerError
		}
	}
	return result, err
}

func queryImpl(s file.Store, fileName string, query []byte, result *api.Result, opts []QueryOption) (err error) {
	// Capture any panic's
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("panic: %v", err1)
		}
	}()

	q, err := ql.NewParser().ParseBytes(fileName, query)
	if err != nil {
		return err
	}

	if OptQuery.Present(opts) {
		// TODO reimplement
		//result.AddMeta("query", q.String())
	}

	qp, err := NewQueryPlan(s, q)
	if err != nil {
		return err
	}

	if OptQueryPlan.Present(opts) {
		result.AddMeta("queryPlan", qp)
	}

	// Note: Copy as result needs a pointer
	r := qp.QueryRange
	result.Range = &r
	err = qp.Execute(result)

	return err
}
