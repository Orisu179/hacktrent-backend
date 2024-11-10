package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Animal struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Province string `db:"province"`
}

type Sighting struct {
	ID           int       `db:"id"`
	AnimalID     int       `db:"animal_id"`
	Quantity     int       `db:"quantity"`
	Latitude     float64   `db:"latitude"`
	Longitude    float64   `db:"longitude"`
	SightingTime time.Time `db:"sighting_time"`
}

type AnimalModel struct {
	DB *pgxpool.Pool
}

func (m *AnimalModel) NewAnimal(name string, province string) error {
	stmt := `INSERT INTO animals (name, province) VALUES ($1, $2)`
	_, err := m.DB.Exec(context.Background(), stmt, name, province)
	if err != nil {
		return err
	}
	return nil
}

func (m *AnimalModel) GetAnimal(id int) (Animal, error) {
	stmt := `SELECT id, name, province FROM animals WHERE id = $1`
	var a Animal
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&a.ID, &a.Name, &a.Province)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Animal{}, pgx.ErrNoRows
		} else {
			return Animal{}, err
		}
	}
	return a, nil
}

func (m *AnimalModel) NewSighting(name string, quantity int, latitude float64, longitude float64) error {
	animalQuery := `SELECT id FROM animals WHERE name = $1`
	var id int
	err := m.DB.QueryRow(context.Background(), animalQuery, name).Scan(&id)
	if err != nil {
		err := m.NewAnimal(name, "CA")
		if err != nil {
			return err
		}
	}

	stmt := `INSERT INTO sightings (animal_id, quantity, latitude, longitude, sighting_time) 
			VALUES ($1, $2, $3, $4, UTC_TIMESTAMP())`
	_, err = m.DB.Exec(context.Background(), stmt, id, quantity, latitude, longitude)
	if err != nil {
		return err
	}
	return nil
}

func (m *AnimalModel) GetSighting(animalName string) ([]Sighting, error) {
	stmt := `SELECT 
    a.id AS animal_id,
    a.name AS animal_name,
    s.latitude,
    s.longitude,
    s.sighting_time
	FROM 
		animals a
	JOIN 
		sightings s ON a.id = s.animal_id
	WHERE 
		a.name = $1
	ORDER BY 
		s.sighting_time DESC;`
	rows, err := m.DB.Query(context.Background(), stmt, animalName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sighting []Sighting
	for rows.Next() {
		var s Sighting
		err = rows.Scan(&s.ID, &s.AnimalID, &s.Quantity, &s.Latitude, &s.Longitude, &s.SightingTime)
		if err != nil {
			return nil, err
		}
		sighting = append(sighting, s)
	}

	return sighting, nil
}

func (m *AnimalModel) GetAllSighting() ([]Sighting, error) {
	stmt := `
	SELECT 
		a.id AS animal_id,
		a.name AS animal_name,
		s.latitude,
		s.longitude,
		s.sighting_time
	FROM 
		animals a
	JOIN 
		sightings s ON a.id = s.animal_id
	ORDER BY 
		s.sighting_time DESC;`
	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sighting []Sighting
	for rows.Next() {
		var s Sighting
		err = rows.Scan(&s.ID, &s.AnimalID, &s.Quantity, &s.Latitude, &s.Longitude, &s.SightingTime)
		if err != nil {
			return nil, err
		}
		sighting = append(sighting, s)
	}

	return sighting, nil
}

func (m *AnimalModel) GetLatestSighting() (Sighting, error) {
	stmt := `SELECT DISTINCT ON (a.id) 
    a.id AS animal_id,
    a.name AS animal_name,
    s.latitude,
    s.longitude,
    s.sighting_time
FROM 
    animals a
JOIN 
    sightings s ON a.id = s.animal_id
ORDER BY 
    a.id, s.sighting_time DESC;
`
	var s Sighting
	row, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return Sighting{}, err
	}
	err = row.Scan(&s.ID, &s.AnimalID, &s.Quantity, &s.Latitude, &s.Longitude)
	if err != nil {
		return Sighting{}, err
	}
	return s, nil
}
