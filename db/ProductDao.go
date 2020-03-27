package db

import (
	"PriceWatch/model"

	"github.com/sirupsen/logrus"
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
	if !dao.isProductInDatabase(priceData.URL) {
		logrus.Info("Adding new product: ", priceData.Title)
		query := "INSERT INTO `products` (`url`, `affiliate_link`, `image_url`, `site_id`, `title`, `image`) VALUES (?,?,?,?,?,?)"
		return dao.connector.InsertData(query, priceData.URL, priceData.AffiliateShortURL, priceData.ImageURL, 1, priceData.Title, priceData.ImageData)
	}

	logrus.Info("Product is already tracked: ", priceData.Title)
	return true
}

// IsProductInDatabase checks if a certain product under URL is already in the database or not
func (dao *ProductDao) isProductInDatabase(url string) bool {
	query := "SELECT COUNT(*) FROM `products` WHERE `url` = ?"
	return dao.connector.IsInDatabase(query, url)
}

//func (dao *ProductDao) findProduct(URL string) sql.NullString {
//	var s sql.NullString
//	err := dao.connector.db.QueryRow("SELECT id FROM `products` WHERE url = ?", URL).Scan(&s)
//	if err != nil {
//		log.Printf("Couldn't find priceData entry in DB for url: %s", URL)
//	}
//	return s
//}
