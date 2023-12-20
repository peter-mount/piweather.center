package exec

import "github.com/peter-mount/piweather.center/store/ql"

func (ex *Executor) resetStack() {
	ex.stack = nil
}

func (ex *Executor) stackEmpty() bool {
	return len(ex.stack) == 0
}

func (ex *Executor) Push(v ql.Value) {
	ex.stack = append(ex.stack, v)
}

func (ex *Executor) Pop() (ql.Value, bool) {
	if ex.stackEmpty() {
		return ql.Value{}, false
	}
	sl := len(ex.stack) - 1
	r := ex.stack[sl]
	ex.stack = ex.stack[:sl]
	return r, true
}
