package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

func checkPrice(w http.ResponseWriter, r *http.Request) {
    requestedPage := string(r.URL.Query()["url"][0])
    log.Println("Price check for url: " + requestedPage)

    priceData := scrapePage(requestedPage)
    log.Println(priceData)

    json.NewEncoder(w).Encode(priceData)
}

func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/price", checkPrice)

    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]bool{"ok": true})
    })

    log.Fatal(http.ListenAndServe(":8333", router))
}
