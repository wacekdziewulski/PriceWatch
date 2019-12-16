package web

import (
	"PriceWatch/model"
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/johnreutersward/opengraph"
)

// ScrapePage extracts the price data from OpenGraph data of a chinese store under a certain product url
func ScrapePage(url string) model.PriceData {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to scrape url: '" + url + "', because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
		return model.PriceData{}
	}
	defer resp.Body.Close()

	md, _ := opengraph.Extract(resp.Body)
	if err != nil {
		log.Println("Failed to extract OpenGraph data because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
		return model.PriceData{}
	}

	data := model.PriceData{}
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

// DownloadImage downloads an image from a certain url
func DownloadImage(imageURL string) string {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Println("Failed to download image from: '" + imageURL + "', because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String()
}
