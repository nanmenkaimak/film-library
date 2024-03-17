package usecase

import (
	"context"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"time"
)

func (a *Service) CreateFilm(ctx context.Context, request entity.FilmeWithActors) (*CreateFilmRequest, error) {
	_, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		return nil, fmt.Errorf("time.Parse err: %w", err)
	}
	filmID, err := a.repo.CreateFilm(request)
	if err != nil {
		return nil, fmt.Errorf("CreateFilm err: %w", err)
	}
	resp := &CreateFilmRequest{
		ID: filmID,
	}

	return resp, nil
}

func (a *Service) UpdateFilm(ctx context.Context, request entity.UpdateMap) error {
	err := a.repo.UpdateFilm(request)
	if err != nil {
		return fmt.Errorf("UpdateFilm err: %w", err)
	}
	return nil
}

func (a *Service) DeleteFilm(ctx context.Context, request DeleteFilmRequest) error {
	err := a.repo.DeleteFilm(request.ID)
	if err != nil {
		return fmt.Errorf("DeleteFilm err: %w", err)
	}
	return nil
}

func (a *Service) GetFilms(ctx context.Context, request GetFilmsRequest) ([]entity.FilmeWithActors, error) {
	var films []entity.FilmeWithActors
	var err error
	if len(request.Name) > 0 {
		films, err = a.repo.GetFilmsByName(request.Name)
	} else {
		films, err = a.repo.GetFilms(request.Sorting)
	}
	if err != nil {
		return nil, fmt.Errorf("GetFilms or GetFilmsByName err: %w", err)
	}
	return films, nil
}