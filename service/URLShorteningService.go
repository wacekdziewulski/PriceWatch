package service

import (
	"PriceWatch/configuration"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

// URLShorteningService proxies calls to bit.ly to shorten affiliate URLs
type URLShorteningService struct {
	configuration *configuration.Configuration
}

// NewURLShorteningService creates a new URLShorteningService
func NewURLShorteningService(configuration *configuration.Configuration) *URLShorteningService {
	return &URLShorteningService{configuration}
}

type bitlyResponse struct {
	Link string `json:"link"`
}

// ShortenURL sends a request to the URL shortening application (e.g. bit.ly) and returns the short URL
func (service *URLShorteningService) ShortenURL(URL string) <-chan string {
	shortURL := make(chan string, 1)

	log.Println("Shortening URL for: " + URL)

	requestBody, err := json.Marshal(map[string]string{
		"long_url":   URL,
		"group_guid": service.configuration.URLShortening.GroupGUID,
	})

	url := service.configuration.URLShortening.Scheme + "://" + service.configuration.URLShortening.APIHost + service.configuration.URLShortening.APIPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Host", service.configuration.URLShortening.APIHost)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+service.configuration.URLShortening.AccessToken)

	dumpedReq, _ := httputil.DumpRequest(req, true)
	log.Println(string(dumpedReq))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print("Failed to shorten URL: " + URL)
		log.Println(err)
		shortURL <- ""
		return shortURL
	}

	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	var bitlyResp bitlyResponse
	json.Unmarshal([]byte(response), &bitlyResp)

	shortURL <- bitlyResp.Link

	return shortURL
}
