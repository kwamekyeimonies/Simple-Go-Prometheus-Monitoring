package middleware

// import (
// 	"net/http"

// 	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/services"
// )

// func PrometheusMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		rw := w
// 		next.ServeHTTP(rw, r)

// 		services.TotalRequest.WithValues(path).Inc()
// 	})
// }
