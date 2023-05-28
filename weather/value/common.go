package value

import (
	"errors"
	"sync"
)

var (
	nilErr        = errors.New("nil Value")
	nan           = errors.New("NaN")
	pInf          = errors.New("+Inf")
	nInf          = errors.New("-Inf")
	mutex         sync.Mutex
	units         = make(map[string]*Unit)
	groups        = make(map[string]*Group)
	transformers  = make(map[string]Transformer)
	uncategorized = &Group{
		name: "Uncategorized",
		err:  errors.New("not uncategorized"),
	}
)

func init() {
	groups["uncategorized"] = uncategorized
}
