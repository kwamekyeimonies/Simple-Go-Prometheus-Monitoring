package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var TotalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name:"http_requests_total",
		Help:"Number of get request",
	},
	[]string{"path"},
)