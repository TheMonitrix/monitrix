package telemetrytypes

import "github.com/ezeslucky/monitrix/pkg/valuer"

type Signal struct {
	valuer.String
}

var (
	SignalTraces      = Signal{valuer.NewString("traces")}
	SignalLogs        = Signal{valuer.NewString("logs")}
	SignalMetrics     = Signal{valuer.NewString("metrics")}
	SignalUnspecified = Signal{valuer.NewString("")}
)
