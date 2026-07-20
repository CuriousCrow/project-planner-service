package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_http_requests_total",
		Help: "Общее количество HTTP-запросов",
	})
)

func init() {
	prometheus.MustRegister(RequestCounter)
}
