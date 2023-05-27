package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/middleware"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/router"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/service"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.Register(service.TotalRequests)
	prometheus.Register(service.ResponseStatus)
	prometheus.Register(service.HttpDuration)
}

func main() {

	server := mux.NewRouter()

	server.Use(middleware.PrometheusMiddleware)

	fmt.Println("Serving requests on port 9090")

	router.ServiceRouter(server)
	router.PrometheusRouter(server)

	http.ListenAndServe(":9090", server)
}
