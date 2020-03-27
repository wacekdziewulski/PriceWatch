package service

import (
	"PriceWatch/db"
	"PriceWatch/model"
	"PriceWatch/web"

	"github.com/sirupsen/logrus"
)

// PriceService - handles adding new product, or just updating the price of an existing one
type PriceService struct {
	productDao           *db.ProductDao
	priceDao             *db.PriceDao
	urlShorteningService *URLShorteningService
}

// NewPriceService creates a new PriceService structure
func NewPriceService(productDao *db.ProductDao, priceDao *db.PriceDao, urlShorteningService *URLShorteningService) *PriceService {
	return &PriceService{productDao, priceDao, urlShorteningService}
}

// addPrice adds a new web scraped PriceData into the Database
func (service *PriceService) addProduct(priceData *model.PriceData) bool {
	addedSuccessfully := service.productDao.AddProduct(priceData)

	if addedSuccessfully {
		logrus.Infof("Added product: %s with price: %.2f %s", priceData.Title, priceData.PriceAmount, priceData.PriceCurrency)
	} else {
		logrus.Warn("Failed to add entry to the database: ", priceData)
	}

	return addedSuccessfully
}

// addPrice adds a new web scraped PriceData into the Database
func (service *PriceService) addPrice(priceData *model.PriceData) bool {
	addedSuccessfully := service.priceDao.AddPrice(priceData)

	if addedSuccessfully {
		logrus.Infof("Added price: %.2f %s for product: %s", priceData.PriceAmount, priceData.PriceCurrency, priceData.Title)
	} else {
		logrus.Warn("Failed to add entry to the database: ", priceData)
	}

	return addedSuccessfully
}

// AddProductPriceByURL adds product to Database and returns the priceData result
func (service *PriceService) AddProductPriceByURL(urlContext *model.URLContext) *model.PriceData {

	// (/) 1. Check if the product url is in the Database
	// 2. If we have it, check the time of last update and see if makes sense to query the web page again
	// 3a. If we have just get the data from the database and present it to the user
	// 3b. If we haven't updated for a while, download the price data
	// 4b. If we don't have an image yet, use the ImageURL to download the image
	// 5. Attach the image data upon adding the entry to the database
	// 6. Update the price of the product if we already have a product entry
	priceData := <-web.ExtractPriceDataFromURL(urlContext.URL.String())
	priceData.ImageData = <-web.DownloadImage(priceData.ImageURL)
	priceData.AffiliateShortURL = urlContext.AffiliateShortURL

	service.addProduct(&priceData)
	service.addPrice(&priceData)

	return &priceData
}
