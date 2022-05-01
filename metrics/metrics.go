package metrics

import "github.com/prometheus/client_golang/prometheus"

var Session = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "session",
}, []string{"msg"})
