package main

type RequestConfig struct {
	baseURL   string
	Path      string
	APIKey    string
	Params    []string
	ParamVals []string
}

type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Subtitle    string  `json:"subtitle"`
	Author      string  `json:"author"`
	Publisher   string  `json:"publisher"`
	Desc        string  `json:"desc"`
	PublishDate string  `json:"publish_date"`
	Country     string  `json:"country"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"image_url"`
	PurchaseURL string  `json:"purchase_url"`
	PDF         bool    `json:"pdf"`
}

// ** FOR TESTING ** //

type Publisher struct {
	ID            string `json:"id"`
	PublisherName string `json:"publisher_name"`
}
