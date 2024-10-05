package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type MetricsRecord interface {
	IncrementTotalRequest(tags ...string)
}

type metricsRecord struct {
	counter *prometheus.CounterVec
}

func (m *metricsRecord) IncrementTotalRequest(tags ...string) {
	m.counter.WithLabelValues(tags...).Inc()
}

func NewMetricsRecord() MetricsRecord {
	counter := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "The total number of http_requests",
	}, []string{"method", "path", "status_code"})

	return &metricsRecord{
		counter: counter,
	}
}
