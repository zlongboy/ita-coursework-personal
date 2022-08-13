package main

// 	nodemon --exec 'go' run cmd/main.go

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Printf("Starting server on %v \n", port)
	log.Fatal(http.ListenAndServe(port, router()))
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", Health)
	r.Post("/admin/add", AddAuthor)

	return r
}
