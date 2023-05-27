package model

import "net/http"

type Response struct {
	Message []string `json:"message"`
}

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode uint
}
