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

func getBooks(rc RequestConfig) {

	// CONSTRUCT URI
	URI, err := url.Parse(rc.baseURL)
	if err != nil {
		fmt.Printf("Error parsing client url: %s \n", err.Error())
	}

	URI.Path += rc.Path

	params := url.Values{}
	for i := 0; i < len(rc.Params); i++ {
		params.Add(rc.Params[i], rc.ParamVals[i])
	}
	URI.RawQuery = params.Encode()
	// fmt.Printf("Print constructed URI: %s \n", URI)

	// BUILD REQUEST
	req, err := http.NewRequest("GET", URI.String(), nil)
	req.Header.Add("Content-Type", "application/json")

	// SEND REQUEST
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s \n", err.Error())
	}

	defer resp.Body.Close()

	// PARSE RESPONSE
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s \n", err.Error())
	}

	fmt.Println(resp.Status)
	fmt.Println(len(string(resBody)))

	// return json.NewDecoder(resp.Body).Decode(&book)

	// TODO: PARSE RESPONSE
}

func booksConfig(searchTerm string) RequestConfig {
	client = &http.Client{Timeout: 10 * time.Second}

	var rc RequestConfig
	rc.baseURL = "https://www.googleapis.com/books/v1"
	rc.Path = "/volumes"
	rc.Params = []string{"projection", "q", "key"}
	rc.ParamVals = []string{"lite", searchTerm, os.Getenv("GOOGLE_API_KEY")}

	return rc
}
