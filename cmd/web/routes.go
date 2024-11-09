package main

import (
	"github.com/rs/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /send/{animal}", app.jsonTest)
	handler := cors.Default().Handler(mux)

	return handler
}
