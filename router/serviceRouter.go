package router

import (
	"github.com/gorilla/mux"
	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/controller"
)

func ServiceRouter(router *mux.Router) {
	router.HandleFunc("/service", controller.Service).Methods("GET")
}
