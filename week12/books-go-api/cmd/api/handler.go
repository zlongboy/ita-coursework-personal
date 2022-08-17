package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New health check request")
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	// TODO: Read array of authors and make chans to handle process for each one

	fmt.Println("New add books request")
	w.Header().Set("Content-Type", "application/json")

	// GET QUERY PARAMS
	author := r.URL.Query().Get("author")
	key := r.URL.Query().Get("key")

	// HANDLE BAD REQUEST
	if key != os.Getenv("TEST_API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		w.Write(ResponseMsg(RespInvalidAuth))
		return
	} else if len(author) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ResponseMsg(RespMissingAuthor))
		return
	}

	// HANDLE SUCCESSFUL REQUEST
	w.WriteHeader(http.StatusOK)

	toSave := getBooks(booksConfig(author)) // Google books request

	qResults := OpenDB(toSave, getAuthors(toSave), getPublishers(toSave)) // Run database queries

	w.Write(SuccessMsg(RespSuccess, qResults))

}

// HELPER FUNCTIONS: Deconstruct book values
func getAuthors(bv []BookValues) map[string]string {
	var as []string

	for _, b := range bv {
		as = append(as, b.Author)
	}

	authorIds := make(map[string]string)
	for _, ua := range unique(as) {
		aid := uuid.New().String()
		authorIds[ua] = aid
	}

	return authorIds
}

func getPublishers(bv []BookValues) map[string]string { // TODO: Refactor, combine with getAuthors(), one function returns two maps
	var as []string

	for _, b := range bv {
		as = append(as, b.Publisher)
	}

	publisherIds := make(map[string]string)
	for _, ua := range unique(as) {
		aid := uuid.New().String()
		publisherIds[ua] = aid
	}

	return publisherIds

}

// RESPONSE MESSAGE BUILDERS
func ResponseMsg(msg string) []byte {
	resp := make(map[string]string)

	resp["message"] = msg
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling %v to JSON", resp)
	}

	return respJSON
}

func SuccessMsg(msg string, r Success) []byte {
	resp := make(map[string]string)

	resp["Message"] = msg
	resp["Books added"] = strconv.Itoa(int(r.Books))
	resp["Authors added"] = strconv.Itoa(int(r.Authors))
	resp["Publishers added"] = strconv.Itoa(int(r.Publishers))
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling %v to JSON", resp)
	}

	return respJSON

}

// UTILS
func unique(ss []string) []string {
	m := make(map[string]bool)
	var u []string

	for _, v := range ss {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}

	return u
}
