package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"reflect"
)

var (
	pseudoVisitor = NewBuilder[reflect.Value]().
		Calculation(pseudoSetPositionImpl[Calculation]).
		Build()
)

func pseudoSetPosition(pos lexer.Position) Visitor[reflect.Value] {
	return pseudoVisitor.Clone().Set(reflect.ValueOf(pos))
}

func pseudoSetPositionImpl[T any](v Visitor[reflect.Value], d *T) error {
	if d != nil {
		tv := reflect.ValueOf(d)
		t := tv.Type()
		if t == nil {
			return nil
		}

		val := tv.Elem().FieldByName("Pos")
		if val.IsValid() && val.CanSet() {
			val.Set(v.Get())
		}
	}
	return nil
}
