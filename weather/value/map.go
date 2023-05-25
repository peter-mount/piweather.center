package value

import (
	"context"
	"strings"
)

type Map struct {
	m map[string]Value
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
	return m.m[strings.ToLower(k)]
}

func (m *Map) GetAll(keys ...string) []Value {
	var args []Value
	for _, key := range keys {
		args = append(args, m.Get(key))
	}
	return args
}
func (m *Map) Put(k string, v Value) {
	m.m[strings.ToLower(k)] = v
}

func (m *Map) Reset() {
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
