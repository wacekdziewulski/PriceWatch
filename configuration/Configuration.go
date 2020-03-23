package configuration

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
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

// URLShorteningConfiguration contains settings for bit.ly to shorten the affiliate links
type URLShorteningConfiguration struct {
	AccessToken string `env:"PRICEWATCH_URL_SHORTENER_ACCESS_TOKEN"`
	APIPath     string `env:"PRICEWATCH_URL_SHORTENER_API_PATH"`
	APIHost     string `env:"PRICEWATCH_URL_SHORTENER_API_HOST"`
	Scheme      string `env:"PRICEWATCH_URL_SHORTENER_API_SCHEME"`
	GroupGUID   string `env:"PRICEWATCH_URL_SHORTENER_GROUP_GUID"`
}

// Configuration contains the PriceWatch settings
type Configuration struct {
	Server        ServerConfiguration
	Database      DatabaseConfiguration
	URLShortening URLShorteningConfiguration
}

// NewConfiguration creates the configuration object
func NewConfiguration() *Configuration {
	logrus.Info("Reading configuration from: ", configurationFile)

	configuration := &Configuration{}
	err := gonfig.GetConf(configurationFile, configuration)
	if err != nil {
		panic(err)
	}

	logrus.Debug(configuration.String())

	return configuration
}

func (sc *ServerConfiguration) String() string {
	return fmt.Sprintf("ServerConfiguration: %v, %v, %d", sc.Scheme, sc.Host, sc.Port)
}

func (dc *DatabaseConfiguration) String() string {
	return fmt.Sprintf("DatabaseConfiguration: %v, %v, %d, %v, %v, %v", dc.Type, dc.Host, dc.Port, dc.User, dc.Password, dc.Schema)
}

func (usc *URLShorteningConfiguration) String() string {
	return fmt.Sprintf("URL Shortening Configuration: %v, %v, %v, %v, %v", maskSensitiveInformation(usc.AccessToken), usc.Scheme, usc.APIHost, usc.APIPath, maskSensitiveInformation(usc.GroupGUID))
}

// String implements the Stringer interface for a string representation
func (c *Configuration) String() string {
	return fmt.Sprintf("Read Configuration: %v, %v, %v", c.Server, c.Database, c.URLShortening)
}

// GetServerURL returns the server full URL with scheme, hostname and port
func (c *Configuration) GetServerURL() string {
	return c.Server.Host + ":" + strconv.Itoa(c.Server.Port)
}

// GetDatabaseConnectionString returns a connection string in a format accepted by database engines
func (c *Configuration) GetDatabaseConnectionString() string {
	return c.Database.User + ":" + c.Database.Password + "@" + c.Database.Protocol + "(" + c.Database.Host + ":" + strconv.Itoa(c.Database.Port) + ")/" + c.Database.Schema
}

func maskSensitiveInformation(input string) string {
	re := regexp.MustCompile(`(.{2}).*`)
	return re.ReplaceAllString(input, `$1\*\*\*\*\*`)
}
