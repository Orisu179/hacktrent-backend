package models

import (
	"context"
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

func (m *AnimalModel) NewAnimal(name string, province string) (string, error) {
	stmt := `INSERT INTO animal (name, province) VALUES ($1, $2)`
	result, err := m.DB.Exec(context.Background(), stmt)
	if err != nil {
		return "error", err
	}
	return result.String(), nil
}

func (m *AnimalModel) GetAnimal(id int) (Animal, error) {
	return Animal{}, nil
}

func (m *AnimalModel) NewSighting(name string, quantity int, latitude float64, longitude float64) (int, error) {
	return 0, nil
}

func (m *AnimalModel) GetSighting(animalId int) ([]Sighting, error) {
	return []Sighting{}, nil
}

func (m *AnimalModel) GetAllSighting() ([]Sighting, error) {
	return []Sighting{}, nil
}

func (m *AnimalModel) GetLatestSighting() (Sighting, error) {
	return Sighting{}, nil
}
