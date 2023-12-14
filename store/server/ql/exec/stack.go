package exec

func (ex *Executor) resetStack() {
	ex.stack = nil
}

func (ex *Executor) stackEmpty() bool {
	return len(ex.stack) == 0
}

func (ex *Executor) push(v Value) {
	ex.stack = append(ex.stack, v)
}

func (ex *Executor) pop() (Value, bool) {
	if ex.stackEmpty() {
		return Value{}, false
	}
	sl := len(ex.stack) - 1
	r := ex.stack[sl]
	ex.stack = ex.stack[:sl]
	return r, true
}
