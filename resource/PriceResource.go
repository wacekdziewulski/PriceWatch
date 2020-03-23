package resource

import (
	"PriceWatch/service"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PriceResource manages endpoints and database operations regarding pricing
type PriceResource struct {
	priceService *service.PriceService
}

// NewPriceResource creates a new PriceResource structure
func NewPriceResource(priceService *service.PriceService) *PriceResource {
	return &PriceResource{priceService}
}

// CheckPrice checks the product price for the chinese store url given
func (resource *PriceResource) CheckPrice(w http.ResponseWriter, r *http.Request) {
	requestedPage := string(r.URL.Query()["url"][0])
	logrus.Info("Price check for url: ", requestedPage)

	priceData := resource.priceService.AddProductPriceByURL(requestedPage)

	json.NewEncoder(w).Encode(priceData)
}
