package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client *http.Client

func getBooks(rc RequestConfig) []BookValues {

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

	// BUILD REQUEST
	req, err := http.NewRequest("GET", URI.String(), nil)
	req.Header.Add("Content-Type", "application/json")

	// SEND REQUEST
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s \n", err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("Google Books response: %v \n", resp.Status)

	// PARSE RESPONSE
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s \n", err.Error())
	}

	var resObj Books
	json.Unmarshal(resBody, &resObj)

	// resultsTotal := resObj.TotalItems
	fmt.Println(extractBooks(resObj))

	return extractBooks(resObj)
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

func extractBooks(ob Books) []BookValues {
	items := ob.Items
	var bv []BookValues

	for _, b := range items {
		var s BookValues
		s.ID = b.ID
		s.Author = parseAuthor(b.VolumeInfo.Authors)
		s.Publisher = b.VolumeInfo.Publisher
		s.Title = b.VolumeInfo.Title
		s.Subtitle = b.VolumeInfo.Subtitle
		s.Desc = b.VolumeInfo.Description
		s.PublishDate = b.VolumeInfo.PublishedDate
		s.Country = b.SaleInfo.Country
		s.Price = float32(b.SaleInfo.ListPrice.Amount)
		s.ImageURL = b.VolumeInfo.ImageLinks.Thumbnail
		s.PurchaseURL = b.SaleInfo.BuyLink
		s.PDF = b.AccessInfo.Pdf.IsAvailable

		bv = append(bv, s)
	}
	return bv
}

func parseAuthor(a []string) string {
	if len(a) > 0 {
		return a[0]
	}

	return "NO AUTHOR FOUND"
}
