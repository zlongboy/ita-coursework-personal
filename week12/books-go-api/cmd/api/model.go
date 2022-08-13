package main

type Publisher struct {
	ID            string `json:"id"`
	PublisherName string `json:"publisher_name"`
}

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
