package value

import (
	"context"
	"strings"
	"sync"
)

type Map struct {
	mutex sync.Mutex
	m     map[string]Value
}

func WithMap(ctx context.Context) context.Context {
	return context.WithValue(ctx, "valueMap", &Map{
		m: make(map[string]Value),
	})
}

func MapFromContext(ctx context.Context) *Map {
	return ctx.Value("valueMap").(*Map)
}

func (m *Map) Get(k string) Value {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.get(k)
}

func (m *Map) get(k string) Value {
	return m.m[strings.ToLower(k)]
}

func (m *Map) GetAll(keys ...string) []Value {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var args []Value
	for _, key := range keys {
		args = append(args, m.get(key))
	}
	return args
}
func (m *Map) Put(k string, v Value) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.m[strings.ToLower(k)] = v
}

func (m *Map) Reset() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m != nil {
		m.m = make(map[string]Value)
	}
}

func ResetMap(ctx context.Context) error {
	m := MapFromContext(ctx)
	if m != nil {
		m.Reset()
	}
	return nil
}

func (m *Map) GetKeys() []string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var a []string
	for k, _ := range m.m {
		a = append(a, k)
	}
	return a
}
