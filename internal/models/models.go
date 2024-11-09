package models

import "time"

type Animal struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	ScientificName string `db:"scientific_name"`
	Province       string `db:"province"`
}

type Sighting struct {
	ID           int       `db:"id"`
	AnimalID     int       `db:"animal_id"`
	Latitude     float64   `db:"latitude"`
	Longitude    float64   `db:"longitude"`
	SightingTime time.Time `db:"sighting_time"`
}
