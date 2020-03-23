package db

import (
	"PriceWatch/model"
)

// ProductDao handles the database operations regarding adding Product data
type ProductDao struct {
	connector *Connector
}

// NewProductDao creates a new DAO for managing Products
func NewProductDao(connector *Connector) *ProductDao {
	return &ProductDao{connector}
}

// AddProduct adds a PriceData object to the Product table in the database
func (dao *ProductDao) AddProduct(priceData *model.PriceData) bool {
	query := "INSERT INTO `products` (`url`, `affiliate_link`, `image_url`, `site_id`, `title`, `image`) VALUES (?,?,?,?,?,?)"
	return dao.connector.InsertData(query, priceData.URL, priceData.AffiliateLink, priceData.ImageURL, 1, priceData.Title, priceData.ImageData)
}

// AddPrice adds a PriceData object to the Product table in the database
func (dao *ProductDao) AddPrice(priceData *model.PriceData) bool {
	query := "INSERT INTO `prices` (`product_id`, `price`, `currency`) VALUES ((SELECT `id` FROM `products` WHERE url = ?), ?, ?)"
	return dao.connector.InsertData(query, priceData.URL, priceData.PriceAmount, priceData.PriceCurrency)
}

//func (dao *ProductDao) findProduct(URL string) sql.NullString {
//	var s sql.NullString
//	err := dao.connector.db.QueryRow("SELECT id FROM `products` WHERE url = ?", URL).Scan(&s)
//	if err != nil {
//		log.Printf("Couldn't find priceData entry in DB for url: %s", URL)
//	}
//	return s
//}
