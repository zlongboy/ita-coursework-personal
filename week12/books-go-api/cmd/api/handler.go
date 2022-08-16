package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SERVER IS RUNNING"))
}

func AddBooks(w http.ResponseWriter, r *http.Request) {

	// Make 3 channels: 1 to get and run authors info queries/publishers/books
	// Join up --> send response successMsg()
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	toSave := getBooks(booksConfig(ap)) // Google books request

	// OpenDB(toSave) // Run database queries

	getAuthors(toSave)
	getPublishers(toSave)

	w.Write(ResponseMsg(RespSuccess)) // TODO: Replace with SuccessMsg when values are available
}

// DECONSTRUCTING BOOKVALUES
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

	// var ajs []JoinBook
	// for _, b := range bv {
	// 	var jb JoinBook
	// 	jb.BookID = b.ID

	// }

	// fmt.Println(authorIds)
	return authorIds
}

func getPublishers(bv []BookValues) map[string]string { // TODO: Refactor, combine with getAuthors
	var as []string

	for _, b := range bv {
		as = append(as, b.Publisher)
	}

	publisherIds := make(map[string]string)
	for _, ua := range unique(as) {
		aid := uuid.New().String()
		publisherIds[ua] = aid
	}

	// fmt.Println(publisherIds)
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

func SuccessMsg(msg string, b, a, p int) []byte {
	resp := make(map[string]string)

	resp["message"] = msg
	resp["Books added"] = strconv.Itoa(b)
	resp["Authors added"] = strconv.Itoa(a)
	resp["Publishers added"] = strconv.Itoa(p)
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
