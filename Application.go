package main

import (
	"PriceWatch/configuration"
	"PriceWatch/db"
	"PriceWatch/resource"
	"PriceWatch/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Application describes the whole PriceWatch application with endpoints, dao etc.
type Application struct {
	configuration        *configuration.Configuration
	connector            *db.Connector
	productDao           *db.ProductDao
	priceService         *service.PriceService
	urlShorteningService *service.URLShorteningService
	priceResource        *resource.PriceResource
}

// NewApplication creates a new application, which spans the PriceWatch functionality together
func NewApplication(c *configuration.Configuration, dbc *db.Connector, pd *db.ProductDao, ps *service.PriceService, uss *service.URLShorteningService, pr *resource.PriceResource) *Application {
	return &Application{configuration: c, connector: dbc, productDao: pd, priceService: ps, urlShorteningService: uss, priceResource: pr}
}

func (app *Application) start() {
	if app.connector != nil {
		defer app.connector.CloseConnection()
	}

	serverURL := app.configuration.GetServerURL()

	log.Println("Starting PriceWatch on: " + serverURL)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/price", app.priceResource.CheckPrice)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println(http.ListenAndServe(serverURL, router))
}
