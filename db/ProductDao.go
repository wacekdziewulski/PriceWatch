package db

import (
	"PriceWatch/model"
	"log"
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
	log.Println("Running SQL query: " + query)
	log.Println(priceData)
	statement, err := dao.connector.getDb().Prepare(query)

	if err != nil {
		log.Print("Failed to prepare DB Statement:" + query)
		log.Println(err)
		return false
	}

	result, err := statement.Exec(priceData.URL, priceData.AffiliateLink, priceData.ImageURL, 1, priceData.Title, priceData.ImageData)

	if err != nil {
		log.Print("Failed to execute DB Statement:" + query)
		log.Println(err)
		return false
	}
	defer statement.Close()

	rows, err := result.RowsAffected()
	if err != nil {
		log.Print("Failed to insert DB data")
		log.Println(err)
		return false
	}

	if rows > 0 {
		return true
	}

	return false
}

//func (dao *ProductDao) findProduct(URL string) sql.NullString {
//	var s sql.NullString
//	err := dao.connector.db.QueryRow("SELECT id FROM `products` WHERE url = ?", URL).Scan(&s)
//	if err != nil {
//		log.Printf("Couldn't find priceData entry in DB for url: %s", URL)
//	}
//	return s
//}
