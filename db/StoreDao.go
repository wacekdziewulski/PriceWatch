package db

import "PriceWatch/model"

// StoreDao handles the database operations regarding reading of the Store data
type StoreDao struct {
	connector *Connector
}

// NewStoreDao creates a new DAO for managing Stores
func NewStoreDao(connector *Connector) *StoreDao {
	return &StoreDao{connector}
}

// GetStoreDataByHostname gets the store data from the database based on the hostname from the URL
func (dao *StoreDao) GetStoreDataByHostname(hostname string) *model.StoreData {
	query := "SELECT `name`, `affiliate_param`, `affiliate_value` FROM `stores` WHERE `url_matcher` = ?"
	return dao.connector.GetStoreFromDatabase(query, hostname)
}
