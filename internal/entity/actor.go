package entity

import (
	"github.com/google/uuid"
)

type Actor struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Gender   string    `json:"gender" db:"gender"`
	BirthDay string    `json:"birth_day" db:"birth_day"`
}

type ActorWithFilms struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Gender   string    `json:"gender" db:"gender"`
	BirthDay string    `json:"birth_day" db:"birth_day"`
	Films    []Film      `json:"films" db:"films"`
}
