package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const listenURL = "localhost"
const listenPort = "8333"

func checkPrice(w http.ResponseWriter, r *http.Request) {
	requestedPage := string(r.URL.Query()["url"][0])
	log.Println("Price check for url: " + requestedPage)

	var output = scrapePage(requestedPage)
	log.Println(output)

	json.NewEncoder(w).Encode(output)
}

func main() {
	serverURL := listenURL + ":" + listenPort
	log.Println("Starting PriceWatch on: " + serverURL)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/price", checkPrice)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	log.Println(http.ListenAndServe(serverURL, router))
}
