package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/johnreutersward/opengraph"
)

// PriceData Scraped data from Gearbest or Banggood containing the price, image url, product name etc.
type PriceData struct {
	Site          string  `json:"site"`
	Title         string  `json:"title"`
	URL           string  `json:"url"`
	PriceAmount   float64 `json:"price_amount"`
	PriceCurrency string  `json:"price_currency"`
	ImageURL      string  `json:"image_url"`
}

func scrapePage(url string) PriceData {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to scrape url: '" + url + "', because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
		return PriceData{}
	}
	defer resp.Body.Close()

	md, _ := opengraph.Extract(resp.Body)
	if err != nil {
		log.Println("Failed to extract OpenGraph data because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
		return PriceData{}
	}

	data := PriceData{}
	for i := range md {
		log.Printf("Found OpenGraph: %s = %s\n", md[i].Property, md[i].Content)

		switch md[i].Property {
		case "site_name":
			data.Site = md[i].Content
		case "title":
			data.Title = md[i].Content
		case "image":
			data.ImageURL = md[i].Content
		case "url":
			data.URL = md[i].Content
		case "price:amount":
			data.PriceAmount, _ = strconv.ParseFloat(md[i].Content, 32)
		case "price:currency":
			data.PriceCurrency = md[i].Content
		}
	}

	return data
}
