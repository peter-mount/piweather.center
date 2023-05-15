package value

import (
	"errors"
	"sync"
)

var (
	nan          = errors.New("NaN")
	pInf         = errors.New("+Inf")
	nInf         = errors.New("-Inf")
	mutex        sync.Mutex
	units        = make(map[string]Unit)
	transformers = make(map[string]Transformer)
)
