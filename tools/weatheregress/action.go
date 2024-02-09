package weatheregress

import (
	"github.com/peter-mount/go-script/executor"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheregress/lang"
)

type action struct {
	metric     api.Metric
	metrics    *lang.Metric
	routingKey string
	message    any
	exec       executor.Executor
}

func newAction(metric api.Metric, metrics *lang.Metric) *action {
	return &action{
		metric:     metric,
		metrics:    metrics,
		routingKey: "metric." + metric.Metric,
		message:    metric.String(),
		exec:       executor.NewExpressionExecutor(),
	}
}
