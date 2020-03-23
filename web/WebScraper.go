package web

import (
	"PriceWatch/model"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/johnreutersward/opengraph"
	"github.com/sirupsen/logrus"
)

// ScrapePage extracts the price data from OpenGraph data of a chinese store under a certain product url
func ScrapePage(url string) <-chan model.PriceData {
	priceData := make(chan model.PriceData, 1)

	resp, err := http.Get(url)

	if err != nil {
		bytes, _ := httputil.DumpResponse(resp, true)
		logrus.Warnf("Failed to scrape url: %s, because of: %+v. HttpResponse: %s", url, err, bytes)
		priceData <- model.PriceData{}
		return priceData
	}
	defer resp.Body.Close()

	md, _ := opengraph.Extract(resp.Body)
	if err != nil {
		logrus.Warnf("Failed to extract OpenGraph data because of: %+v", err)
		priceData <- model.PriceData{}
		return priceData
	}

	data := model.PriceData{}
	for i := range md {
		logrus.Debugf("Found OpenGraph: %s = %s", md[i].Property, md[i].Content)

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

	priceData <- data

	return priceData
}
