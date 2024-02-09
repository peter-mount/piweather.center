package weatheregress

import (
	"github.com/peter-mount/go-script/calculator"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type action struct {
	metric  api.Metric
	metrics []*lang.Metric
	message any
	calc    calculator.Calculator
}
