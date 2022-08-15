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

	// GET QUERY PARAMS
	ap := r.URL.Query().Get("author")
	key := r.URL.Query().Get("key")

	// HANDLE BAD REQUEST
	if key != os.Getenv("TEST_API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		w.Write(ResponseMsg(RespInvalidAuth))
		return
	} else if len(ap) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ResponseMsg(RespMissingAuthor))
		return
	}

	// HANDLE SUCCESSFUL REQUEST
	w.WriteHeader(http.StatusOK)

	toSave := getBooks(booksConfig(ap))

	// TO DELETE - TEST BLOCK //
	toPrint, err := json.MarshalIndent(toSave, "", " ")
	if err != nil {
		fmt.Printf("Error marshalling: %s \n", err.Error())
	}
	fmt.Printf("%s \n", toPrint)
	// END TEST BLOCK //

	w.Write(ResponseMsg(RespSuccess))
}

func ResponseMsg(msg string) []byte {
	resp := make(map[string]string)

	resp["message"] = msg
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling %v to JSON", resp)
	}

	return respJSON
}
