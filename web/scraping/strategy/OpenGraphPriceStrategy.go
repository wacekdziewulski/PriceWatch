package strategy

import (
	"PriceWatch/model"
	"strconv"
	"strings"

	"github.com/johnreutersward/opengraph"
	"github.com/sirupsen/logrus"
)

// OpenGraphPriceStrategy defines the strategy for processing OpenGraph data
type OpenGraphPriceStrategy struct {
}

// NewOpenGraphPriceStrategy returns the OpenGraphPriceStrategy
func NewOpenGraphPriceStrategy() *OpenGraphPriceStrategy {
	return &OpenGraphPriceStrategy{}
}

// FillPriceData extracts the prices from OpenGraph - implements PriceStrategy
func (p *OpenGraphPriceStrategy) FillPriceData(pageBody string, priceData *model.PriceData) {
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
