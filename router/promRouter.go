package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusRouter(router *mux.Router) {
	router.HandleFunc("/metrics", promhttp.Handler().ServeHTTP).Methods("GET")
}
