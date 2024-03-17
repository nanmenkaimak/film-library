package usecase

import (
	"context"
	"fmt"
	"github.com/nanmenkaimak/film_library/internal/entity"
	"time"
)

func (a *Service) CreateActor(ctx context.Context, request entity.Actor) (*RegisterUserResponse, error) {
	_, err := time.Parse("2006-01-02", request.BirthDay)
	if err != nil {
		return nil, fmt.Errorf("time.Parse err: %w", err)
	}
	actorID, err := a.repo.CreateActor(request)
	if err != nil {
		return nil, fmt.Errorf("CreateActor err: %w", err)
	}
	resp := &RegisterUserResponse{
		ID: actorID,
	}

	return resp, nil
}

func (a *Service) UpdateActor(ctx context.Context, request entity.UpdateMap) error {
	err := a.repo.UpdateActor(request)
	if err != nil {
		return fmt.Errorf("UpdateActor err: %w", err)
	}
	return nil
}

func (a *Service) DeleteActor(ctx context.Context, request DeleteActorRequest) error {
	err := a.repo.DeleteActor(request.ID)
	if err != nil {
		return fmt.Errorf("DeleteActor err: %w", err)
	}
	return nil
}

func (a *Service) GetActors(ctx context.Context, request GetActorRequest) ([]entity.ActorWithFilms, error) {
	actors, err := a.repo.GetActors(request.Name)
	if err != nil {
		return nil, fmt.Errorf("GetActors err: %w", err)
	}
	return actors, nil
}