package web

import (
	"PriceWatch/model"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/johnreutersward/opengraph"
	"github.com/sirupsen/logrus"
)

func fillPriceDataBasedOnOpenGraph(pageBody string, priceData *model.PriceData) {
	md, err := opengraph.Extract(strings.NewReader(pageBody))

	if err != nil {
		logrus.Warnf("Failed to extract OpenGraph data because of: %+v", err)
		return
	}

	for i := range md {
		logrus.Debugf("Found OpenGraph: %s = %s", md[i].Property, md[i].Content)

		switch md[i].Property {
		case "site_name":
			priceData.Site = md[i].Content
		case "title":
			priceData.Title = md[i].Content
		case "image":
			priceData.ImageURL = md[i].Content
		case "url":
			priceData.URL = md[i].Content
		case "price:amount":
			priceData.PriceAmount, _ = strconv.ParseFloat(md[i].Content, 32)
		case "price:currency":
			priceData.PriceCurrency = md[i].Content
		}
	}
}

// ScrapePageContents gets the contents of a web page and returns it as a string
func scrapePageContents(url string) <-chan string {
	pageContents := make(chan string, 1)

	resp, err := http.Get(url)

	if err != nil {
		bytes, _ := httputil.DumpResponse(resp, true)
		logrus.Warnf("Failed to scrape url: %s, because of: %+v. HttpResponse: %s", url, err, bytes)
		pageContents <- ""
		return pageContents
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	pageContents <- string(bytes)

	return pageContents
}

// ExtractPriceDataFromURL extracts the price data from OpenGraph data of a chinese store under a certain product url
func ExtractPriceDataFromURL(url string) <-chan model.PriceData {
	priceData := make(chan model.PriceData, 1)

	pageContents := <-scrapePageContents(url)
	data := model.PriceData{}

	if pageContents != "" {
		fillPriceDataBasedOnOpenGraph(pageContents, &data)
	}

	priceData <- data

	return priceData
}
