package db

import (
	"PriceWatch/model"
)

// PriceDao handles the database operations regarding adding Product data
type PriceDao struct {
	connector *Connector
}

// NewPriceDao creates a new DAO for managing Products
func NewPriceDao(connector *Connector) *PriceDao {
	return &PriceDao{connector}
}

// AddPrice adds a PriceData object to the Product table in the database
func (dao *PriceDao) AddPrice(priceData *model.PriceData) bool {
	query := "INSERT INTO `prices` (`product_id`, `price`, `currency`) VALUES ((SELECT `id` FROM `products` WHERE url = ?), ?, ?)"
	return dao.connector.InsertData(query, priceData.URL, priceData.PriceAmount, priceData.PriceCurrency)
}
