package model

import "net/url"

// URLContext describes the URL passed to the application through the resource as well as metadata - e.g. the affiliate code or the store name
type URLContext struct {
	URL               *url.URL
	StoreName         string
	AffiliateURL      *url.URL
	AffiliateShortURL string
}
