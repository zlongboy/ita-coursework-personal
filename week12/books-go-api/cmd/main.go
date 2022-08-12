package main

// cli start server with nodemon:
// 	    nodemon --exec 'go' run cmd/main.go

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var list []Book

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtId(params["id"])

	json.NewEncoder(w).Encode(list)
}

func _deleteItemAtId(uid string) {
	for i, book := range list {
		if book.ID == uid {
			list = append(list[:i], list[i+1:]...)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	list = append(list, book)

	json.NewEncoder(w).Encode(book)
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(list)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/list", getList).Methods("GET")
	router.HandleFunc("/add", addBook).Methods("POST")
	router.HandleFunc("/list/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	list = append(list, Book{
		ID:    "123abc",
		Name:  "My Diary",
		Price: 9.99,
	})

	handleRequests()
}
