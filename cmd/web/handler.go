package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sightingRequest struct {
	Animal    string  `json:"animal"`
	Quantity  string  `json:"quantity"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Time      string  `json:"time"`
}

type animalRequest struct {
	Animal   string `json:"animal"`
	Province string `json:"province"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	//_, err := w.Write([]byte("Hello from rare animal!"))
	animals, err := app.animals.GetLatestSighting()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	_, err = fmt.Fprintf(w, "%v\n", animals)
	if err != nil {
		return
	}
}

func (app *application) jsonTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	input := r.PathValue("animal")
	responseStruct := &sightingRequest{
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

func (app *application) postAnimal(w http.ResponseWriter, r *http.Request) {
	var animal animalRequest
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		app.logger.Error("Invalid json", "error: ", err.Error())
		http.Error(w, "Invalid json", http.StatusInternalServerError)
	}

	err = app.animals.NewAnimal(animal.Animal, animal.Province)
	if err != nil {
		app.logger.Error("Invalid animal", "error: ", err.Error())
		http.Error(w, "Invalid animal", http.StatusInternalServerError)
	}
}

func (app *application) getAnimal(w http.ResponseWriter, r *http.Request) {

}
