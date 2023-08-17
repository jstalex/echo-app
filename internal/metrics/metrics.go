package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var NameMetric = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "names",
	Name:      "vasyas",
	Help:      "Vasya's users count",
})

func AddVasya() {
	NameMetric.Inc()
}
