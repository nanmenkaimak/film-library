package repository

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/database/postgres"
	"github.com/nanmenkaimak/film_library/internal/entity"
)

type Repository interface {
	AuthRepository
	UserRepository
	ActorRepository
	FilmRepository
}

type AuthRepository interface {
	CreateUserToken(userToken entity.UserToken) error
	UpdateUserToken(userToken entity.UserToken) error
}

type UserRepository interface {
	CreateUser(newUser entity.User) (uuid.UUID, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type FilmRepository interface {
	CreateFilm(newFilm entity.FilmeWithActors) (uuid.UUID, error)
	UpdateFilm(film entity.UpdateMap) error
	DeleteFilm(filmID uuid.UUID) error
	GetFilms(sorting string) ([]entity.FilmeWithActors, error)
	GetFilmsByName(name string) ([]entity.FilmeWithActors, error)
}

type ActorRepository interface {
	CreateActor(newActor entity.Actor) (uuid.UUID, error)
	UpdateActor(actor entity.UpdateMap) error
	DeleteActor(actorID uuid.UUID) error
	GetActors(name string) ([]entity.ActorWithFilms, error)
}

type Repo struct {
	main    *postgres.Db
	replica *postgres.Db
}

func NewRepository(connMain *postgres.Db, connReplica *postgres.Db) *Repo {
	return &Repo{
		main:    connMain,
		replica: connReplica,
	}
}
