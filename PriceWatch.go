package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Reading configuration from: " + configurationFile)
	configuration := NewConfiguration()
	log.Println(configuration)
	dbConnector := NewDbConnector(configuration)
	if dbConnector != nil {
		defer dbConnector.CloseConnection()
	}
	serverURL := configuration.GetServerURL()
	productDao := NewProductDao(dbConnector)
	priceResource := NewPriceResource(productDao)

	log.Println("Starting PriceWatch on: " + serverURL)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/price", priceResource.checkPrice)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println(http.ListenAndServe(serverURL, router))

}
