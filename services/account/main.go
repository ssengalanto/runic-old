package main

import (
	"log"
	"net/http"

	"github.com/ssengalanto/runic/pkg/http/mux"
)

func main() {
	r := mux.New()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Runic Account Service!"))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8082", r)) //nolint:gosec //todo
}
