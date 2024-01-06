package functions

import (
	"fmt"
	_ "github.com/peter-mount/piweather.center/astro/calculator"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strings"
	"sync"
)

type Function struct {
	Name       string
	Calculator value.Calculator
	Op         func(value.Value) (value.Value, error)
	BiOp       func(value.Value, value.Value) (value.Value, error)
	TriOp      func(value.Value, value.Value, value.Value) (value.Value, error)
	MathOp     func(float64) float64
	MathBiOp   func(float64, float64) float64
}

func (f *Function) IsEmpty() bool {
	return f.Calculator == nil &&
		f.Op == nil && f.BiOp == nil && f.TriOp == nil &&
		f.MathOp == nil && f.MathBiOp == nil
}

var (
	functions     = make(map[string]*Function)
	functionMutex sync.Mutex
)

func init() {
	if err := AddFunctions(
		Function{Name: "max", MathBiOp: math.Max},
		Function{Name: "min", MathBiOp: math.Min},
		Function{Name: "floor", MathOp: math.Floor},
		Function{Name: "ceil", MathOp: math.Ceil},
		Function{Name: "AbsoluteHumidity", BiOp: measurement.GetAbsoluteHumidity},
		Function{Name: "DewPoint", BiOp: measurement.GetDewPoint},
		Function{Name: "FeelsLike", TriOp: measurement.FeelsLike},
		Function{Name: "HeatIndex", BiOp: measurement.HeatIndex},
		Function{Name: "SolarAltitude"},
		Function{Name: "SolarAzimuth"},
		Function{Name: "WindChill", BiOp: measurement.WindChill},
	); err != nil {
		panic(err)
	}
}

func AddFunctions(fs ...Function) error {
	for _, f := range fs {
		if err := AddFunction(f); err != nil {
			return err
		}
	}
	return nil
}

func AddFunction(f Function) error {
	f.Name = strings.ToLower(f.Name)

	if f.IsEmpty() {
		c, err := value.GetCalculator(f.Name)
		if err == nil {
			f.Calculator = c
		}
	}

	functionMutex.Lock()
	defer functionMutex.Unlock()

	if _, exists := functions[f.Name]; exists {
		return fmt.Errorf("function %q already defined", f.Name)
	}
	functions[f.Name] = &f
	return nil
}

func LookupFunction(n string) (*Function, bool) {
	f, e := lookupFunction(n)
	if !e {
		c, err := value.GetCalculator(n)
		if err == nil {
			_ = AddFunction(Function{Name: n, Calculator: c})
			return lookupFunction(n)
		}
	}
	return f, e
}

func lookupFunction(n string) (*Function, bool) {
	n = strings.ToLower(n)

	functionMutex.Lock()
	defer functionMutex.Unlock()

	f, e := functions[n]

	return f, e
}
