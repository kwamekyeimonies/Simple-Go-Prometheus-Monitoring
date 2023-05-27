package controller

import (
	"net/http"

	"github.com/kwamekyeimonies/Simple-Go-Prometheus-Monitoring/model"
)

func NewResponseWriter(w http.ResponseWriter) *model.ResponseWriter {
	return &model.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}
