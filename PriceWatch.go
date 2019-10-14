package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

const configurationFile = "./pricewatch.json"

// Configuration contains the PriceWatch settings
type Configuration struct {
	ServerURL string `env:"PRICEWATCH_SERVER_URL"`
}

func initConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(configurationFile, &configuration)
	if err != nil {
		panic(err)
	}
	log.Println(configuration.ServerURL)
	return configuration
}

func checkPrice(w http.ResponseWriter, r *http.Request) {
	requestedPage := string(r.URL.Query()["url"][0])
	log.Println("Price check for url: " + requestedPage)

	var output = scrapePage(requestedPage)
	log.Println(output)

	json.NewEncoder(w).Encode(output)
}

func main() {
	log.Println("Reading configuration from: " + configurationFile)
	configuration := initConfiguration()

	log.Println("Starting PriceWatch on: " + configuration.ServerURL)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/price", checkPrice)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println(http.ListenAndServe(configuration.ServerURL, router))
}
