package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tkanos/gonfig"
)

const configurationFile = "./pricewatch.json"

// ServerConfiguration defines where and how the server is deployed
type ServerConfiguration struct {
	Scheme string `env:"PRICEWATCH_SERVER_SCHEME"`
	Host   string `env:"PRICEWATCH_SERVER_HOST"`
	Port   int    `env:"PRICEWATCH_SERVER_PORT"`
}

// DatabaseConfiguration contains the data to connect to the database
type DatabaseConfiguration struct {
	Type     string `env:"PRICEWATCH_DB_TYPE"`
	Host     string `env:"PRICEWATCH_DB_HOST"`
	Port     int    `env:"PRICEWATCH_DB_PORT"`
	Protocol string `env:"PRICEWATCH_DB_PROTOCOL"`
	User     string `env:"PRICEWATCH_DB_USER"`
	Password string `env:"PRICEWATCH_DB_PASSWORD"`
	Schema   string `env:"PRICEWATCH_DB_SCHEMA"`
}

// Configuration contains the PriceWatch settings
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// NewConfiguration creates the configuration object
func NewConfiguration() *Configuration {
	log.Println("Reading configuration from: " + configurationFile)

	configuration := &Configuration{}
	err := gonfig.GetConf(configurationFile, configuration)
	if err != nil {
		panic(err)
	}

	log.Println(configuration.String())

	return configuration
}

func (sc *ServerConfiguration) String() string {
	return fmt.Sprintf("ServerConfiguration: %v, %v, %d", sc.Scheme, sc.Host, sc.Port)
}

func (dc *DatabaseConfiguration) String() string {
	return fmt.Sprintf("DatabaseConfiguration: %v, %v, %d, %v, %v, %v", dc.Type, dc.Host, dc.Port, dc.User, dc.Password, dc.Schema)
}

// String implements the Stringer interface for a string representation
func (c *Configuration) String() string {
	return fmt.Sprintf("Read Configuration: %v, %v", c.Server, c.Database)
}

// GetServerURL returns the server full URL with scheme, hostname and port
func (c *Configuration) GetServerURL() string {
	return c.Server.Host + ":" + strconv.Itoa(c.Server.Port)
}

// GetDatabaseConnectionString returns a connection string in a format accepted by database engines
func (c *Configuration) GetDatabaseConnectionString() string {
	return c.Database.User + ":" + c.Database.Password + "@" + c.Database.Protocol + "(" + c.Database.Host + ":" + strconv.Itoa(c.Database.Port) + ")/" + c.Database.Schema
}
