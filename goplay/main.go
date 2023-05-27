package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "processed_ops_total",
		Help: "The total number of processed events",
	})
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(time.Second * 2)
		}
	}()
}

func main() {

	RecordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
