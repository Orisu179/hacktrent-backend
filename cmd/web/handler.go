package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sightingRequest struct {
	Animal    string  `json:"animal"`
	Quantity  int     `json:"quantity"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
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
	w.WriteHeader(http.StatusAccepted)
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
	w.WriteHeader(http.StatusAccepted)
}

func (app *application) postSightings(w http.ResponseWriter, r *http.Request) {
	var sighting sightingRequest
	err := json.NewDecoder(r.Body).Decode(&sighting)
	if err != nil {
		app.logger.Error("Invalid json", "error: ", err.Error())
		http.Error(w, "Invalid json", http.StatusInternalServerError)
	}

	err = app.animals.NewSighting(sighting.Animal, sighting.Quantity, sighting.Latitude, sighting.Longitude)
	if err != nil {
		app.logger.Error("Invalid sighting", "error: ", err.Error())
		http.Error(w, "Invalid sighting", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
}

func (app *application) getSightings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.PathValue("animal")
	response, err := app.animals.GetSighting(name)
	if err != nil {
		app.logger.Error("No sighting for animal", "error: ", err.Error())
		http.Error(w, "No sighting for animal", http.StatusInternalServerError)
	}
	json, err := json.Marshal(response)
	if err != nil {
		app.logger.Error("Invalid json", "error: ", err.Error())
		http.Error(w, "Invalid json", http.StatusInternalServerError)
	}
	_, err = w.Write(json)
	if err != nil {
		app.logger.Error(err.Error(), "animal", response)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (app *application) getAllSightings(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := app.animals.GetAllSighting()
	if err != nil {
		app.logger.Error("No sighting in database", "error: ", err.Error())
		http.Error(w, "No sighting in database", http.StatusNotFound)
	}
	json, err := json.Marshal(response)
	if err != nil {
		app.logger.Error("Invalid json", "error: ", err.Error())
		http.Error(w, "Invalid json", http.StatusInternalServerError)
	}
	_, err = w.Write(json)
	if err != nil {
		app.logger.Error(err.Error(), "parsing error", "error: ", err.Error())
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// could work if we are trying to get a region of animals
// func (app *application) getAnimal(w http.ResponseWriter, r *http.Request) {
//
//}
