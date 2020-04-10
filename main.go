package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://kiam.fr", http.StatusPermanentRedirect)
}

func main() {
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
