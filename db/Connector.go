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
func (connector *Connector) InsertData(statement string, args ...interface{}) int64 {
	result, err := connector.database.Exec(statement, args)
	if err != nil {
		log.Print("Failed to execute DB Statement:" + statement)
		log.Println(err)
		return 0
	}
	rows, _ := result.RowsAffected()
	if rows > 0 {
		log.Println("Inserted data to database with rows: " + string(rows))
	}

	return 1

	//stmt, err := connector.database.Prepare(statement)
	//if err != nil {
	//log.Fatal("Failed to prepare DB Statement:" + statement)
	//log.Fatalln(err.Error())
	//return 0
	//}

	//result, err1 := stmt.Exec(args)
	//if err1 != nil {
	//log.Fatal("Failed to execute DB Query:" + statement)
	//log.Fatalln(err.Error())
	//return 0
	//}

	//rowsAffected, err2 := result.RowsAffected()
	//if err2 != nil {
	//	log.Fatal("Failed to return DB Affected rows")
	//	log.Fatalln(err.Error())
	//	return 0
	//}

	//return rowsAffected
}

// CloseConnection closes the Database Connection
func (connector *Connector) CloseConnection() {
	if connector.database != nil {
		connector.database.Close()
	}
}
