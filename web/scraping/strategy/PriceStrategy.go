package strategy

import "PriceWatch/model"

// PriceStrategy defines how a certain web scraper should process web page data to extract the PriceData object
type PriceStrategy interface {
	FillPriceData(pageBody string, priceData *model.PriceData)
}
