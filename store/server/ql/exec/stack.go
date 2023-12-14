package exec

import "github.com/peter-mount/go-kernel/v2/log"

func (ex *executor) resetStack() {
	ex.stack = nil
}

func (ex *executor) stackEmpty() bool {
	return len(ex.stack) == 0
}

func (ex *executor) push(v Value) {
	ex.stack = append(ex.stack, v)
}

func (ex *executor) pop() (Value, bool) {
	if ex.stackEmpty() {
		log.Printf("StackEmpty")
		return Value{IsNull: true}, false
	}
	sl := len(ex.stack) - 1
	r := ex.stack[sl]
	ex.stack = ex.stack[:sl]
	log.Printf("Pop %d -> %d %v", sl, len(ex.stack), r)
	return r, true
}
