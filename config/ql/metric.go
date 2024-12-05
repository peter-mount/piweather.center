package ql

import "github.com/alecthomas/participle/v2/lexer"

// Metric handles a metric reference
type Metric struct {
	Pos lexer.Position

	Metric []string `parser:"@Ident ( '.' @Ident )*"`
	Name   string
}

func (v *visitor) Metric(b *Metric) error {
	if b != nil && v.metric != nil {
		return v.metric(v, b)
	}
	return nil
}

func (b *builder) Metric(f func(Visitor, *Metric) error) Builder {
	b.common.metric = f
	return b
}
