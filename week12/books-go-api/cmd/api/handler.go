package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {

	// Pass in to booksClient(searchTerm)
	// Pass into getBooks()
	// ResponseParsed =
	// Make 3 channels: 1 to get and run authors info queries/publishers/books
	// Join up --> send response (query all records with that author id, return as JSON object)
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// OpenDB()

	ap := r.URL.Query().Get("author")
	key := r.URL.Query().Get("key")

	if key != os.Getenv("TEST_API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		w.Write(BadRequest(RespInvalidAuth))
		return
	} else if len(ap) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(BadRequest(RespMissingAuthor))
		return
	}

	w.WriteHeader(http.StatusOK)
	// successJSON, err := json.MarshalIndent(getBooks(booksConfig(ap)), "", " ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// w.Write(getBooks(booksConfig(ap)))
	// w.Write(successJSON)

	items := getBooks(booksConfig(ap)).Items

	for i, b := range items {
		if len(b.VolumeInfo.Authors) > 0 {
			fmt.Printf("%v) %v \n", i, b.VolumeInfo.Authors[0])
		}
	}

}

func BadRequest(msg string) []byte {
	resp := make(map[string]string)

	resp["message"] = msg
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling %v to JSON", resp)
	}

	return respJSON
}
