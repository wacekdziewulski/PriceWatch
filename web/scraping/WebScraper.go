package scraping

import (
	"PriceWatch/model"
	"PriceWatch/web/scraping/strategy"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
)

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
func ExtractPriceDataFromURL(urlContext *model.URLContext) <-chan model.PriceData {
	priceData := make(chan model.PriceData, 1)

	pageContents := <-scrapePageContents(urlContext.URL.String())
	data := model.PriceData{}

	if pageContents != "" {
		priceScrapingStrategy := strategy.NewPriceStrategyFactory().ProvidePriceScrapingStrategy(urlContext.StoreName)
		priceScrapingStrategy.FillPriceData(pageContents, &data)
	}

	priceData <- data

	return priceData
}
