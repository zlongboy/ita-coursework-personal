package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client *http.Client

func getBooks(r BooksReq) {

	// CONSTRUCT URI
	URI, err := url.Parse(r.baseURL)
	if err != nil {
		fmt.Printf("Error parsing client url: %s \n", err.Error())
	}

	URI.Path += r.Path

	params := url.Values{}
	for i := 0; i < len(r.Params); i++ {
		params.Add(r.Params[i], r.ParamVals[i])
	}
	URI.RawQuery = params.Encode()
	// fmt.Printf("Print constructed URI: %s \n", URI)

	// BUILD REQUEST
	req, err := http.NewRequest("GET", URI.String(), nil)
	req.Header.Add("Content-Type", "application/json")

	// SEND REQUEST - HANDLE RESPONSE
	// var book Book

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s \n", err.Error())
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s \n", err.Error())
	}

	fmt.Println(resp.Status)
	fmt.Println(len(string(resBody)))

	// return json.NewDecoder(resp.Body).Decode(&book)
}

func makeRequest() {
	client = &http.Client{Timeout: 10 * time.Second}

	var br BooksReq
	br.baseURL = "https://www.googleapis.com/books/v1"
	br.Path = "/volumes"
	br.SearchTerm = "emily-st-john-mandel" // TODO: write function to get value from a param
	br.Params = []string{"projection", "q", "key"}
	br.ParamVals = []string{"lite", br.SearchTerm, os.Getenv("GOOGLE_API_KEY")}

	getBooks(br)
}
