package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()


	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	router.Path("/prometheus").Handler(promhttp.Handler())

	err := http.ListenAndServe(":9090", router)
	log.Fatal(err)
}
