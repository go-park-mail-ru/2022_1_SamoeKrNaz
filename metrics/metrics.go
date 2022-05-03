package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Session = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "session",
}, []string{"status", "msg"})

var DurationSession = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "duration_of_request_session",
}, []string{"method"})

var User = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "user",
}, []string{"status", "msg"})

var DurationUser = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "duration_of_request_user",
}, []string{"method"})
