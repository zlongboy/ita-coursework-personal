package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/zlongboy/ita-coursework-personal/week12/books-api/util"
)

func main() {
	// port := fmt.Sprintf(":%v", getEnvMap()["PORT"])
	port := ":8000"
	fmt.Printf("Starting server on %v \n", port)
	log.Fatal(http.ListenAndServe(port, router()))
}

func router() http.Handler {
	r := chi.NewRouter()

	return r
}
