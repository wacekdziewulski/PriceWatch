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

func (connector *Connector) getDb() *sql.DB {
	return connector.database
}

// InsertData puts data into the database using an "insert into" statement
func (connector *Connector) InsertData(query string, args ...interface{}) bool {
	log.Println("Running SQL query: " + query)
	statement, err := connector.getDb().Prepare(query)

	if err != nil {
		log.Print("Failed to prepare DB Statement:" + query)
		log.Println(err)
		return false
	}

	result, err := statement.Exec(args...)

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

// CloseConnection closes the Database Connection
func (connector *Connector) CloseConnection() {
	if connector.database != nil {
		connector.database.Close()
	}
}
