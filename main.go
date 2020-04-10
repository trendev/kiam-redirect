package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "https://kiam.fr"

var newURL string

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, newURL, http.StatusPermanentRedirect)
}

func main() {

	if rawURL := os.Getenv("URL"); rawURL != "" {
		u, err := url.Parse(rawURL)
		if err != nil {
			log.Fatal(err)
		} else {
			newURL = u.String() // rawURL can also be used there
		}
	} else {
		newURL = baseURL
	}

	log.Printf("Server starting, listening on port :9000 and redirecting request to %v", newURL)

	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
