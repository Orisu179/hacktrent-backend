package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	log.Print("Starting server at 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
