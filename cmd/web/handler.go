package main

import (
	"encoding/json"
	"net/http"
)

type request struct {
	Animal    string  `json:"animal"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Time      string  `json:"time"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	_, err := w.Write([]byte("Hello from rare animal!"))
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) jsonTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input := r.PathValue("animal")
	responseStruct := &request{
		Animal:    input,
		Longitude: 32.3,
		Latitude:  58.2,
		Time:      "2023-03-23",
	}
	responseJson, err := json.Marshal(responseStruct)
	if err != nil {
		app.serverError(w, r, err)
	}
	_, err = w.Write(responseJson)
	if err != nil {
		app.logger.Error(err.Error(), "animal", responseJson)
		http.Error(w, "server json parsing error", http.StatusInternalServerError)
	}
}
