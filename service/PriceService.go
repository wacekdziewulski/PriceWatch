package service

import (
	"PriceWatch/db"
	"PriceWatch/model"
	"PriceWatch/web"
	"log"
	"os"
)

// PriceService - handles adding new product, or just updating the price of an existing one
type PriceService struct {
	productDao *db.ProductDao
}

// NewPriceService creates a new PriceService structure
func NewPriceService(productDao *db.ProductDao) *PriceService {
	return &PriceService{productDao}
}

// AddPrice adds a new web scraped PriceData into the Database
func (service *PriceService) AddPrice(priceData model.PriceData) {
	service.productDao.AddProduct(priceData)
}

// AddProductPriceByURL adds product to Database and returns the priceData result
func (service *PriceService) AddProductPriceByURL(url string) *model.PriceData {

	// 1. Check if the product url is in the Database
	// 2. If we have it, check the time of last update and see if makes sense to query the web page again
	// 3a. If we have just get the data from the database and present it to the user
	// 3b. If we haven't updated for a while, download the price data
	// 4b. If we don't have an image yet, use the ImageURL to download the image
	// 5. Attach the image data upon adding the entry to the database
	// 6. Update the price of the product if we already have a product entry
	var priceData = web.ScrapePage(url)
	log.Println(priceData)

	outputFileName := "/tmp/PriceWatch.jpg"

	f, err := os.Create(outputFileName)
	if err == nil {
		defer f.Close()
		log.Println("Downloaded image from: " + priceData.ImageURL + " into " + outputFileName)
		f.WriteString(web.DownloadImage(priceData.ImageURL))
		f.Sync()
	}

	service.AddPrice(priceData)

	return &priceData
}
