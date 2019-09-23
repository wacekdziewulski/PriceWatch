package main

import (
    "log"
    "io/ioutil"
    "fmt"
    "net/http"
    "strings"
    "github.com/dyatlov/go-opengraph/opengraph"
)

type PriceData struct {
    title string
    url string
    price int
    imageUrl int
}

func scrapePage(url string) PriceData {
    webSiteContents := getWebsiteContents(url)
    return extractPriceData(webSiteContents)
}

func getWebsiteContents(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to scrape url: '" + url + "', because: '" + err.Error() + "'")
        return ""
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    body := string(bodyBytes)

    return body
}

func extractPriceData(webPageContents string) PriceData {
    og := opengraph.NewOpenGraph()
    err := og.ProcessHTML(strings.NewReader(webPageContents))

    if err != nil {
        fmt.Println("Failed to extract OpenGraph data from: '" + webPageContents + "', because: '" + err.Error() + "'")
        return PriceData{}
    }

    priceData := PriceData{title: og.Title, url: og.URL}

    return priceData
}
