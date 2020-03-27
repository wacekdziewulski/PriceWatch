package resource

import (
	"PriceWatch/service"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PriceResource manages endpoints and database operations regarding pricing
type PriceResource struct {
	priceService      *service.PriceService
	urlParsingService *service.URLParsingService
}

// NewPriceResource creates a new PriceResource structure
func NewPriceResource(urlParsingService *service.URLParsingService, priceService *service.PriceService) *PriceResource {
	return &PriceResource{priceService, urlParsingService}
}

// CheckPrice checks the product price for the chinese store url given
func (resource *PriceResource) CheckPrice(w http.ResponseWriter, r *http.Request) {
	requestedPage := string(r.URL.Query()["url"][0])
	logrus.Info("Price check for url: ", requestedPage)

	urlContext := resource.urlParsingService.CreateURLContext(requestedPage)
	priceData := resource.priceService.AddProductPriceByURL(&urlContext)

	w.Write([]byte(fmt.Sprintf("Product: %s with price: %.2f %s added!", priceData.Title, priceData.PriceAmount, priceData.PriceCurrency)))
}
