package main

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {

}
