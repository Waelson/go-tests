package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricsRecord interface {
	IncrementTotalRequest()
}

type metricsRecord struct {
	counter prometheus.Counter
}

func (m *metricsRecord) IncrementTotalRequest() {
	m.counter.Inc()
}

func NewMetricsRecord() MetricsRecord {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "The total number of requests",
	})

	return &metricsRecord{
		counter: counter,
	}
}
