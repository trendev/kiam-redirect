package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "https://kiam.fr"

var newURL string

func main() {

	if rawURL := os.Getenv("URL"); rawURL != "" {
		u, err := url.ParseRequestURI(rawURL) // little control...
		if err != nil {
			log.Fatal(err)
			panic("invalid URL")
		} else {
			newURL = u.String() // rawURL can also be used there
		}
	} else {
		newURL = baseURL
	}

	log.Printf("Server starting, listening on port :9000 and redirecting request to %q", newURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, newURL, http.StatusPermanentRedirect)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
