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

func getBooks(c Client) {

	// CONSTRUCT URI
	URI, err := url.Parse(c.baseURL)
	if err != nil {
		fmt.Printf("Error parsing client url: %s \n", err.Error())
	}

	URI.Path += c.Path

	params := url.Values{}
	for i := 0; i < len(c.Params); i++ {
		params.Add(c.Params[i], c.ParamVals[i])
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

func booksClient(searchTerm string) Client {
	client = &http.Client{Timeout: 10 * time.Second}

	var bc Client
	bc.baseURL = "https://www.googleapis.com/books/v1"
	bc.Path = "/volumes"
	bc.Params = []string{"projection", "q", "key"}
	bc.ParamVals = []string{"lite", searchTerm, os.Getenv("GOOGLE_API_KEY")}

	return bc
}
