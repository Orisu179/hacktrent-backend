package main

import (
	"github.com/rs/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /animals", app.jsonTest)
	mux.HandleFunc("POST /animals", app.postAnimal)
	mux.HandleFunc("POST /sightings", app.postSightings)
	handler := cors.Default().Handler(mux)

	return handler
}
