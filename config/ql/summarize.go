package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
)

type Summarize struct {
	Pos  lexer.Position
	With string `parser:"('summarised'|'summarized') ('with' @String)"`
}

func (v *visitor[T]) Summarize(b *Summarize) error {
	var err error
	if b != nil {
		if v.summarize != nil {
			err = v.summarize(v, b)
			if errors.IsVisitorStop(err) || errors.IsVisitorExit(err) {
				return nil
			}
		}
	}
	return err
}

func (b *builder[T]) Summarize(f func(Visitor[T], *Summarize) error) Builder[T] {
	b.common.summarize = f
	return b
}
