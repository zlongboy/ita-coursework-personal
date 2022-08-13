package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Handler(r chi.Router) {
	r.Get("/admin/health", Health)
	r.Post("/admin/add", AddAuthor)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {

}
