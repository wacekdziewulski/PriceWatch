package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// PriceResource manages endpoints and database operations regarding pricing
type PriceResource struct {
	productDao *ProductDao
}

// NewPriceResource creates a new PriceResource structure
func NewPriceResource(productDao *ProductDao) *PriceResource {
	return &PriceResource{productDao}
}

func (resource *PriceResource) checkPrice(w http.ResponseWriter, r *http.Request) {
	requestedPage := string(r.URL.Query()["url"][0])
	log.Println("Price check for url: " + requestedPage)

	var priceData = scrapePage(requestedPage)
	log.Println(priceData)
	resource.productDao.AddProduct(priceData, 1)

	json.NewEncoder(w).Encode(priceData)
}
