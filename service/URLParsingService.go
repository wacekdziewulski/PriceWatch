package service

import (
	"PriceWatch/db"
	"PriceWatch/model"
	"net/url"
	"strings"
)

// URLParsingService handles processing of raw URLs and produces the URLContext
type URLParsingService struct {
	storeDao             *db.StoreDao
	urlShorteningService *URLShorteningService
}

// NewURLParsingService creates a new URLParsingService structure
func NewURLParsingService(storeDao *db.StoreDao, urlShorteningService *URLShorteningService) *URLParsingService {
	return &URLParsingService{storeDao, urlShorteningService}
}

// CreateURLContext creates the URLContext object from just the URL given
func (service *URLParsingService) CreateURLContext(rawurl string) model.URLContext {
	var context model.URLContext = model.URLContext{}
	context.URL, _ = url.Parse(rawurl)
	context.AffiliateURL = service.createAffiliateURL(context.URL)
	context.AffiliateShortURL = <-service.urlShorteningService.ShortenURL(context.AffiliateURL.String())
	return context
}

func (service *URLParsingService) createAffiliateURL(input *url.URL) *url.URL {
	affiliateURL := url.URL{Scheme: "https", Host: input.Host}

	var host string = input.Host
	if strings.HasPrefix(host, "m.") || strings.HasPrefix(host, "www.") {
		hostFragments := strings.Split(host, ".")
		host = strings.Join(hostFragments[1:], ".")
	}

	storeData := service.storeDao.GetStoreDataByHostname(host)

	values := url.Values{}
	values.Add(storeData.AffiliateParam, storeData.AffiliateValue)
	affiliateURL.RawQuery = values.Encode()

	return &affiliateURL
}
