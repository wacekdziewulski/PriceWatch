package db

import (
	"PriceWatch/configuration"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("Failed to open database connection because: %+v", err)
		return nil
	}

	connector.database = db

	logrus.Infof("Connected to database on: %s, port: %d", configuration.Database.Host, configuration.Database.Port)

	return connector
}

func (connector *Connector) getDb() *sql.DB {
	return connector.database
}

// InsertData puts data into the database using an "insert into" statement
func (connector *Connector) InsertData(query string, args ...interface{}) bool {
	logrus.Debug("Running SQL query: ", query)
	statement, err := connector.getDb().Prepare(query)

	if err != nil {
		logrus.Warnf("Failed to prepare DB Statement for query: %s, because: %+v", query, err)
		return false
	}

	result, err := statement.Exec(args...)

	if err != nil {
		logrus.Warnf("Failed to execute DB Statement for query: %s, because: %+v", query, err)
		return false
	}
	defer statement.Close()

	rows, err := result.RowsAffected()
	if err != nil {
		logrus.Warnf("Failed to insert DB data for query: %s, because: %+v", query, err)
		return false
	}

	if rows > 0 {
		return true
	}

	return false
}

// IsInDatabase is used for SELECT COUNT(*) queries to express if an entry is in the database or not
func (connector *Connector) IsInDatabase(query string, args ...interface{}) bool {
	rows, err := connector.getDb().Query(query, args...)
	if err != nil {
		logrus.Warnf("Failed to get DB row count for query: %s, because: %+v", query, err)
		return true
	}
	defer rows.Close()

	var count int32

	if rows.Next() {
		rows.Scan(&count)
	}

	return count > 0
}

// CloseConnection closes the Database Connection
func (connector *Connector) CloseConnection() {
	if connector.database != nil {
		connector.database.Close()
	}
}
