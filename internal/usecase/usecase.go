package usecase

import (
	"context"
	"github.com/nanmenkaimak/film_library/internal/entity"
)

type UseCase interface {
	User
	UserToken
	Actor
	Film
}

type User interface {
	RegisterUser(ctx context.Context, request entity.User) (*RegisterUserResponse, error)
}

type UserToken interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JwtUserToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*JwtRenewToken, error)
}

type Actor interface {
	CreateActor(ctx context.Context, request entity.Actor) (*RegisterUserResponse, error)
	UpdateActor(ctx context.Context, request entity.UpdateMap) error
	DeleteActor(ctx context.Context, request DeleteActorRequest) error
	GetActors(ctx context.Context, request GetActorRequest) ([]entity.ActorWithFilms, error)
}

type Film interface {
	CreateFilm(ctx context.Context, request entity.FilmeWithActors) (*CreateFilmRequest, error)
	UpdateFilm(ctx context.Context, request entity.UpdateMap) error
	DeleteFilm(ctx context.Context, request DeleteFilmRequest) error
	GetFilms(ctx context.Context, request GetFilmsRequest) ([]entity.FilmeWithActors, error)
}
