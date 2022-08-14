package main

type RequestConfig struct {
	baseURL   string
	Path      string
	APIKey    string
	Params    []string
	ParamVals []string
}

type Book struct {
	Kind       string `json:"kind,omitempty"`
	TotalItems int    `json:"totalItems,omitempty"`
	Items      []struct {
		Kind       string `json:"kind,omitempty"`
		ID         string `json:"id,omitempty"`
		Etag       string `json:"etag,omitempty"`
		SelfLink   string `json:"selfLink,omitempty"`
		VolumeInfo struct {
			Title         string   `json:"title,omitempty"`
			Subtitle      string   `json:"subtitle,omitempty"`
			Authors       []string `json:"authors,omitempty"`
			Publisher     string   `json:"publisher,omitempty"`
			PublishedDate string   `json:"publishedDate,omitempty"`
			Description   string   `json:"description,omitempty"`
			ReadingModes  struct {
				Text  bool `json:"text,omitempty"`
				Image bool `json:"image,omitempty"`
			} `json:"readingModes,omitempty"`
			MaturityRating      string `json:"maturityRating,omitempty"`
			AllowAnonLogging    bool   `json:"allowAnonLogging,omitempty"`
			ContentVersion      string `json:"contentVersion,omitempty"`
			PanelizationSummary struct {
				ContainsEpubBubbles  bool `json:"containsEpubBubbles,omitempty"`
				ContainsImageBubbles bool `json:"containsImageBubbles,omitempty"`
			} `json:"panelizationSummary,omitempty"`
			ImageLinks struct {
				SmallThumbnail string `json:"smallThumbnail,omitempty"`
				Thumbnail      string `json:"thumbnail,omitempty"`
			} `json:"imageLinks,omitempty"`
			PreviewLink         string `json:"previewLink,omitempty"`
			InfoLink            string `json:"infoLink,omitempty"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink,omitempty"`
		} `json:"volumeInfo,omitempty"`
		SaleInfo struct {
			Country   string `json:"country,omitempty"`
			ListPrice struct {
				Amount       float64 `json:"amount,omitempty"`
				CurrencyCode string  `json:"currencyCode,omitempty"`
			} `json:"listPrice,omitempty"`
			RetailPrice struct {
				Amount       float64 `json:"amount,omitempty"`
				CurrencyCode string  `json:"currencyCode,omitempty"`
			} `json:"retailPrice,omitempty"`
			BuyLink string `json:"buyLink,omitempty"`
			Offers  []struct {
				FinskyOfferType int `json:"finskyOfferType,omitempty"`
				ListPrice       struct {
					AmountInMicros int    `json:"amountInMicros,omitempty"`
					CurrencyCode   string `json:"currencyCode,omitempty"`
				} `json:"listPrice,omitempty"`
				RetailPrice struct {
					AmountInMicros int    `json:"amountInMicros,omitempty"`
					CurrencyCode   string `json:"currencyCode,omitempty"`
				} `json:"retailPrice,omitempty"`
				Giftable bool `json:"giftable,omitempty"`
			} `json:"offers,omitempty"`
		} `json:"saleInfo,omitempty"`
		AccessInfo struct {
			Country string `json:"country,omitempty"`
			Epub    struct {
				IsAvailable  bool   `json:"isAvailable,omitempty"`
				AcsTokenLink string `json:"acsTokenLink,omitempty"`
			} `json:"epub,omitempty"`
			Pdf struct {
				IsAvailable bool `json:"isAvailable,omitempty"`
			} `json:"pdf,omitempty"`
			AccessViewStatus string `json:"accessViewStatus,omitempty"`
		} `json:"accessInfo,omitempty"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet,omitempty"`
		} `json:"searchInfo,omitempty"`
	} `json:"items,omitempty"`
}

const (
	RespInvalidAuth   string = "Invalid API Key"
	RespMissingAuthor string = "Missing required 'author' param"
)

// ** FOR TESTING ** //

type Publisher struct {
	ID            string `json:"id"`
	PublisherName string `json:"publisher_name"`
}
