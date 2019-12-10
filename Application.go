package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	configuration *Configuration
	dbConnector   *DBConnector
	productDao    *ProductDao
	priceResource *PriceResource
}

func NewApplication(c *Configuration, dbc *DBConnector, pd *ProductDao, pr *PriceResource) *Application {
	return &Application{configuration: c, dbConnector: dbc, productDao: pd, priceResource: pr}
}

func (app *Application) start() {
	if app.dbConnector != nil {
		defer app.dbConnector.CloseConnection()
	}

	serverURL := app.configuration.GetServerURL()

	log.Println("Starting PriceWatch on: " + serverURL)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/price", app.priceResource.checkPrice)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println(http.ListenAndServe(serverURL, router))
}
