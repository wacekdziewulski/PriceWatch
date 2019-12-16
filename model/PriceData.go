package model

// PriceData Scraped data from Gearbest or Banggood containing the price, image url, product name etc.
type PriceData struct {
	Site          string  `json:"site"`
	Title         string  `json:"title"`
	URL           string  `json:"url"`
	PriceAmount   float64 `json:"price_amount"`
	PriceCurrency string  `json:"price_currency"`
	ImageURL      string  `json:"image_url"`
}