package util

import (
	"net/http"
	"strconv"
)

type MetricMiddleware struct {
	metricsRecord MetricsRecord
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (m *MetricMiddleware) Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{ResponseWriter: w}
		next.ServeHTTP(&rec, r)
		m.metricsRecord.IncrementTotalRequest(r.Method, r.URL.Path, strconv.Itoa(rec.status))
	}
	return http.HandlerFunc(fn)
}

func NewMetricMiddleware(metricsRecord MetricsRecord) *MetricMiddleware {
	return &MetricMiddleware{
		metricsRecord: metricsRecord,
	}
}
