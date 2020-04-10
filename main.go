package main

import (
	"log"
	"net/http"
	"os"
)

const baseURL = "https://kiam.fr"

var url string

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func main() {

	if u := os.Getenv("URL"); u != "" {
		url = u
	} else {
		url = baseURL
	}

	log.Printf("Server starting, listening on port :9000 and redirecting request to %v", url)

	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
