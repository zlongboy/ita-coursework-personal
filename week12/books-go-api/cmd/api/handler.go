package main

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	// TODO: get author param of incoming request
	// Pass in to booksClient(searchTerm)
	// Pass into getBooks()
	// ResponseParsed =
	// Make 3 channels: 1 to get and run authors info queries/publishers/books
	// Join up --> send response (query all records with that author id, return as JSON object)
}

func Test(w http.ResponseWriter, r *http.Request) {
	OpenDB()
}
