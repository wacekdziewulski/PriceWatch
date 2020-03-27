package model

import "fmt"

// PriceData Scraped data from Gearbest or Banggood containing the price, image url, product name etc.
type PriceData struct {
	Site              string  `json:"site"`
	Title             string  `json:"title"`
	URL               string  `json:"url"`
	PriceAmount       float64 `json:"price_amount"`
	PriceCurrency     string  `json:"price_currency"`
	ImageURL          string  `json:"image_url"`
	ImageData         string
	AffiliateShortURL string
}

func (priceData PriceData) String() string {
	return fmt.Sprintf("Site: %v, URL: %v, Title: %v, Price: %f %v, AffiliateLink: %v", priceData.Site, priceData.URL, priceData.Title, priceData.PriceAmount, priceData.PriceCurrency, priceData.AffiliateShortURL)
}
