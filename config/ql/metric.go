package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

// Metric handles a metric reference
type Metric struct {
	Pos lexer.Position

	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

func (v *visitor[T]) Metric(b *Metric) error {
	if b != nil && v.metric != nil {
		return v.metric(v, b)
	}
	return nil
}

func initMetric(_ Visitor[*parserState], b *Metric) error {
	b.Name = strings.Join(b.Metric, ".")
	return nil
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.common.metric = f
	return b
}
