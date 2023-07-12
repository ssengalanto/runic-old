package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Runic Account Service!"))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8082", mux)) //nolint:gosec //todo
}
