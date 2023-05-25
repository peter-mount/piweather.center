package value

import (
	"errors"
	"sync"
)

var (
	nilErr       = errors.New("nil Value")
	nan          = errors.New("NaN")
	pInf         = errors.New("+Inf")
	nInf         = errors.New("-Inf")
	mutex        sync.Mutex
	units        = make(map[string]*Unit)
	transformers = make(map[string]Transformer)
)
