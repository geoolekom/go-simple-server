package views

import (
	"net/http"
)

func NotFoundHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(404)
	w.Write([]byte("{}"))
}

func BadRequestHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(400)
	w.Write([]byte("{}"))
}

func EmptyBodyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}