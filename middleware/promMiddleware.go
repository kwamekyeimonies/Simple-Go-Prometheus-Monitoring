package middleware

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/controller"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/service"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathRegexp()

		timer := prometheus.NewTimer(service.HttpDuration.WithLabelValues(path))
		rw := controller.NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.StatusCode

		service.ResponseStatus.WithLabelValues(strconv.Itoa(int(statusCode))).Inc()
		service.TotalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()

	})
}
