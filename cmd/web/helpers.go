package main

import (
	"context"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) createDb() error {
	animalsTable := `
	CREATE TABLE IF NOT EXISTS animal (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    province VARCHAR(3));`

	animalSightingTable := `
	CREATE TABLE IF NOT EXISTS sighting ( 
    id SERIAL PRIMARY KEY,
    animal_id INT NOT NULL,
    latitude NUMERIC(9, 6) NOT NULL,
    longitude NUMERIC(9, 6) NOT NULL,
    sighting_time TIMESTAMP NOT NULL,
    FOREIGN KEY (animal_id) REFERENCES animal (id) ON DELETE CASCADE
);`
	_, err := app.animals.DB.Exec(context.Background(), animalsTable)
	if err != nil {
		return err
	}

	_, err = app.animals.DB.Exec(context.Background(), animalSightingTable)
	if err != nil {
		return err
	}
	return nil
}
