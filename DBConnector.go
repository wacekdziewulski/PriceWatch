package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// DBConnector manages database connectivity
type DBConnector struct {
	db *sql.DB
}

// NewDbConnector creates and initializes a new database connection
func NewDbConnector(configuration *Configuration) *DBConnector {
	dbConnector := &DBConnector{}

	dbType := configuration.Database.Type
	dbConnectionString := configuration.GetDatabaseConnectionString()
	db, err := sql.Open(dbType, dbConnectionString)

	if err != nil {
		log.Println("Failed to open database connection: " + err.Error())
		return nil
	}

	dbConnector.db = db

	log.Println("Connected to database on: " + configuration.Database.Host + ", port: " + strconv.Itoa(configuration.Database.Port))

	return dbConnector
}

// CloseConnection closes the Database Connection
func (connector *DBConnector) CloseConnection() {
	if connector.db != nil {
		connector.db.Close()
	}
}
