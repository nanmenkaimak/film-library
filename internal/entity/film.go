package entity

import (
	"github.com/google/uuid"
	"time"
)

type Film struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Rating      float64   `json:"rating" db:"rating"`
}

type FilmeWithActors struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	ReleaseDate string    `json:"release_date" db:"release_date"`
	Rating      float64   `json:"rating" db:"rating"`
	Actors      []Actor   `json:"actors" db:"actors"`
}

type UpdateMap struct {
	ID     uuid.UUID              `json:"id"`
	Values map[string]interface{} `json:"values"`
}
