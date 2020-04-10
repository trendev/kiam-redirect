package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusPermanentRedirect)
	w.Write([]byte(`{"message": "we moved to https://kiam.fr"}`))
}

func main() {
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
