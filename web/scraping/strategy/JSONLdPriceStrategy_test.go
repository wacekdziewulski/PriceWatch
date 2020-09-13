package strategy

import (
	"PriceWatch/model"
	"io/ioutil"
	"testing"
)

func TestJSONLdPriceStrategy(*testing.T) {
	pageContent, _ := ioutil.ReadFile("test/data/banggood-emax-tinyhawk-2.html")
	jsonStrategy := NewJSONLdPriceStrategy()
	priceData := model.PriceData{}
	jsonStrategy.FillPriceData(string(pageContent), &priceData)
}
