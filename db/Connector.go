package db

import (
	"PriceWatch/configuration"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Connector manages database connectivity
type Connector struct {
	database *sql.DB
}

// NewConnector creates and initializes a new database connection
func NewConnector(configuration *configuration.Configuration) *Connector {
	connector := &Connector{}

	dbType := configuration.Database.Type
	dbConnectionString := configuration.GetDatabaseConnectionString()
	db, err := sql.Open(dbType, dbConnectionString)

	if err != nil {
		log.Println("Failed to open database connection: " + err.Error())
		return nil
	}

	connector.database = db

	log.Println("Connected to database on: " + configuration.Database.Host + ", port: " + strconv.Itoa(configuration.Database.Port))

	return connector
}

// CloseConnection closes the Database Connection
func (connector *Connector) CloseConnection() {
	if connector.database != nil {
		connector.database.Close()
	}
}
