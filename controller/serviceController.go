package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/model"
)

func Service(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("An Error occured %v", resp.Status)
	}

	var response model.Response
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(response)
}
