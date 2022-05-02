package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Session = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "session",
}, []string{"status", "msg"})

var Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "duration_of_request",
}, []string{"method"})
