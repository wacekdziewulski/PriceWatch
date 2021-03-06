package service

import (
	"PriceWatch/configuration"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
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

	logrus.Info("Shortening URL for: ", URL)

	requestBody, err := json.Marshal(map[string]string{
		"long_url":   URL,
		"group_guid": service.configuration.URLShortening.GroupGUID,
	})

	url := service.configuration.URLShortening.Scheme + "://" + service.configuration.URLShortening.APIHost + service.configuration.URLShortening.APIPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Host", service.configuration.URLShortening.APIHost)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+service.configuration.URLShortening.AccessToken)

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		dumpedReq, _ := httputil.DumpRequest(req, true)
		logrus.Debug(dumpedReq)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		logrus.Warnf("Failed to shorten URL: %s, because: %+v", URL, err)
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
