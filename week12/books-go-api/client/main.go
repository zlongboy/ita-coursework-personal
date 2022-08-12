package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/joho/godotenv"
)

// TODO: Remove duplicate getEnvMap function and import util
func getEnvMap() map[string]string {

	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error reading .env to map")
	}
	return envMap
}

var client *http.Client

// Request struct
type BooksReq struct {
	baseURL    string
	Path       string
	APIKey     string
	SearchTerm string
	Params     []string
	ParamVals  []string
}

// Response structs
type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	Desc        string  `json:"desc"`
	PublishDate string  `json:"publish_date"`
	Country     string  `json:"country"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"image_url"`
	PurchaseURL string  `json:"purchase_url"`
	PDF         bool    `json:"pdf"`
}

type Author struct {
	Name string `json:"name"`
}

type Publisher struct {
	PublisherName string `json:"publisher_name"`
}

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

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	var br BooksReq
	br.baseURL = "https://www.googleapis.com/books/v1"
	br.Path = "/volumes"
	br.SearchTerm = "emily-st-john-mandel" // TODO: write function to get value from a param
	br.Params = []string{"projection", "q", "key"}
	br.ParamVals = []string{"lite", br.SearchTerm, getEnvMap()["GOOGLE_API_KEY"]}

	getBooks(br)
}
