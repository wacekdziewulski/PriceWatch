package strategy

import "github.com/sirupsen/logrus"

// PriceStrategyFactory defines the factory for creating PriceStrategies
type PriceStrategyFactory struct{}

// NewPriceStrategyFactory creates a new PriceStrategyFactory
func NewPriceStrategyFactory() *PriceStrategyFactory {
	return &PriceStrategyFactory{}
}

// ProvidePriceScrapingStrategy creates the necessary strategy for scraping price data based on the type of store that is being queried
func (f *PriceStrategyFactory) ProvidePriceScrapingStrategy(storeType string) PriceStrategy {
	logrus.Infof("Store: %s", storeType)
	if storeType == "Banggood" {
		return NewJSONLdPriceStrategy()
	}
	return NewOpenGraphPriceStrategy()
}
